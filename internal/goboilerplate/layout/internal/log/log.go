package log

import (
	_ "embed"
	"html/template"
	"os"
)

// Embed the template file
//
//go:embed templates/logContent.tpl
var logContent string

func CreatePkgLogFiles() (err error) {
	var logFile *os.File

	err = os.MkdirAll("./internal/pkg/utils/log", os.ModeDir)

	if err == nil {
		logFile, err = os.Create("./internal/pkg/utils/log/log.go")
		if err == nil {
			var configTmpl *template.Template
			configTmpl, err = template.New("config").Parse(logContent)
			if err == nil {
				err = configTmpl.Execute(logFile, nil)
			}
		}
	}

	return err
}
