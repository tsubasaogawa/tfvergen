// main package is the main of tfvergen.
package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

const VERSION string = "0.0.1"

// getRequiredVersion returns required_version from tf files in the cwd.
func getRequiredVersion() (string, error) {
	module, _ := tfconfig.LoadModule(".")
	if len(module.RequiredCore) < 1 {
		return "", fmt.Errorf("tfvergen %s\nThere is no required version.", VERSION)
	}
	constraint := module.RequiredCore[0]

	r := regexp.MustCompile(`\d+\.\d+\.\d+`)
	version := r.FindString(constraint)
	if version == "" {
		fmt.Fprintf(os.Stderr, "Fail to extract version from %s\n", constraint)
	}

	return version, nil
}

func main() {
	version, err := getRequiredVersion()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", version)
}
