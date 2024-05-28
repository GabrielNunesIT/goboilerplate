package api

import (
	"fmt"
	"os"
	"strings"
)

func CreateApiFolder(verbose bool) (err error) {
	var apiFolder = ""

	fmt.Println()
	fmt.Println("Will your project include an API? [y/N]")

	_, err = fmt.Scan(&apiFolder)
	apiFolder = strings.ToLower(apiFolder)

	if apiFolder != "y" && apiFolder != "n" {
		if verbose {
			fmt.Printf(`\n Bad Input: %s. Considering "N" `, apiFolder)
		}
		apiFolder = "N"
	}

	if apiFolder == "y" {
		err = os.MkdirAll("./api", os.ModeDir)
		if err == nil {
			var file *os.File
			file, err = os.Create("./api/README.md")
			if err == nil {
				err = writeApiReadme(file)
			}
		}
	}

	if err != nil && verbose {
		fmt.Printf("\n Error while creating API folder: %s", err)
	}
	return err
}

func writeApiReadme(file *os.File) (err error) {
	contents := []string{
		"# `/api`\n",
		"\nOpenAPI/Swagger specs, JSON schema files, protocol definition files.\n",
	}

	for _, content := range contents {
		_, err = file.WriteString(content)
	}

	return err
}
