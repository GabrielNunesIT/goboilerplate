package layout

import (
	"fmt"
	"goboilerplate/internal/goboilerplate/layout/api"
	"goboilerplate/internal/goboilerplate/layout/assets"
	"goboilerplate/internal/goboilerplate/layout/bin"
	"goboilerplate/internal/goboilerplate/layout/build"
	"goboilerplate/internal/goboilerplate/layout/cmd"
	"goboilerplate/internal/goboilerplate/layout/configs"
	"goboilerplate/internal/goboilerplate/layout/docs"
	"goboilerplate/internal/goboilerplate/layout/githooks"
	"goboilerplate/internal/goboilerplate/layout/internal"
	"os"
	"os/exec"
	"strings"
)

var appFolder = "./"
var appName = ""

func CreateLayout(verbose bool) (err error) {
	fmt.Println()
	fmt.Println("What do you want to name the app?")

	_, err = fmt.Scanln(&appName)

	if err == nil && appName != "" {
		fmt.Println()
		fmt.Println("Where do you want to create the project? (Default is in the current folder.)")
		fmt.Println("Press ENTER for default.")

		_, err = fmt.Scanln(&appFolder)
		if err != nil {
			if err.Error() == ("unexpected newline") {
				err = nil
				appFolder = "./"
			}
		}
	}

	if err == nil && appFolder != "" {
		if verbose {
			fmt.Println()
			fmt.Println("Initiating project creation...")
		}

		err = execGoModInit()
		if err == nil {
			err = api.CreateApiFolder(verbose)
		}
		if err == nil {
			err = assets.CreateAssetsFolder(verbose)
		}
		if err == nil {
			err = bin.CreateBinFolder(verbose)
		}
		if err == nil {
			err = build.CreateBuildFolder(verbose)
		}
		if err == nil {
			err = cmd.CreateCmdFolder(verbose, appName)
		}
		if err == nil {
			err = configs.CreateConfigsFolder(verbose)
		}
		if err == nil {
			err = docs.CreateDocsFolder(verbose)
		}
		if err == nil {
			err = githooks.CreateGithooksFolder(verbose)
		}
		if err == nil {
			err = internal.CreateInternalFolder(verbose, appName)
		}
		if err == nil {
			err = writeProjectReadme()
		}

		fmt.Println("\n" + strings.Repeat(string("#"), 50))
		fmt.Println("Project creation finalized!!")
		fmt.Println(strings.Repeat(string("#"), 50) + "\n")
	}

	return err
}

func execGoModInit() (err error) {
	if appFolder != "./" {
		err = os.MkdirAll(appFolder, os.ModeDir)
		if err == nil {
			err = os.Chdir(appFolder)
		}
	}

	if err == nil {
		var goPath string
		goPath, err = exec.LookPath("go")
		if err == nil {
			createCmd := exec.Command(goPath, "mod", "init", appName)
			err = createCmd.Run()
		}
	}

	return err
}

func writeProjectReadme() (err error) {
	var file *os.File
	file, err = os.Create("./README.md")
	if err == nil {
		contents := []string{
			"# `/" + appName + "`\n",
			"\nDetail your project Information.\n",
		}

		for _, content := range contents {
			_, err = file.WriteString(content)
		}
	}

	return err
}
