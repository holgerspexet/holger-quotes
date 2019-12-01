// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	err:= vfsgen.Generate(http.Dir("./assets"), vfsgen.Options{
		PackageName:  "templates",
		BuildTags:    "release",
		VariableName: "assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
