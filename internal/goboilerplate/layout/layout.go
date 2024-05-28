package layout

import (
	"fmt"
	"goboilerplate/internal/goboilerplate/layout/customlayout"
	"goboilerplate/internal/goboilerplate/layout/defaultlayout"
	"os"
	"os/exec"
	"strings"
)

var appFolder = "./"
var appName = ""

// CreateLayout creates a layout for the application.
// It prompts the user for the app name and project folder,
// and then initiates the project creation process.
// If verbose is set to true, it prints additional information during the process.
// The function returns an error if any error occurs during the creation process.
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
			var file *os.File
			file, err = os.Open(appFolder + "/layout.yaml")
			if err != nil {
				err = defaultlayout.CreateDefaultLayout(verbose, appName)
			} else {
				err = customlayout.CreateCustomLayout(verbose, appName, file)
			}
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
