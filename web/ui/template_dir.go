// +build dev

package ui

import "net/http"

// TemplateDir is http dir var, used with conjunction with vfsgen
var TemplateDir http.FileSystem = http.Dir("../../web/ui/templates")
