package list

import (
	"log"
	"net/http"
	"text/template"

	"github.com/holgerspexet/holger-quotes/storage"
)

func ListHandler(store storage.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// TODO investigate how these paths works
		tmpl, err := template.ParseFiles("./templates/base.html", "./templates/list.html")
		if err != nil {
			log.Panic(err.Error())
		}

		err = tmpl.Execute(w, ListPageData{Quotes: store.Get()})
		if err != nil {
			log.Panic(err)
		}
	}
}

type ListPageData struct {
	Quotes []storage.QuoteInfo
}
