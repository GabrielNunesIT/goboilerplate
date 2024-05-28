package internalproject

import (
	"fmt"
	"goboilerplate/internal/goboilerplate/layout/defaultlayout/internalproject/config"
	"goboilerplate/internal/goboilerplate/layout/defaultlayout/internalproject/log"
	"os"
)

func CreateInternalFolder(verbose bool, appName string) (err error) {
	err = os.MkdirAll("./internal/pkg", os.ModeDir)
	if err == nil {
		err = os.MkdirAll("./internal/app/"+appName, os.ModeDir)
	}
	if err == nil {
		var file *os.File
		file, err = os.Create("./internal/README.md")
		if err == nil {
			err = writeInternalReadme(file)
		}
	}

	if err == nil {
		err = config.CreateAppConfigFiles(appName)
	}
	if err == nil {
		err = log.CreatePkgLogFiles()
	}

	if err != nil && verbose {
		fmt.Printf("\n Error while creating Internal folder: %s", err)
	}
	return err
}

func writeInternalReadme(file *os.File) (err error) {
	contents := []string{
		"# `/internal`\n",
		"\nPrivate application and library code. This is the code you don't want others importing in their applications or libraries. Note that this layout pattern is enforced by the Go compiler itself. See the Go 1.4 [`release notes`](https://golang.org/doc/go1.4#internalpackages) for more details. Note that you are not limited to the top level `internal` directory. You can have more than one `internal` directory at any level of your project tree.\n",
		"\nYou can optionally add a bit of extra structure to your internal packages to separate your shared and non-shared internal code. It's not required (especially for smaller projects), but it's nice to have visual clues showing the intended package use. Your actual application code can go in the `/internal/app` directory (e.g., `/internal/app/myapp`) and the code shared by those apps in the `/internal/pkg` directory (e.g., `/internal/pkg/myprivlib`).\n",
	}

	for _, content := range contents {
		_, err = file.WriteString(content)
	}

	return err
}
