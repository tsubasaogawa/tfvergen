package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"testing"
	"time"
)

var testTfFile string

const tfTemplate = `
terraform {
	required_version = "%s"
  } 
`

func TestMain(m *testing.M) {
	if existsTfFile() {
		fmt.Fprintln(os.Stderr, "Please (re)move *.tf file(s) in your current directory.")
		os.Exit(1)
	}

	testTfFile = fmt.Sprintf("%d.tf", time.Now().UnixNano())

	code := m.Run()
	os.Exit(code)
}

func existsTfFile() bool {
	paths, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	result := false
	filepath.Walk(paths, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(".tf$", f.Name())
			if err == nil && r {
				result = true
				return nil
			}
		}
		return nil
	})
	return result
}

func TestPrintRequiredVersion(t *testing.T) {
	tests := map[string]struct {
		ver_string string
		want       string
	}{
		"equal":            {ver_string: "= 1.1.0", want: "1.1.0"},
		"no_operator":      {ver_string: "1.1.0", want: "1.1.0"},
		"rightmost":        {ver_string: "~> 1.1.0", want: "1.1.0"},
		"two_conditions_1": {ver_string: ">= 1.0.0, < 1.1.0", want: "1.0.0"},
		"two_conditions_2": {ver_string: "< 1.1.0, >= 1.0.0", want: "1.1.0"},
	}

	fp, err := os.Create(testTfFile)
	if err != nil {
		t.Fatal(err)
	}
	defer fp.Close()
	defer os.Remove(testTfFile)

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			_, err = fp.Seek(0, 0)
			if err != nil {
				t.Fatal(err)
			}
			fp.WriteString(fmt.Sprintf(tfTemplate, tt.ver_string))

			ver := getRequiredVersion()
			if ver != tt.want {
				t.Errorf("Got: %s, Want: %s", ver, tt.want)
			}
		})
	}
}
