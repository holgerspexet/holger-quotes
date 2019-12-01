// +build !release

package static

//go:generate go run generate_vfsdata.go

import "net/http"

// Assets is a http.FileSystem containing the static files
var Assets http.FileSystem = http.Dir("static/assets")
