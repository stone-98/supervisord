//go:build !release
// +build !release

package main_

import (
	"net/http"
)

//HTTP auto generated
var HTTP http.FileSystem = http.Dir("./webgui")
