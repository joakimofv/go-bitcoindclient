package main

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

//Arguments:
//1. blockhash    (string, required) The block hash
//2. verbosity    (numeric, optional, default=1) 0 for hex-encoded data, 1 for a json object, and 2 for json object with transaction data
//2. dummy        (numeric, optional) API-Compatibility for previous API. Must be zero or null.
var reArgument = regexp.MustCompile(`^[0-9]+.\s+([^\s]+)\s+\(([^\),]+)(?:|, ([a-z]*)(?:|, default=([^\)]*)))\)\s*(.*)`)

//                DEPRECATED. For forward compatibility use named arguments and omit this parameter.
var reContinuedComment = regexp.MustCompile(`^\s+([^\s\[\{\.\]\}].*)`)

//      [
var reArgumentArrayBegin = regexp.MustCompile(`^\s*\[`)

//      {
var reArgumentStructBegin = regexp.MustCompile(`^\s*\{`)

//        "descriptor",             (string) An output descriptor
// vout_index,                  (numeric) The zero-based output index, before a change output is added.
var reSingleMember = regexp.MustCompile(`^\s*"?([^",]*)"?,\s+\(([^\),]+)(?:|, required)\)\s*(.*)`)

//        {                         (json object) An object with output descriptor and metadata
var reAnonymousStructStart = regexp.MustCompile(`^\s*\{\s*\(json object[^\)]*\)\s*(.*)`)

//          "desc": "str",          (string, required) An output descriptor
//  "hash" : "hex",                 (string) the block hash (same as provided)
//  "confirmations" : n,            (numeric) The number of confirmations, or -1 if the block is not on the main chain
//  "nextblockhash" : "hex"         (string, optional) The hash of the next block (if available)
//  "in_active_chain" : true|false,    (boolean) Whether specified block is in the active chain or not (only present with explicit "blockhash" argument)
//          "range": n or [n,n],    (numeric or array, optional, default=1000) Up to what index HD chains should be explored (either end or [begin,end])
var reStructMemberBasic = regexp.MustCompile(`^\s*"([^"]+)"\s*:\s*[^\[\{\(\s][^\(]*\(([^\),]+)(?:|, ([a-z]+)(?:|, default=([^\)]*)))\)\s*(.*)`)

//        },
var reStructEnd = regexp.MustCompile(`^\s*\}`)

//        ...
var reArrayDeadSpace = regexp.MustCompile(`^\s*\.\.\.,?\s*$`)

//      ]
var reArrayEnd = regexp.MustCompile(`^\s*\]`)

// Result:
// Multiple top level struct marker (inserted by our parser)
////ALTERNATIVE: (for verbosity = 2)
var reAlternative = regexp.MustCompile(`^` + alternativeMarker + `(.+)`)

//n    (numeric) The total amount in BTC received for this wallet.
//"str"    (string) The new bitcoin address
var reSingle = regexp.MustCompile(`^\s*(?:"([^",]*)"|([^"\s\{\[]+))\s+\(([^\),]+)\)\s*(.*)`)

//null    (json null)
var reEmpty = regexp.MustCompile(`^\s*null\s*\(json null\)`)

//{                                 (json object)
//reAnonymousStructStart

// [                               (json array) The result of the mempool acceptance test for each raw transaction in the input array.
var reAnonymousArrayStart = regexp.MustCompile(`^\s*\[\s*\(json array\)\s*(.*)`)

//                                 Returns results for each transaction in the same order they were passed in.
//                                 It is possible for transactions to not be fully validated ('allowed' unset) if another transaction failed.
//reContinuedComment

//      "scriptSig" : {                (json object) The script
var reStructMemberStructStart = regexp.MustCompile(`^\s*"([^"]+)"\s*:\s*\{[^\(]*\(json object(?:|, ([a-z]+)[^\)]*)\)\s*(.*)`)

//  "tx" : [                        (json array) The transaction ids
var reStructMemberArrayStart = regexp.MustCompile(`^\s*"([^"]+)"\s*:\s*\[[^\(]*\(json array(?:|, ([a-z]+)[^\)]*)\)\s*(.*)`)

var reStructMember = regexp.MustCompile(`^\s*"[^"]+"\s*:\s*[^\(]+\s+\([^\(]+\)`)

