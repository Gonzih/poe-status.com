// +build ignore

package main

import (
	"github.com/shurcooL/vfsgen"
	"github.com/Gonzih/poe-status.com/util"
	"github.com/Gonzih/poe-status.com/web/assets"
)

func main() {
	err := vfsgen.Generate(assets.AssetsDir, vfsgen.Options{
		PackageName:  "assets",
		BuildTags:    "!dev",
		VariableName: "AssetsDir",
	})
	util.Must(err)
}
