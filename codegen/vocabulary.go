package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// like strings.Split but returns nil if the string is empty.
func split(s string, sep string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, sep)
}

func golangifyType(t string) string {
	switch t {
	case "string":
		return "string"
	case "integer":
		return "int64"
	case "numeric":
		return "float64"
	case "boolean":
		return "bool"
	case "array":
		return "[2]int64"
	}
	// Handle "or" between types by selecting a preferred one according to the list.
	alternatives := strings.Split(t, " or ")
	if len(alternatives) < 2 {
		alternatives = strings.Split(t, " / ")
	}
	if len(alternatives) < 2 {
		// Caller should panic.
		return ""
	}
	preferredType := alternatives[0]
	for _, alternative := range alternatives[1:] {
		if preferredTypes[alternative] > preferredTypes[preferredType] {
			preferredType = alternative
		}
	}
	return golangifyType(preferredType)
}

var preferredTypes = map[string]int{
	"array":   3,
	"integer": 2,
	"numeric": 1,
	"string":  1,
	"boolean": 1,
}

var reDigits = regexp.MustCompile(`^([0-9]+)`)

// making camel case for some particular words.
func camelCase(original string) string {
	lower := strings.ToLower(original)
	// Clear out snake case and such.
	purged := strings.ReplaceAll(lower, "_", "")
	purged = strings.ReplaceAll(purged, "-", "")
	purged = strings.ReplaceAll(purged, " ", "")
	purged = strings.ReplaceAll(purged, "(", "")
	purged = strings.ReplaceAll(purged, ")", "")
	purged = strings.ReplaceAll(purged, "'", "")
	purged = strings.ReplaceAll(purged, `"`, "")
	purged = strings.ReplaceAll(purged, "/", "Or")
	purged = strings.ReplaceAll(purged, "|", "Or")
	purged = strings.ReplaceAll(purged, "==", "Equals")
	purged = strings.ReplaceAll(purged, "=", "Equals")
	titled := strings.Title(purged)
	var parts []string
OUTER:
	for {
		if m := reDigits.FindStringSubmatch(titled); len(m) == 2 {
			parts = append(parts, m[1])
			titled = strings.Title(strings.TrimPrefix(titled, m[1]))
			if titled == "" {
				break OUTER
			}
		}
		for _, word := range vocabulary {
			if strings.HasPrefix(titled, word) {
				if replacement := compositeWords[word]; replacement != "" {
					parts = append(parts, replacement)
				} else {
					parts = append(parts, word)
				}
				titled = strings.Title(strings.TrimPrefix(titled, word))
				if titled == "" {
					break OUTER
				}
				continue OUTER
			}
		}
		break
	}
	if titled != "" {
		tidyVocabulary()
		panic(fmt.Sprintf("%q  %q  %v", original, titled, parts))
		parts = append(parts, titled)
	}
	return strings.Join(parts, "")
}

// Prints a tidy vocabulary, which is sorted correctly. Use to update this code if adding new words.
func tidyVocabulary() {
	sort.Slice(vocabulary, func(i, j int) bool {
		if strings.HasPrefix(vocabulary[j], vocabulary[i]) {
			return false
		}
		if strings.HasPrefix(vocabulary[i], vocabulary[j]) {
			return true
		}
		if vocabulary[i] < vocabulary[j] {
			return true
		}
		return false
	})
	fmt.Println(`var vocabulary = []string {`)
	prevWord := ""
	for _, word := range vocabulary {
		if word == prevWord {
			continue
		}
		prevWord = word
		fmt.Printf("\t\"%s\",\n", word)
	}
	fmt.Println(`}`)
}

var compositeWords = map[string]string{
	"Blockstats":     "BlockStats",
	"Id":             "ID",
	"Ids":            "IDs",
	"Txoutset":       "TxOutSet",
	"Txstats":        "TxStats",
	"Txsize":         "TxSize",
	"P2sh":           "P2SH",
	"Bip":            "BIP",
	"Totalsize":      "TotalSize",
	"Vsize":          "VSize",
	"Descendantsize": "DescendantSize",
	"Ancestorsize":   "AncestorSize",
	"Wtx":            "WTx",
	"Prevout":        "PrevOut",
	"Incycle":        "InCycle",
	"Bip152hb":       "BIP152Hb",
	"Scriptsig":      "ScriptSig",
	"Hd":             "HD",
}