//       "address": amount,    (numeric or string, required) The bitcoin address is the key, the numeric amount (can be string) in BTC is the value
var reKeyValue = regexp.MustCompile(`[Kk]ey.*[Vv]alue`)

//  ...,              Same output as verbosity = 1
//      ...,          The transactions in the format of the getrawtransaction RPC. Different from verbosity = 1 "tx" result
//      ...                                    The layout is the same as the output of decoderawtransaction.
var reManualInterpretationNeeded = regexp.MustCompile(`^\s*\.\.\.,?\s+(.+)`)

var reDeprecated = regexp.MustCompile(`DEPRECATED|[Ii]gnored dummy value|[Tt]his value is ignored|[Rr]emains for backward compatibility`)

//    "scanning" : {                          (json object) current scanning details, or false if no scan is in progress
var reMayBeFalse = regexp.MustCompile(`, or false`)

type structInfo struct {
	TypeName string
	Comment  string

	Members []memberInfo

	// A "special kind" of of struct is a map. Members should be empty if these are set.
	MapKeyType     string
	MapValueType   string
	MapValueStruct structInMapInfo

	// Modify the type and custom marshalling behaviour.
	ReplaceByMember bool
	Anonymous       bool // Aka inline struct.
	ArrayLevel      int
	NilIsFalse      bool
}

type structInMapInfo struct {
	TypeName string
	Comment  string
	Members  []memberInfo

	// Map structs may not in itself be a map.
}

type basicTypeInfo struct {
	Name       string
	TypeName   string
	Comment    string
	JsonName   string
	ArrayLevel int
}

type memberInfo struct {
	JsonName string
	Name     string
	Optional bool
	Comment  string
	Default  string

	// It can be either a basic type, a struct or an array (of the basic type or struct).
	ArrayLevel int
	// Use BasicType if non-zero, else use Struct.
	BasicType string
	Struct    structInfo
}

func newStructInfo(name string, jsonComments []string) structInfo {
	s := structInfo{
		TypeName: name,
	}
	if len(jsonComments) == 0 {
		return s
	}
	if reEmpty.MatchString(jsonComments[0]) {
		return s
	}

	lf := NewLineFeeder(jsonComments)
	if reArgument.MatchString(jsonComments[0]) {
		s.buildArgumentStruct(lf)
	} else {
		s.buildResultStruct(lf)
	}
	return s
}

func (s *structInfo) buildArgumentStruct(lf *LineFeeder) {
ARGUMENTS:
	for lf.Scan() {
		m := reArgument.FindStringSubmatch(lf.Text())
		if len(m) != 6 {
			panic(lf.Debug())
		}
		member := memberInfo{
			JsonName: m[1],
			Name:     camelCase(m[1]),
			Optional: m[3] == "optional",
			Comment:  appendComments(m[5], continuedComment(lf), defaultComment(m[4])),
			Default:  m[4],
		}

		if reDeprecated.MatchString(member.Comment) {
			// Skip past this member.
			for lf.Scan() {
				if reArgument.MatchString(lf.Text()) {
					lf.Rewind()
					break
				}
			}
			continue ARGUMENTS
		}

		if m[2] == "json array" {
			// Walk past the opening symbol "[".
			lf.Scan()
			if !reArgumentArrayBegin.MatchString(lf.Text()) {
				panic(lf.Debug())
			}
			// Interpret the following lines as array.
			var basicType basicTypeInfo
			basicType, member.Struct = buildArray(s.TypeName+member.Name, lf)
			member.BasicType = basicType.TypeName
			member.Comment = appendComments(member.Comment, basicType.Comment)
			member.ArrayLevel = basicType.ArrayLevel + member.Struct.ArrayLevel
		} else if m[2] == "json object" {
			// Walk past the opening symbol "{".
			lf.Scan()
			if !reArgumentStructBegin.MatchString(lf.Text()) {
				panic(lf.Debug())
			}
			// Interpret the following lines as struct.
			member.Struct = buildStruct(s.TypeName+member.Name, "", lf)
			if !member.Optional {
				// Let it be inline embedded in the parent structInfo.
				member.Struct.Anonymous = true
			}
		} else {
			member.BasicType = golangifyType(m[2])
			if member.BasicType == "" {
				panic(lf.Debug())
			}
		}
		s.Members = append(s.Members, member)
	}
}

