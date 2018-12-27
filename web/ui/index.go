//go:generate go run -tags=dev assets_generate.go

package ui

import (
	"net/http"
	// "gitlab.com/Gonzih/poe-status.com/db"
)

// AssetsHandler will serve static assets
var AssetsHandler = http.FileServer(AssetsDir)