var vocabulary = []string{
	"Abandoned",
	"Abandon",
	"Abort",
	"Absolute",
	"Accepted",
	"Accept",
	"Action",
	"Activation",
	"Active",
	"Added",
	"Addresses",
	"Address",
	"Addr",
	"Add",
	"Allowed",
	"Amounts",
	"Amount",
	"Analyze",
	"Ancestorsize",
	"Ancestors",
	"Ancestor",
	"And",
	"Asm",
	"As",
	"Automatic",
	"Aux",
	"Available",
	"Avg",
	"Avoid",
	"Backup",
	"Balances",
	"Balance",
	"Banned",
	"Ban",
	"Base",
	"Best",
	"Bind",
	"Bip152hb",
	"Bip",
	"Bits",
	"Bit",
	"Blank",
	"Blockchain",
	"Blockhash",
	"Blockstats",
	"Blocks",
	"Block",
	"Bogo",
	"Branch",
	"Bump",
	"Bytes",
	"By",
	"Capabilities",
	"Category",
	"Chain",
	"Changeaddress",
	"Change",
	"Check",
	"Chunks",
	"Clear",
	"Coinbase",
	"Combine",
	"Commands",
	"Command",
	"Comment",
	"Commitment",
	"Complete",
	"Compressed",
	"Config",
	"Confirmations",
	"Conflicts",
	"Conf",
	"Connected",
	"Connections",
	"Connection",
	"Conn",
	"Convert",
	"Count",
	"Created",
	"Create",
	"Credentials",
	"Current",
	"Cur",
	"Data",
	"Decoded",
	"Decode",
	"Default",
	"Delta",
	"Depends",
	"Derive",
	"Derivs",
	"Descendantsize",
	"Descendants",
	"Descendant",
	"Descriptors",
	"Descriptor",
	"Desc",
	"Destination",
	"Details",
	"Difficulty",
	"Dir",
	"Disable",
	"Disconnect",
	"Disk",
	"Display",
	"Download",
	"Dummy",
	"Dump",
	"Duration",
	"Elapsed",
	"Embedded",
	"Empty",
	"Enabled",
	"Encrypt",
	"Entry",
	"Enumerate",
	"Equals",
	"Errors",
	"Error",
	"Estimated",
	"Estimate",
	"Exclude",
	"External",
	"Extract",
	"Ex",
	"False",
	"Fees",
	"Fee",
	"File",
	"Filter",
	"Finalize",
	"Final",
	"Fingerprint",
	"Flag",
	"Flight",
	"Forks",
	"Format",
	"For",
	"Found",
	"Frame",
	"Free",
	"From",
	"Funded",
	"Funds",
	"Fund",
	"Generated",
	"Generate",
	"Genesis",
	"Get",
	"Groupings",
	"Hash",
	"Has",
	"Hd",
	"Headers",
	"Header",
	"Height",
	"Help",
	"Hex",
	"Historical",
	"Hwm",
	"Ids",
	"Id",
	"If",
	"Immature",
	"Import",
	"Inbound",
	"Include",
	"Increase",
	"Incremental",
	"Incycle",
	"Inc",
	"Index",
	"Info",
	"Initial",
	"Inputs",
	"Input",
	"Ins",
	"Internal",
	"Interval",
	"Involves",
	"In",
	"Is",
	"Join",
	"Keypool",
	"Keys",
	"Key",
	"Labels",
	"Label",
	"Last",
	"Left",
	"Len",
	"Level",
	"Limited",
	"Limit",
	"List",
	"Loaded",
	"Load",
	"Local",
	"Locked",
	"Lock",
	"Logging",
	"Log",
	"Long",
	"Malloc",
	"Many",
	"Mapped",
	"Master",
	"Maximum",
	"Max",
	"Median",
	"Memory",
	"Mempool",
	"Merkle",
	"Message",
	"Method",
	"Millis",
	"Mine",
	"Minimum",
	"Mining",
	"Min",
	"Missing",
	"Mode",
	"Modified",
	"Multisig",
	"Multi",
	"Mu",
	"Names",
	"Name",
	"Networks",
	"Network",
	"Net",
	"New",
	"Next",
	"Node",
	"Nonce",
	"Non",
	"Notifications",
	"Not",
	"No",
	"Num",
	"N",
	"Objects",
	"Offset",
	"Oldest",
	"Old",
	"Only",
	"On",
	"Ops",
	"Options",
	"Op",
	"Orig",
	"Or",
	"Otherwise",
	"Outputs",
	"Output",
	"Outs",
	"Out",
	"P2sh",
	"Package",
	"Parent",
	"Partial",
	"Passphrase",
	"Path",
	"Pay",
	"Peer",
	"Pending",
	"Percentiles",
	"Period",
	"Permissions",
	"Permit",
	"Ping",
	"Poll",
	"Pooled",
	"Port",
	"Position",
	"Possible",
	"Pos",
	"Precious",
	"Previous",
	"Prevout",
	"Prev",
	"Prioritise",
	"Private",
	"Privkeys",
	"Privkey",
	"Process",
	"Program",
	"Progress",
	"Proof",
	"Proposal",
	"Protocol",
	"Proxy",
	"Pruned",
	"Prune",
	"Pruning",
	"Psbts",
	"Psbt",
	"Ps",
	"Pubkeys",
	"Pubkey",
	"Purpose",
	"Query",
	"Randomize",
	"Range",
	"Rate",
	"Raw",
	"Reachable",
	"Reached",
	"Reason",
	"Received",
	"Recv",
	"Redeem",
	"Refill",
	"Reject",
	"Relay",
	"Remaining",
	"Removed",
	"Remove",
	"Replaceable",
	"Requests",
	"Request",
	"Required",
	"Req",
	"Rescan",
	"Result",
	"Reused",
	"Reuse",
	"Rewards",
	"Root",
	"Rpc",
	"Rules",
	"Rule",
	"Safe",
	"Save",
	"Scanning",
	"Scan",
	"Score",
	"Scriptsig",
	"Scripts",
	"Script",
	"Seed",
	"Segwit",
	"Send",
	"Sent",
	"Sequence",
	"Serialized",
	"Serve",
	"Services",
	"Set",
	"Signatures",
	"Signature",
	"Signers",
	"Signer",
	"Sign",
	"Sigs",
	"Sig",
	"Since",
	"Size",
	"Skip",
	"Smart",
	"Soft",
	"Solvable",
	"Spendable",
	"Spent",
	"Starting",
	"Startup",
	"Start",
	"State",
	"Statistics",
	"Stats",
	"Status",
	"Stop",
	"String",
	"Stripped",
	"Str",
	"Submit",
	"Subsidy",
	"Subtract",
	"Sub",
	"Success",
	"Sum",
	"Sw",
	"Synced",
	"Table",
	"Target",
	"Template",
	"Test",
	"The",
	"Threshold",
	"Timestamp",
	"Time",
	"Tips",
	"Totalsize",
	"Totals",
	"Total",
	"To",
	"Transactions",
	"Transaction",
	"Tries",
	"True",
	"Trusted",
	"Txes",
	"Txoutset",
	"Txsize",
	"Txstats",
	"Txs",
	"Tx",
	"Type",
	"Unbroadcast",
	"Unclaimed",
	"Unconfirmed",
	"Unknown",
	"Unload",
	"Unlocked",
	"Unlock",
	"Unsafe",
	"Unspendables",
	"Unspendable",
	"Unspents",
	"Unspent",
	"Until",
	"Untrusted",
	"Update",
	"Upgrade",
	"Upload",
	"Uptime",
	"Usage",
	"Used",
	"Use",
	"Utxo",
	"Validate",
	"Valid",
	"Value",
	"Vb",
	"Verbose",
	"Verbosity",
	"Verification",
	"Verify",
	"Version",
	"Ver",
	"Vin",
	"Vout",
	"Vsize",
	"Wait",
	"Wallets",
	"Wallet",
	"Warnings",
	"Warning",
	"Was",
	"Watching",
	"Watch",
	"Weight",
	"When",
	"Window",
	"With",
	"Witness",
	"Work",
	"Wtx",
	"X",
	"Zmq",
}