func (s *structInfo) buildResultStruct(lf *LineFeeder) {
	var order []string
	var basicTypes []basicTypeInfo
	var structs []structInfo
	char := byte('A')

	for lf.Scan() {
		structTypeName := s.TypeName
		if m := reAlternative.FindStringSubmatch(lf.Text()); len(m) == 2 {
			structTypeName += camelCase(m[1])
			if !lf.Scan() {
				panic(lf.Debug())
			}
		}
		// It's either null, a single member, an anonymous struct, or an anonymous array.
		if reEmpty.MatchString(lf.Text()) {
			continuedComment(lf)
			continue
		} else if m := reSingle.FindStringSubmatch(lf.Text()); len(m) == 5 {
			typeName := golangifyType(m[3])
			if typeName == "" {
				panic(lf.Debug())
			}
			name := m[1]
			if name == "" {
				name = m[2]
			}
			basicTypes = append(basicTypes, basicTypeInfo{
				Name:     camelCase(name),
				TypeName: typeName,
				Comment:  appendComments(m[4], continuedComment(lf)),
			})
			order = append(order, "basic")
		} else if m = reAnonymousStructStart.FindStringSubmatch(lf.Text()); len(m) == 2 {
			comment := appendComments(m[1], continuedComment(lf))
			subStruct := buildStruct(structTypeName, comment, lf)
			structs = append(structs, subStruct)
			order = append(order, "struct")
		} else if m = reAnonymousArrayStart.FindStringSubmatch(lf.Text()); len(m) == 2 {
			comment := appendComments(m[1], continuedComment(lf))
			typePrefix := structTypeName + "Element"
			if char > 'A' {
				typePrefix += string([]byte{char})
			}
			basicType, subStruct := buildArray(typePrefix, lf)
			char++
			if !reflect.ValueOf(basicType).IsZero() {
				basicType.Comment = appendComments(comment, basicType.Comment)
				basicTypes = append(basicTypes, basicType)
				order = append(order, "basic")
			} else if !reflect.ValueOf(subStruct).IsZero() {
				subStruct.Comment = appendComments(comment, subStruct.Comment)
				structs = append(structs, subStruct)
				order = append(order, "struct")
			} else {
				panic(lf.Debug())
			}
		} else {
			panic(lf.Debug())
		}
	}

	if len(basicTypes) == 0 && len(structs) == 1 && structs[0].ArrayLevel == 0 {
		// Use the substruct as the main struct, but keep the name.
		structs[0].TypeName = s.TypeName
		*s = structs[0]
		return
	}

	// Gather the alternative types into the base struct.
	s.ReplaceByMember = true
	for _, t := range order {
		if t == "basic" {
			basicType := basicTypes[0]
			basicTypes = basicTypes[1:]
			s.Members = append(s.Members, memberInfo{
				Name:       basicType.Name,
				Comment:    basicType.Comment,
				BasicType:  basicType.TypeName,
				ArrayLevel: basicType.ArrayLevel,
			})
		} else {
			subStruct := structs[0]
			structs = structs[1:]
			if subStruct.ArrayLevel == 0 {
				// Would like to do this, but can't since some structs need to be borrowed by other Resps
				// (those crossreferences are done by todofix after codegen has finished).
				//subStruct.Anonymous = true
			}
			s.Members = append(s.Members, memberInfo{
				Name:       strings.Replace(strings.TrimPrefix(subStruct.TypeName, s.TypeName), "Element", "Array", 1),
				Comment:    subStruct.Comment,
				Struct:     subStruct,
				ArrayLevel: subStruct.ArrayLevel,
			})
		}
	}
}

func continuedComment(lf *LineFeeder) string {
	comment := ""
	// Gather the comment continuation.
	for lf.Scan() {
		if reStructMember.MatchString(lf.Text()) || reSingleMember.MatchString(lf.Text()) || reArrayDeadSpace.MatchString(lf.Text()) {
			// Real content, rewind and break.
		} else if m := reContinuedComment.FindStringSubmatch(lf.Text()); len(m) == 2 {
			if comment != "" {
				comment += "\n"
			}
			comment += m[1]
			continue
		}
		lf.Rewind()
		break
	}
	return comment
}

