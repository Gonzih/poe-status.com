// +build ignore

package main

import (
	"github.com/shurcooL/vfsgen"
	"gitlab.com/Gonzih/poe-status.com/config"
	"gitlab.com/Gonzih/poe-status.com/util"
)

func main() {
	err := vfsgen.Generate(config.DataDir, vfsgen.Options{
		PackageName:  "config",
		BuildTags:    "!dev",
		VariableName: "DataDir",
	})
	util.Must(err)
}
