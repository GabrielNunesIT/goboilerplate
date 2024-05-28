package build

import (
	"fmt"
	"os"
)

func CreateBuildFolder(verbose bool) (err error) {
	err = os.MkdirAll("./build", os.ModeDir)
	if err == nil {
		var file *os.File
		file, err = os.Create("./build/README.md")
		if err == nil {
			err = writeBuildReadme(file)
		}
	}

	if err != nil && verbose {
		fmt.Printf("\n Error while creating Build folder: %s", err)
	}
	return err
}

func writeBuildReadme(file *os.File) (err error) {
	contents := []string{
		"# `/build`\n",
		"\nPackaging and Continuous Integration.\n",
		"\nPut your cloud (AMI), container (Docker), OS (deb, rpm, pkg) package configurations and scripts in the `/build/package` directory.\n",
		"\nPut your CI (travis, circle, drone) configurations and scripts in the `/build/ci` directory. Note that some of the CI tools (e.g., Travis CI) are very picky about the location of their config files. Try putting the config files in the `/build/ci` directory linking them to the location where the CI tools expect them when possible (don't worry if it's not and if keeping those files in the root directory makes your life easier :-)).\n",
	}

	for _, content := range contents {
		_, err = file.WriteString(content)
	}

	return err
}
