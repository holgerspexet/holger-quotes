package create

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/holgerspexet/holger-quotes/storage"
)

// CreateHandler returns a http.HandlerFunc
// which retuns a form for creating quotes on GET
// and creates adds a new quote to the storage system on POST
func CreateHandler(store storage.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "POST":
			if err := req.ParseForm(); err != nil {
				http.Error(w, "Unable to parse the form", http.StatusUnprocessableEntity)
				return
			}

			store.Store(storage.QuoteInfo{
				Quote: strings.Join(req.Form["quote"], ""),
				Who:   strings.Join(req.Form["who"], ""),
				Where: strings.Join(req.Form["where"], ""),
				When:  time.Now(),
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
