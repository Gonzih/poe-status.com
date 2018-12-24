// +build dev

package config

import "net/http"

var DataDir http.FileSystem = http.Dir("data")
