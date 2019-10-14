package create

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/holgerspexet/holger-quotes/storage"
)

func CreateHandler(storage.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "POST":
			if err := req.ParseForm(); err != nil {
				http.Error(w, "Unable to parse the form", http.StatusUnprocessableEntity)
				return
			}

			fmt.Printf("%s", req.Form)
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
