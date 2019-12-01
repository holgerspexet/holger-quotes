// +build !release

package templates

//go:generate go run generate_vfsdata.go

import "net/http"

// assets is a http.FileSystem containing the template files
var assets = http.Dir("templates/assets")
