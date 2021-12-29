package main

import (
	"bufio"
	"bytes"
	"regexp"
	"strings"
)

type groupInfo struct {
	name    string
	methods []string
}

func parseHelp(help []byte) (groups []groupInfo) {
	reGroup := regexp.MustCompile(`== ([A-z]+) ==`)
	reMethod := regexp.MustCompile(`^([^\s]+)`)

	scanner := bufio.NewScanner(bytes.NewBuffer(help))
	var currGroup groupInfo
	for scanner.Scan() {
		if m := reGroup.FindSubmatch(scanner.Bytes()); len(m) == 2 {
			if len(currGroup.methods) > 0 {
				groups = append(groups, currGroup)
			}
			currGroup = groupInfo{name: string(m[1])}
		} else if m := reMethod.FindSubmatch(scanner.Bytes()); len(m) == 2 {
			currGroup.methods = append(currGroup.methods, string(m[1]))
		}
	}
	if len(currGroup.methods) > 0 {
		groups = append(groups, currGroup)
	}
	return
}

type methodInfo struct {
	Lowercase string
	Camelcase string
	Comment   string
	Arguments []string
	Results   []string

	Req  structInfo
	Resp structInfo
}

var alternativeMarker = "// ALTERNATIVE "

func parseMethodHelp(help []byte) (info methodInfo, err error) {
	scanner := bufio.NewScanner(bytes.NewBuffer(help))

	reLowercase := regexp.MustCompile(`^([^\s]+)`)
	for scanner.Scan() {
		if m := reLowercase.FindSubmatch(scanner.Bytes()); len(m) == 2 {
			info.Lowercase = string(m[1])
			break
		}
	}
	info.Camelcase = camelCase(info.Lowercase)

	reNonEmpty := regexp.MustCompile(`[^\s]+`)
	reArguments := regexp.MustCompile(`^Arguments:`)
	reResult := regexp.MustCompile(`^Result(:| ([^:]+):)`)

	for scanner.Scan() {
		if reArguments.Match(scanner.Bytes()) {
			break
		}
		if m := reResult.FindSubmatch(scanner.Bytes()); len(m) > 0 {
			if len(m) > 2 && len(m[2]) > 0 {
				info.Results = append(info.Results, alternativeMarker+string(m[2]))
			}
			goto RESULT
		}
		if reNonEmpty.Match(scanner.Bytes()) {
			if info.Comment != "" {
				info.Comment += "\n"
			}
			info.Comment += strings.TrimRight(scanner.Text(), " \t")
		}
	}

	for scanner.Scan() {
		if reNonEmpty.Match(scanner.Bytes()) {
			info.Arguments = append(info.Arguments, strings.TrimRight(scanner.Text(), " \t"))
		} else {
			break
		}
	}

	for scanner.Scan() {
		if m := reResult.FindSubmatch(scanner.Bytes()); len(m) > 0 {
			if len(m) > 2 && len(m[2]) > 0 {
				info.Results = append(info.Results, alternativeMarker+string(m[2]))
			}
			goto RESULT
		}
	}
RESULT:
	for scanner.Scan() {
		if reNonEmpty.Match(scanner.Bytes()) {
			info.Results = append(info.Results, strings.TrimRight(scanner.Text(), " \t"))
		} else {
			// Peek at the next line.
			scanner.Scan()
			// Is it a new result alternative format?
			if m := reResult.FindSubmatch(scanner.Bytes()); len(m) > 0 {
				if len(m) > 2 && len(m[2]) > 0 {
					info.Results = append(info.Results, alternativeMarker+string(m[2]))
				}
				continue
			}
			// Annoyingly, there is a premature empty line here in the help for "testmempoolaccept".
			// Work around it by see if the next line is json stuff.
			// In the normal case the next line after the end would be "Examples:".
			if strings.Contains(scanner.Text(), "json") {
				info.Results = append(info.Results, strings.TrimRight(scanner.Text(), " \t"))
				continue
			}
			break
		}
	}

	return
}
