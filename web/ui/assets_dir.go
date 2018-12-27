// +build dev

package ui

import "net/http"

// AssetsDir is http dir var, used with conjunction with vfsgen
var AssetsDir = http.Dir("../../web/ui/assets")
