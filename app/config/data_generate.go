// +build ignore

package main

import (
	"github.com/shurcooL/vfsgen"
	"github.com/Gonzih/poe-status.com/app/config"
	"github.com/Gonzih/poe-status.com/util"
)

func main() {
	err := vfsgen.Generate(config.DataDir, vfsgen.Options{
		PackageName:  "config",
		BuildTags:    "!dev",
		VariableName: "DataDir",
	})
	util.Must(err)
}
