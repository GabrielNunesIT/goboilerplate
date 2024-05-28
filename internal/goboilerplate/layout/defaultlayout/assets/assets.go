package assets

import (
	"fmt"
	"os"
	"strings"
)

func CreateAssetsFolder(verbose bool) (err error) {
	var assetsFolder = ""

	fmt.Println()
	fmt.Println("Will your project include assets like images, logos, etc... ? [y/N]")

	_, err = fmt.Scan(&assetsFolder)
	assetsFolder = strings.ToLower(assetsFolder)

	if assetsFolder != "y" && assetsFolder != "n" {
		if verbose {
			fmt.Printf(`\n Bad Input: %s. Considering "N" `, assetsFolder)
		}
		assetsFolder = "N"
	}

	if assetsFolder == "y" {
		err = os.MkdirAll("./assets", os.ModeDir)
		if err == nil {
			var file *os.File
			file, err = os.Create("./assets/README.md")
			if err == nil {
				err = writeAssetsReadme(file)
			}
		}
	}

	if err != nil && verbose {
		fmt.Printf("\n Error while creating Assets folder: %s", err)
	}
	return err
}

func writeAssetsReadme(file *os.File) (err error) {
	contents := []string{
		"# `/Assets`\n",
		"Other assets to go along with your repository (images, logos, etc).\n",
	}

	for _, content := range contents {
		_, err = file.WriteString(content)
	}

	return err
}
