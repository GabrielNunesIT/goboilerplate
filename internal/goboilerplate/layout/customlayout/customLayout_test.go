package customlayout_test

import (
	"goboilerplate/internal/goboilerplate/layout/customlayout"
	"strings"
	"testing"

	"github.com/k0kubun/pp"
)

func TestUnmarshalYAMLFile(t *testing.T) {
	fileContents := `
path: ./
subfolders:
  [
	{
	  path: ./cmd,
	  subfolders:
		[
		  {
			path: ./cmd/batatas,
			components:
			  [
				{ 
				  name: main,
				  path: ./cmd/batatas,
				  type: main,
				},
			  ],
		  },
		],
	},
	{
	  path: ./internal,
	  subfolders:
		[
		  {
			path: ./internal/app/batatas,
			components:
			  [
				{
				  name: echo,
				  path: ./internal/app/batatas/webserver,
				  type: webserver,
				  framework: echo,
				},
				{
				  name: log,
				  path: ./internal/app/batatas/log,
				  type: log,
				  framework: zerolog,
				},
				{
				  name: db,
				  path: ./internal/app/batatas/db,
				  type: db,
				  framework: gorm,
				},
			  ],
		  },
		],
	},
  ]`

	folders, err := customlayout.DecodeYAMLIntoFolders(strings.NewReader(fileContents))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	pp.Println(folders)
}
