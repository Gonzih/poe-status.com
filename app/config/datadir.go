// +build dev

package config

import "net/http"

// DataDir is http dir var, used with conjunction with vfsgen
var DataDir http.FileSystem = http.Dir("../../app/config/data")