func defaultComment(defaultValue string) string {
	if defaultValue == "" {
		return ""
	}
	return "Default: " + defaultValue
}

func appendComments(ss ...string) string {
	output := ""
	for _, s := range ss {
		if s != "" {
			if output != "" {
				output += "\n"
			}
			parts := strings.Split(s, "\n")
			for i, part := range parts {
				if i != 0 {
					output += "\n"
				}
				output += strings.TrimRight(part, " \t")
			}
		}
	}
	return output
}

func appendDefinitions(ss ...string) string {
	output := ""
	for _, s := range ss {
		if s != "" {
			if output != "" {
				output += "\n\n"
			}
			output += strings.Trim(s, "\n")
		}
	}
	return output
}

func buildArray(typePrefix string, lf *LineFeeder) (basicTypeInfo, structInfo) {
	var order []string
	var basicTypes []basicTypeInfo
	var structs []structInfo
	char := byte('A')

	subTypePrefix := func(text string) string {
		subTypePrefix := typePrefix
		addChar := true
		if len(order) == 0 {
			addChar = false
			// If it's a single struct then don't want a strange letter in the name.
			// Peek ahead until the array end marker.
			arrayEnd := "]"
			anotherObjectStart := `[^.\}\]\s]`
			for i := 0; i < len(text); i++ {
				if text[i] == ' ' {
					if i >= 2 {
						arrayEnd = " " + arrayEnd
					}
					anotherObjectStart = " " + anotherObjectStart
				} else {
					break
				}
			}
			anotherObject := regexp.MustCompile("^" + anotherObjectStart)
			wentAhead := 1
			for lf.Scan() {
				if strings.HasPrefix(lf.Text(), arrayEnd) {
					break
				}
				if anotherObject.MatchString(lf.Text()) {
					addChar = true
					break
				}
				wentAhead++
			}
			for i := 0; i < wentAhead; i++ {
				lf.Rewind()
			}
		}
		if addChar {
			subTypePrefix += string([]byte{char})
		}
		char++
		return subTypePrefix
	}

	for lf.Scan() {
		if reArrayEnd.MatchString(lf.Text()) {
			break
		}
		// It's either a single member, an anonymous struct, or another level of array.
		if m := reSingleMember.FindStringSubmatch(lf.Text()); len(m) == 4 {
			typeName := golangifyType(m[2])
			if typeName == "" {
				panic(lf.Debug())
			}
			name := camelCase(m[1])
			if name == "" {
				name = string([]byte{char})
				char++
			}
			basicTypes = append(basicTypes, basicTypeInfo{
				Name:       name,
				TypeName:   typeName,
				Comment:    appendComments(m[3], continuedComment(lf)),
				JsonName:   m[1],
				ArrayLevel: 1,
			})
			order = append(order, "basic")
		} else if m = reAnonymousStructStart.FindStringSubmatch(lf.Text()); len(m) == 2 {
			text := lf.Text()
			comment := appendComments(m[1], continuedComment(lf))
			s := buildStruct(subTypePrefix(text), comment, lf)
			s.ArrayLevel = 1
			structs = append(structs, s)
			order = append(order, "struct")
		} else if m = reAnonymousArrayStart.FindStringSubmatch(lf.Text()); len(m) == 2 {
			text := lf.Text()
			comment := appendComments(m[1], continuedComment(lf))
			basicType, subStruct := buildArray(subTypePrefix(text), lf)
			if !reflect.ValueOf(basicType).IsZero() {
				basicType.Comment = appendComments(comment, basicType.Comment)
				basicType.ArrayLevel += 1
				basicTypes = append(basicTypes, basicType)
				order = append(order, "basic")
			} else if !reflect.ValueOf(subStruct).IsZero() {
				subStruct.Comment = appendComments(comment, subStruct.Comment)
				subStruct.ArrayLevel += 1
				structs = append(structs, subStruct)
				order = append(order, "struct")
			} else {
				panic(lf.Debug())
			}
		} else if reArrayDeadSpace.MatchString(lf.Text()) {
			continue
		} else {
			panic(lf.Debug())
		}
	}

	if len(basicTypes) == 0 && len(structs) == 0 {
		// Make a struct with a TODO.
		return basicTypeInfo{}, structInfo{
			TypeName:   typePrefix,
			Members:    []memberInfo{{Comment: "FIX " + typePrefix}},
			ArrayLevel: 1,
		}
	}
	if len(basicTypes) == 0 && len(structs) == 1 {
		return basicTypeInfo{}, structs[0]
	}
	if len(structs) == 0 {
		if len(basicTypes) > 0 {
			basicType := basicTypes[0].TypeName
			name := basicTypes[0].Name
			comment := strings.Trim(fmt.Sprintf("Element: %s\t%s", basicTypes[0].Name, basicTypes[0].Comment), " \t")

			if len(basicTypes) > 1 {
				comment = strings.Trim(fmt.Sprintf("Element: %s\t%s", basicTypes[0].JsonName, basicTypes[0].Comment), " \t")
				// See if the type is the same for all, then we can make do with just one parameter after all.
				for i := 1; i < len(basicTypes); i++ {
					bt := basicTypes[i]
					if bt.TypeName != basicType {
						goto NOT_COMPATIBLE
					}
					name += "Or" + bt.Name
					comment += strings.Trim(fmt.Sprintf("\nElement: %s\t%s", bt.JsonName, bt.Comment), " \t")
				}
			}
			basicTypes[0].Name = name
			basicTypes[0].Comment = alignTabs(comment)

			return basicTypes[0], structInfo{}
		}
	NOT_COMPATIBLE:
	}

	// Gather the alternative types into a holder struct.
	holderStruct := structInfo{
		TypeName:        typePrefix,
		Comment:         "Holder of alternative parameter formats, only one will be used, the first that is non-zero.",
		ReplaceByMember: true,
		ArrayLevel:      1,
	}
	for _, t := range order {
		if t == "basic" {
			basicType := basicTypes[0]
			basicTypes = basicTypes[1:]
			basicType.ArrayLevel -= 1
			holderStruct.Members = append(holderStruct.Members, memberInfo{
				Name:       basicType.Name,
				Comment:    basicType.Comment,
				BasicType:  basicType.TypeName,
				ArrayLevel: basicType.ArrayLevel,
			})
		} else {
			s := structs[0]
			structs = structs[1:]
			s.ArrayLevel -= 1
			if s.ArrayLevel < 1 {
				s.Anonymous = true
			}
			holderStruct.Members = append(holderStruct.Members, memberInfo{
				Name:       strings.TrimPrefix(s.TypeName, typePrefix),
				Comment:    s.Comment,
				Struct:     s,
				ArrayLevel: s.ArrayLevel,
			})
		}
	}
	return basicTypeInfo{}, holderStruct
}

