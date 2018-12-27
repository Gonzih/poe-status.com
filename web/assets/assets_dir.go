// +build dev

package assets

import "net/http"

// AssetsDir is http dir var, used with conjunction with vfsgen
var AssetsDir = http.Dir("../../web/assets/assets")
