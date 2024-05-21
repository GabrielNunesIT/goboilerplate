package config

import (
	_ "embed"
	"os"
	"text/template"
)

// Embed the template file
//
//go:embed templates/configContent.tpl
var configContent string

//go:embed templates/loadFromFile.tpl
var loadFromFileContent string

//go:embed templates/loadFromEnvs.tpl
var loadFromEnvsContent string

//go:embed templates/loadFromFlags.tpl
var loadFromFlagContent string

func CreateAppConfigFiles(appName string) (err error) {
	var fileConfig *os.File
	var fileLoadFromFile *os.File
	var fileLoadEnv *os.File
	var fileLoadFlag *os.File

	err = os.MkdirAll("./internal/app/"+appName+"/config", os.ModeDir)

	if err == nil {
		fileConfig, err = os.Create("./internal/app/" + appName + "/config/config.go")
		if err == nil {
			var configTmpl *template.Template
			configTmpl, err = template.New("config").Parse(configContent)
			if err == nil {
				err = configTmpl.Execute(fileConfig, appName)
			}
			// _, err = fileConfig.WriteString(fmt.Sprintf(configContent, appName))
		}
	}

	if err == nil {
		fileLoadFromFile, err = os.Create("./internal/app/" + appName + "/config/loadFromFile.go")
		if err == nil {
			var loadFromFileTmpl *template.Template
			loadFromFileTmpl, err = template.New("loadFromFile").Parse(loadFromFileContent)
			if err == nil {
				err = loadFromFileTmpl.Execute(fileLoadFromFile, appName)
			}
			// _, err = fileLoadFile.WriteString(loadFromFileContent)
		}
	}

	if err == nil {
		fileLoadEnv, err = os.Create("./internal/app/" + appName + "/config/loadFromEnv.go")
		if err == nil {
			var loadFromEnvTmpl *template.Template
			loadFromEnvTmpl, err = template.New("loadFromEnv").Parse(loadFromEnvsContent)
			if err == nil {
				err = loadFromEnvTmpl.Execute(fileLoadEnv, appName)
			}
			// _, err = fileLoadEnv.WriteString(fmt.Sprintf(loadFromEnvContent, strings.ToUpper(appName)))
		}
	}

	if err == nil {
		fileLoadFlag, err = os.Create("./internal/app/" + appName + "/config/loadFromFlag.go")
		if err == nil {
			var loadFromFlagTmpl *template.Template
			loadFromFlagTmpl, err = template.New("loadFromFlag").Parse(loadFromFlagContent)
			if err == nil {
				err = loadFromFlagTmpl.Execute(fileLoadFlag, appName)
			}
			// _, err = fileLoadFlag.WriteString(loadFromFlagContent)
		}
	}

	return err
}
