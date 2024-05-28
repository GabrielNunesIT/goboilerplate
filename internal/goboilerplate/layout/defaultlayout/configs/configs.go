package configs

import (
	"fmt"
	"os"
)

func CreateConfigsFolder(verbose bool) (err error) {
	err = os.MkdirAll("./configs", os.ModeDir)
	if err == nil {
		var file *os.File
		file, err = os.Create("./configs/README.md")
		if err == nil {
			err = writeConfigsReadme(file)
		}
	}

	if err != nil && verbose {
		fmt.Printf("\n Error while creating Configs folder: %s", err)
	}
	return err
}

func writeConfigsReadme(file *os.File) (err error) {
	contents := []string{
		"# `/configs`\n",
		"\nConfiguration file templates or default configs.\n",
		"\nPut your `confd` or `consul-template` template files here.\n",
	}

	for _, content := range contents {
		_, err = file.WriteString(content)
	}

	return err
}