func alignTabs(s string) string {
	parts := strings.Split(s, "\n")
	results := make([]string, len(parts))
	for {
		anyLen := false
		allTabs := true
		highestTabLen := 0
		for i, part := range parts {
			if len(part) == 0 {
				continue
			}
			anyLen = true
			if part[0] == '\t' {
				if len(results[i]) > highestTabLen {
					highestTabLen = len(results[i])
				}
				continue
			}
			allTabs = false
			results[i] += string(part[:1])
			parts[i] = string(part[1:])
		}
		if !anyLen {
			break
		}
		if allTabs {
			for i, part := range parts {
				if len(part) == 0 {
					continue
				}
				parts[i] = string(part[1:])
				for n := len(results[i]); n < highestTabLen+4; n++ {
					results[i] += " "
				}
			}
		}
	}
	return strings.Join(results, "\n")
}

func buildStruct(typePrefix string, comment string, lf *LineFeeder) structInfo {
	s := structInfo{
		TypeName: typePrefix,
		Comment:  comment,
	}

	uniqueMemberName := func(name string) string {
		// Handle duplicate names that we get after camelCase sometimes.
		char := []byte{'1'}
	UNIQUE_NAME:
		for {
			for _, member := range s.Members {
				if name == member.Name {
					name = strings.TrimSuffix(name, string(char))
					char[0] += 1
					name += string(char)
					continue UNIQUE_NAME
				}
			}
			break
		}
		return name
	}

	// Entering this function, we have walked past the initial "{" and comments,
	// so the members should start from the first line we scan.
	for lf.Scan() {
		if reStructEnd.MatchString(lf.Text()) {
			break
		}
		// It's either a basic member, a struct member or an array member.
		if m := reStructMemberBasic.FindStringSubmatch(lf.Text()); len(m) == 6 {
			typeName := golangifyType(m[2])
			if typeName == "" {
				panic(lf.Debug())
			}
			member := memberInfo{
				JsonName:  m[1],
				Name:      camelCase(m[1]),
				Optional:  m[3] == "optional",
				Comment:   appendComments(m[5], continuedComment(lf), defaultComment(m[4])),
				Default:   m[4],
				BasicType: typeName,
			}
			member.Name = uniqueMemberName(member.Name)
			s.Members = append(s.Members, member)
		} else if m = reStructMemberStructStart.FindStringSubmatch(lf.Text()); len(m) == 4 {
			member := memberInfo{
				JsonName: m[1],
				Name:     camelCase(m[1]),
				Optional: m[2] == "optional",
				Comment:  appendComments(m[3], continuedComment(lf)),
			}
			member.Struct = buildStruct(typePrefix+member.Name, "", lf)
			// Special handling of nasty cases where the type might be different.
			if reMayBeFalse.MatchString(member.Comment) {
				member.Optional = true
				member.Struct.NilIsFalse = true
			}
			if !member.Optional {
				// Let it be inline embedded in the parent structInfo. Only arrays need named types.
				member.Struct.Anonymous = true
			}
			member.Name = uniqueMemberName(member.Name)
			s.Members = append(s.Members, member)
		} else if m = reStructMemberArrayStart.FindStringSubmatch(lf.Text()); len(m) == 4 {
			member := memberInfo{
				JsonName: m[1],
				Name:     camelCase(m[1]),
				Optional: m[2] == "optional",
				Comment:  appendComments(m[3], continuedComment(lf)),
			}
			member.Name = uniqueMemberName(member.Name)
			var basicType basicTypeInfo
			basicType, member.Struct = buildArray(typePrefix+member.Name, lf)
			member.BasicType = basicType.TypeName
			member.Comment = appendComments(member.Comment, basicType.Comment)
			member.ArrayLevel = basicType.ArrayLevel + member.Struct.ArrayLevel
			s.Members = append(s.Members, member)
		} else if m = reManualInterpretationNeeded.FindStringSubmatch(lf.Text()); len(m) == 2 {
			member := memberInfo{
				Comment: appendComments(m[1], continuedComment(lf)),
			}
			s.Members = append(s.Members, member)
		} else if reArrayDeadSpace.MatchString(lf.Text()) {
			// How can there be "..." in a struct? Only if it is in fact a dict of key-value pairs.
			// Then we'll remake it into a map, using the only member as the value definition.
			// There must be one previous member: The "key-value pair" definition.
			if len(s.Members) != 1 {
				// Handle special cases, that should get TODOs instead of panicing.
				if s.TypeName == "DecodePsbtRespInputsNonWitnessUtxo" {
					member := memberInfo{
						Comment: "FIX " + typePrefix,
					}
					s.Members = append(s.Members, member)
					continue
				}
				fmt.Println(s.TypeName)
				panic(lf.Debug())
			}

			s.MapKeyType = "string"
			member := s.Members[0]
			if member.BasicType != "" {
				if !reKeyValue.MatchString(member.Comment) {
					// Nevermind, the comments are not consistent.
					//panic(lf.Debug())
				}
				s.MapValueType = member.BasicType
			} else {
				s.MapValueType = member.Struct.TypeName
				// We're gonna remove the struct as a member, but need to still define it so that the code gets printed.
				// A structInfo can't hold another structInfo, that would be a import loop, so shuffle the values to another type.
				s.MapValueStruct = structInMapInfo{
					TypeName: member.Struct.TypeName,
					Comment:  member.Struct.Comment,
					Members:  member.Struct.Members,
				}
			}
			// There may not be any more lines.
			lf.Scan()
			if !reStructEnd.MatchString(lf.Text()) {
				panic(lf.Debug())
			}
			lf.Rewind()

			// Delete the member and instead use the structInfo to represent the map.
			s.Members = nil
			s.Comment = appendComments(s.Comment, member.Comment)
		} else {
			panic(lf.Debug())
		}
	}
	return s
}

