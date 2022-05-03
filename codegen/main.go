package main

import (
	"bytes"
	"embed"
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"golang.org/x/xerrors"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("%+v\n", err)
	}
}

//go:embed templates
var templates embed.FS
var tmpl *template.Template

func init() {
	t, err := template.New("").Funcs(template.FuncMap{
		"HasPrefix": strings.HasPrefix,
		"Split":     split,
		"sub":       func(i int, s int) int { return i - s },
	}).ParseFS(templates, "templates/*.tmpl")
	if err != nil {
		panic(err)
	}
	tmpl = t
}

func run() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	if filepath.Base(wd) != "go-bitcoindclient" {
		return errors.New("This needs to be run in the working directory 'go-bitcoindclient'")
	}
	help, err := exec.Command("bitcoin-cli", "-regtest", "help").Output()
	if err != nil {
		return errors.New("In order to run this you need to have bitcoin-cli installed, and bitcoind running in regtest mode.")
	}
	// in v23 the '== Generating ==' section mysteriously disappeared from the list,
	// but the methods still exist and can be polled by 'help method', so add them manually.
	if !bytes.Contains(help, []byte("== Generating ==")) {
		help = append(help, []byte(`
== Generating ==
generateblock
generatetoaddress
generatetodescriptor
`)...)
	}

	groups := parseHelp(help)
	var allMethods []methodInfo
	var contents [][]byte
	for _, group := range groups {
		var methodInfos []methodInfo
		imports := map[string]struct{}{
			"context": struct{}{},
		}
		for _, method := range group.methods {
			if len(os.Args) == 2 && method != os.Args[1] {
				continue
			}
			help, err := exec.Command("bitcoin-cli", "-regtest", "help", method).CombinedOutput()
			if err != nil {
				return xerrors.Errorf("Could not get help for %q: %s", method, help)
			}
			info, err := parseMethodHelp(help)
			if err != nil {
				return xerrors.Errorf("parseMethodHelp for %q: %w", method, err)
			}
			info.Req = newStructInfo(info.Camelcase+"Req", info.Arguments)
			info.Resp = newStructInfo(info.Camelcase+"Resp", info.Results)
			_, _, imps := info.Req.presentMembers(0)
			for key, _ := range imps {
				imports[key] = struct{}{}
			}
			_, _, imps = info.Resp.presentMembers(0)
			for key, _ := range imps {
				imports[key] = struct{}{}
			}
			if len(info.Resp.Members) > 0 {
				imports["encoding/json"] = struct{}{}
			}
			methodInfos = append(methodInfos, info)
		}
		allMethods = append(allMethods, methodInfos...)
		var sortedImports []string
		for imp, _ := range imports {
			sortedImports = append(sortedImports, imp)
		}
		sort.Strings(sortedImports)
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "group", sortedImports); err != nil {
			return xerrors.Errorf("ExecuteTemplate 'group': %w", err)
		}
		for _, info := range methodInfos {
			if err := tmpl.ExecuteTemplate(buf, "method", info); err != nil {
				return xerrors.Errorf("ExecuteTemplate 'method': %w", err)
			}
		}
		contents = append(contents, buf.Bytes())
	}
	for i, group := range groups {
		filename := strings.ToLower(group.name) + ".go"
		// See if there is already a file and it has identical contents.
		oldContents, err := os.ReadFile(filename)
		if err == nil {
			if bytes.Equal(oldContents, contents[i]) {
				continue
			}
		}
		// Write the content to the file (overwriting any existing file).
		if err := os.WriteFile(filename, contents[i], 0644); err != nil {
			return xerrors.Errorf("Could not write source code file: %w", err)
		}
	}
	// Generate tests too.
	buf := new(bytes.Buffer)
	// Push the "Stop" test to the back because it shuts down the bitcoind.
	// Push the "CreateWallet" test to the front because some other methods depend on there being a wallet.
	sort.Slice(allMethods, func(i, j int) bool {
		if allMethods[i].Camelcase == "CreateWallet" || allMethods[j].Camelcase == "Stop" {
			return true
		}
		if allMethods[i].Camelcase == "Stop" || allMethods[j].Camelcase == "CreateWallet" {
			return false
		}
		return allMethods[i].Camelcase < allMethods[j].Camelcase
	})
	if err := tmpl.ExecuteTemplate(buf, "rpc_test", allMethods); err != nil {
		return xerrors.Errorf("ExecuteTemplate 'rpc_test': %w", err)
	}
	content := buf.Bytes()
	filename := "rpc_test.go"
	// See if there is already a file and it has identical contents.
	oldContent, err := os.ReadFile(filename)
	if err == nil {
		if bytes.Equal(oldContent, content) {
			goto AFTER_TEST_WRITE
		}
	}
	// Write the content to the file (overwriting any existing file).
	if err := os.WriteFile(filename, content, 0644); err != nil {
		return xerrors.Errorf("Could not write test code file: %w", err)
	}
AFTER_TEST_WRITE:

	b, err := exec.Command("go", "fmt").CombinedOutput()
	if err != nil {
		return xerrors.Errorf("Could not go fmt: %s", b)
	}
	return nil
}
