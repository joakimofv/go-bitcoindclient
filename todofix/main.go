package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"

	"sigs.k8s.io/yaml"
)

func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	if err := run(); err != nil {
		log.SetFlags(0)
		log.Fatalln(err)
	}
}

type YamlEntry struct {
	File string `json:"file"`
	Todo string `json:"todo"`
	Fix  string `json:"fix"`
}

func run() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("Usage: %s <yaml-file>", os.Args[0])
	}
	return Run(os.Args[1])
}

func Run(yamlFile string) error {
	b, err := os.ReadFile(yamlFile)
	if err != nil {
		return err
	}

	var todos []YamlEntry
	if err := yaml.UnmarshalStrict(b, &todos); err != nil {
		return err
	}
	var todosPerFile [][]YamlEntry
OUTER:
	for _, todo := range todos {
		for i := range todosPerFile {
			if todo.File == todosPerFile[i][0].File {
				todosPerFile[i] = append(todosPerFile[i], todo)
				continue OUTER
			}
		}
		todosPerFile = append(todosPerFile, []YamlEntry{todo})
	}

	var updatedFiles []fileContents
	for _, group := range todosPerFile {
		b, err := os.ReadFile(group[0].File)
		if err != nil {
			return err
		}
		for _, todo := range group {
			re, err := regexp.Compile(todo.Todo)
			if err != nil {
				return err
			}
			mm := re.FindAllSubmatch(b, -1)
			if len(mm) == 0 {
				return fmt.Errorf("Regexp '%s' failed to match in file '%s'", todo.Todo, todo.File)
			}
			for _, m := range mm {
				var fmtArgs []interface{}
				for _, subMatch := range m[1:] {
					fmtArgs = append(fmtArgs, subMatch)
				}
				fix := fmt.Sprintf(todo.Fix, fmtArgs...)
				b = bytes.Replace(b, m[0], []byte(fix), 1)
			}
		}
		updatedFiles = append(updatedFiles, fileContents{file: group[0].File, b: b})
	}

	for _, uf := range updatedFiles {
		if err := os.WriteFile(uf.file, uf.b, 0777); err != nil {
			return err
		}
	}
	return nil
}

type fileContents struct {
	file string
	b    []byte
}
