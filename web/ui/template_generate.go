// +build ignore

package main

import (
	"github.com/shurcooL/vfsgen"
	"gitlab.com/Gonzih/poe-status.com/util"
	"gitlab.com/Gonzih/poe-status.com/web/ui"
)

func main() {
	err := vfsgen.Generate(ui.TemplateDir, vfsgen.Options{
		PackageName:  "ui",
		BuildTags:    "!dev",
		VariableName: "TemplateDir",
	})
	util.Must(err)
}
