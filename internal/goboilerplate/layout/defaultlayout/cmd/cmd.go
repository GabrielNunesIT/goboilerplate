package cmd

import (
	"fmt"
	"os"
)

func CreateCmdFolder(verbose bool, appName string) (err error) {
	err = os.MkdirAll("./cmd/"+appName, os.ModeDir)
	if err == nil {
		var file *os.File
		file, err = os.Create("./cmd/README.md")
		if err == nil {
			err = writeCmdReadme(file)
		}
	}

	if err != nil && verbose {
		fmt.Printf("\n Error while creating CMD folder: %s", err)
	}
	return err
}

func writeCmdReadme(file *os.File) (err error) {
	contents := []string{
		"# `/cmd`\n",
		"\nMain applications for this project.\n",
		"\nThe directory name for each application should match the name of the executable you want to have (e.g., `/cmd/myapp`).\n",
		"\nDon't put a lot of code in the application directory. If you think the code can be imported and used in other projects, then it should live in the `/pkg` directory. If the code is not reusable or if you don't want others to reuse it, put that code in the `/internal` directory. You'll be surprised what others will do, so be explicit about your intentions!\n",
		"\nIt's common to have a small `main` function that imports and invokes the code from the `/internal` and `/pkg` directories and nothing else.\n",
	}

	for _, content := range contents {
		_, err = file.WriteString(content)
	}

	return err
}
