package customlayout

import (
	"fmt"
	"goboilerplate/pkg/types"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Folder struct {
	Path       string            `yaml:"path"`
	Subfolders []Folder          `yaml:"subfolders"`
	Component  []types.Component `yaml:"components"`
}

// DecodeYAMLIntoFolders decodes the YAML data from the provided reader into a Folder struct.
// It returns the decoded Folder struct and any error encountered during decoding.
func DecodeYAMLIntoFolders(content io.Reader) (folder Folder, err error) {
	err = yaml.NewDecoder(content).Decode(&folder)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
		return folder, err
	}

	return folder, nil
}

// CreateCustomLayout creates a custom layout for the application.
func CreateCustomLayout(verbose bool, appName string, file *os.File) (err error) {
	var customLayout Folder

	customLayout, err = DecodeYAMLIntoFolders(file)
	if err != nil {
		return err
	}

	checkAndCreateLayout(customLayout)

	return nil
}

func checkAndCreateLayout(folder Folder) (err error) {
	if _, err = os.Stat(folder.Path); os.IsNotExist(err) {
		err = os.Mkdir(folder.Path, os.ModePerm)
	}

	if err != nil {
		return err
	}

	for _, component := range folder.Component {
		if component.CompType != component.Framework.GetType() {
			return fmt.Errorf("component type and framework type do not match: %s, %s", component.CompType, component.Framework.GetType())
		}

		if _, err = os.Stat(component.Path); os.IsNotExist(err) {
			err = os.Mkdir(component.Path, os.ModePerm)
		}

		if err != nil {
			return err
		}

		fwFile, err := os.Create(component.Path + "/" + component.Framework.GetName() + ".go")
		if err != nil {
			return err
		}

		template := component.Framework.GetTemplate()
		err = template.Execute(fwFile, nil)
		if err != nil {
			return err
		}
	}

	for _, subFolder := range folder.Subfolders {
		if err != nil {
			return err
		}

		err = checkAndCreateLayout(subFolder)
	}

	return
}
