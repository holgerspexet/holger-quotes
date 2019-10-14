package create

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/holgerspexet/holger-quotes/storage"
)

func CreateHandler(store storage.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "POST":
			if err := req.ParseForm(); err != nil {
				http.Error(w, "Unable to parse the form", http.StatusUnprocessableEntity)
				return
			}

			store.Store(storage.QuoteInfo{
				Who:   strings.Join(req.Form["who"], ""),
				Quote: strings.Join(req.Form["quote"], ""),
				When:  strings.Join(req.Form["when"], ""),
			})
			http.Redirect(w, req, "/", http.StatusSeeOther)
		case "GET":
			// TODO investigate how these paths works
			tmpl, err := template.ParseFiles("./templates/base.html", "./templates/new.html")
			if err != nil {
				log.Panic(err.Error())
			}
			err = tmpl.Execute(w, nil)
			if err != nil {
				log.Panic(err)
			}
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
