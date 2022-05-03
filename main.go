package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

func main() {
	module, _ := tfconfig.LoadModule(".")
	if len(module.RequiredCore) < 1 {
		fmt.Fprintln(os.Stderr, "There is no required version.")
		os.Exit(1)
	}
	versionConstraint := module.RequiredCore[0]

	r := regexp.MustCompile(`\d+\.\d+\.\d+`)
	version := r.FindString(versionConstraint)
	if version == "" {
		fmt.Fprintf(os.Stderr, "Fail to extract version from %s\n", versionConstraint)
	}

	fmt.Printf("%s\n", version)
}
