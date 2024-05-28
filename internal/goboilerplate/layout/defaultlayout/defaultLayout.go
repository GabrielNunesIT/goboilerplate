package defaultlayout

import (
	"goboilerplate/internal/goboilerplate/layout/defaultlayout/api"
	"goboilerplate/internal/goboilerplate/layout/defaultlayout/assets"
	"goboilerplate/internal/goboilerplate/layout/defaultlayout/bin"
	"goboilerplate/internal/goboilerplate/layout/defaultlayout/build"
	"goboilerplate/internal/goboilerplate/layout/defaultlayout/cmd"
	"goboilerplate/internal/goboilerplate/layout/defaultlayout/configs"
	"goboilerplate/internal/goboilerplate/layout/defaultlayout/docs"
	"goboilerplate/internal/goboilerplate/layout/defaultlayout/githooks"
	"goboilerplate/internal/goboilerplate/layout/defaultlayout/internalproject"
)

// CreateDefaultLayout creates the default layout for the application.
// It creates various folders required for the application, such as API, assets, bin, build, cmd, configs, docs, githooks, and internal.
// The function takes a boolean parameter 'verbose' to enable verbose logging and a string parameter 'appName' to specify the application name.
// It returns an error if any of the folder creation operations fail, otherwise it returns nil.
// It is based in https://github.com/golang-standards/project-layout.
func CreateDefaultLayout(verbose bool, appName string) (err error) {
	err = api.CreateApiFolder(verbose)
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
		err = internalproject.CreateInternalFolder(verbose, appName)
	}

	return err
}
