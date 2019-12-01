package templates

import (
	"io"
	"io/ioutil"
	"log"
	"path"
	"text/template"

	"github.com/holgerspexet/holger-quotes/storage"
)

// ExecuteNewTemplate writes the executed template to the provided writer
func ExecuteNewTemplate(wr io.Writer, hosting string) {
	tmpl, err := template.New("new").Parse(readAssets("/base.html", "/new.html"))
	must(err)

	err = tmpl.Execute(
		wr,
		templateHelper{
			Hosting: hosting,
			Join:    path.Join,
		})
	must(err)
}

// ExecuteListTemplate writes the executed template to the provided writer
func ExecuteListTemplate(wr io.Writer, hosting string, quotes []storage.QuoteInfo) {
	tmpl, err := template.New("list").Parse(readAssets("/base.html", "/list.html"))
	must(err)

	err = tmpl.Execute(
		wr,
		struct {
			Quotes []storage.QuoteInfo
			templateHelper
		}{
			Quotes: quotes,
			templateHelper: templateHelper{
				Hosting: hosting,
				Join:    path.Join,
			},
		})
	must(err)
}

func readAssets(paths ...string) string {
	result := ""

	for _, path := range paths {
		file, err := assets.Open(path)
		must(err)

		content, err := ioutil.ReadAll(file)
		must(err)

		result += string(content)
	}

	return result
}

func must(err error) {
	if err != nil {
		log.Panic(err)
	}
}

type templateHelper struct {
	Hosting string
	Join    func(elem ...string) string
}
