package main

import (
	"bytes"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	if err := os.WriteFile("abc.yaml", []byte(`---
- file: hej
  todo: "[0-5]"
  fix: X
- file: hej2
  todo: 12([0-9])45
  fix: hello there %s!
`), 0664); err != nil {
		t.Fatal(err)
	}
	defer os.Remove("abc.yaml")
	if err := os.WriteFile("hej", []byte(`123456789
abcde123456
`), 0664); err != nil {
		t.Fatal(err)
	}
	defer os.Remove("hej")
	if err := os.WriteFile("hej2", []byte(`12A451224512345124451254512945
121451224512345124451254512945
`), 0664); err != nil {
		t.Fatal(err)
	}
	defer os.Remove("hej2")

	if err := Run("abc.yaml"); err != nil {
		t.Fatal(err)
	}

	b, err := os.ReadFile("hej")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, []byte(`XXXXX6789
abcdeXXXXX6
`)) {
		t.Errorf("%s", b)
	}
	b, err = os.ReadFile("hej2")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, []byte(`12A45hello there 2!hello there 3!hello there 4!hello there 5!hello there 9!
hello there 1!hello there 2!hello there 3!hello there 4!hello there 5!hello there 9!
`)) {
		t.Errorf("%s", b)
	}
}
