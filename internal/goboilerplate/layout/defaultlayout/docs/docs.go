package docs

import (
	"fmt"
	"os"
)

func CreateDocsFolder(verbose bool) (err error) {
	err = os.MkdirAll("./docs", os.ModeDir)
	if err == nil {
		var file *os.File
		file, err = os.Create("./docs/README.md")
		if err == nil {
			err = writeDocsReadme(file)
		}
	}

	if err != nil && verbose {
		fmt.Printf("\n Error while creating Docs folder: %s", err)
	}
	return err
}

func writeDocsReadme(file *os.File) (err error) {
	contents := []string{
		"# `/docs`\n",
		"\nDesign and user documents (in addition to your godoc generated documentation).\n",
	}

	for _, content := range contents {
		_, err = file.WriteString(content)
	}

	return err
}
