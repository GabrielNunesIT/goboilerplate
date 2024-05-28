package bin

import (
	"fmt"
	"os"
)

func CreateBinFolder(verbose bool) (err error) {
	err = os.MkdirAll("./bin", os.ModeDir)
	if err == nil {
		var file *os.File
		file, err = os.Create("./bin/README.md")
		if err == nil {
			err = writeBinReadme(file)
		}
	}

	if err != nil && verbose {
		fmt.Printf("\n Error while creating BIN folder: %s", err)
	}
	return err
}

func writeBinReadme(file *os.File) (err error) {
	contents := []string{
		"# `/bin`\n",
		"\nBinaries files for different operating systems and architectures.\n",
	}

	for _, content := range contents {
		_, err = file.WriteString(content)
	}

	return err
}