func (s structInfo) Definition() string {
	output := ""
	if s.NilIsFalse {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "unmarshal-false-to-nil", s); err != nil {
			panic(err)
		}
		output += buf.String()
		s.TypeName += "Contents"
	}
	if s.Comment != "" {
		//output = "// " + s.TypeName + ":"
		output += "// " + strings.ReplaceAll(s.Comment, "\n", "\n// ") + "\n"
	}
	output += "type " + s.TypeName + " struct {\n"
	members, otherDefinitions, _ := s.presentMembers(1)
	output += members + "\n}"
	output = appendDefinitions(output, otherDefinitions)
	return output
}

func (sm structInMapInfo) Definition() string {
	s := structInfo{
		TypeName: sm.TypeName,
		Comment:  sm.Comment,
		Members:  sm.Members,
	}
	return s.Definition()
}

func (s structInfo) presentMembers(indentLevel int) (output string, otherDefinitions string, imports map[string]struct{}) {
	imports = make(map[string]struct{})
	indent := ""
	for i := 0; i < indentLevel; i++ {
		indent += "\t"
	}
	for i, member := range s.Members {
		if i > 0 {
			output += "\n\n"
		}
		if member.Comment != "" {
			if member.BasicType == "" && reflect.ValueOf(member.Struct).IsZero() {
				// Only comment and no contents means it's a TODO.
				output += indent + "// TODO: " + strings.ReplaceAll(member.Comment, "\n", "\n"+indent+"// ")
				continue
			}
			output += indent + "// " + strings.ReplaceAll(member.Comment, "\n", "\n"+indent+"// ") + "\n"
		}
		output += indent + member.Name + " "
		if member.ArrayLevel > 0 {
			for n := 0; n < member.ArrayLevel; n++ {
				output += "[]"
			}
		} else if member.Optional {
			switch member.BasicType {
			case "string":
				goto NO_POINTER
			case "int", "float64":
				if member.Default == "0" {
					goto NO_POINTER
				}
			case "bool":
				if member.Default == "false" {
					goto NO_POINTER
				}
			case "[2]int":
				if member.Default == "[0,0]" {
					goto NO_POINTER
				}
			}
			if member.Struct.MapKeyType != "" {
				goto NO_POINTER
			}
			output += "*"
		NO_POINTER:
		}
		if member.BasicType != "" {
			output += member.BasicType
		} else if !reflect.ValueOf(member.Struct).IsZero() {
			// Either:
			// Name map[MapKeyType]MapValueType `json:"JsonName"`
			// Or:
			// Name struct {
			//		...
			// } `json:"JsonName"`
			// Or:
			// Name StructType `json:"JsonName"`
			if member.Struct.MapKeyType != "" {
				output += fmt.Sprintf("map[%s]%s", member.Struct.MapKeyType, member.Struct.MapValueType)
				if !reflect.ValueOf(member.Struct.MapValueStruct).IsZero() {
					otherDefinitions = appendDefinitions(otherDefinitions, member.Struct.MapValueStruct.Definition())
				}
			} else if member.Struct.Anonymous {
				output += "struct {\n"
				members, moreDefinitions, imps := member.Struct.presentMembers(indentLevel + 1)
				output += members
				otherDefinitions = appendDefinitions(otherDefinitions, moreDefinitions)
				for k, v := range imps {
					imports[k] = v
				}
				output += "\n" + indent + "}"
			} else {
				output += member.Struct.TypeName
				// Define the struct somewhere else.
				otherDefinitions = appendDefinitions(otherDefinitions, member.Struct.Definition())
			}
		} else {
			panic(spew.Sdump(s))
		}
		if !s.ReplaceByMember {
			output += " `json:\"" + member.JsonName
			if member.Optional {
				output += ",omitempty"
			}
			output += "\"`"
		}
	}
	if s.ReplaceByMember {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "marshal", s); err != nil {
			panic(err)
		}
		if err := tmpl.ExecuteTemplate(buf, "unmarshal", s); err != nil {
			panic(err)
		}
		otherDefinitions = appendDefinitions(buf.String(), otherDefinitions)

		imports["bytes"] = struct{}{}
		if len(s.Members) > 1 {
			imports["reflect"] = struct{}{}
		}
	}
	if s.NilIsFalse {
		imports["bytes"] = struct{}{}
	}
	return
}
