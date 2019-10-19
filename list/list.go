package list

import (
	"log"
	"net/http"
	"path"
	"text/template"

	"github.com/holgerspexet/holger-quotes/storage"
)

// Handler returns a http.HandlerFunc which creates a list of all quotes
func Handler(store storage.Store, templateDir string, hosting string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		tmpl, err := template.ParseFiles(path.Join(templateDir, "base.html"), path.Join(templateDir, "list.html"))
		if err != nil {
			log.Panic(err.Error())
		}

		err = tmpl.Execute(w, pageData{Quotes: store.Get(), Hosting: hosting, Join: path.Join})
		if err != nil {
			log.Panic(err)
		}
	}
}

type pageData struct {
	Quotes  []storage.QuoteInfo
	Hosting string
	Join    func(elem ...string) string
}
