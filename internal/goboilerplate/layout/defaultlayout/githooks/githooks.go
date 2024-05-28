package githooks

import (
	"fmt"
	"os"
	"strings"
)

func CreateGithooksFolder(verbose bool) (err error) {
	var githooksFolder = ""

	fmt.Println()
	fmt.Println("Will your project contain Githooks? [y/N]")

	_, err = fmt.Scan(&githooksFolder)
	githooksFolder = strings.ToLower(githooksFolder)

	if githooksFolder != "y" && githooksFolder != "n" {
		if verbose {
			fmt.Printf(`\n Bad Input: %s. Considering "N" `, githooksFolder)
		}
		githooksFolder = "N"
	}

	if githooksFolder == "y" {
		err = os.MkdirAll("./githooks", os.ModeDir)
		if err == nil {
			var file *os.File
			file, err = os.Create("./githooks/README.md")
			if err == nil {
				err = writeGithooksReadme(file)
			}
		}
	}

	if err != nil && verbose {
		fmt.Printf("\n Error while creating Githooks folder: %s", err)
	}
	return err
}

func writeGithooksReadme(file *os.File) (err error) {
	contents := []string{
		"# `/Githooks`\n",
		"\nGit hooks.\n",
	}

	for _, content := range contents {
		_, err = file.WriteString(content)
	}

	return err
}
