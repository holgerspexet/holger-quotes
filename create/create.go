package create

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/holgerspexet/holger-quotes/storage"
)

// Handler returns a http.HandlerFunc
// which retuns a form for creating quotes on GET
// and creates adds a new quote to the storage system on POST
func Handler(store storage.Store, templateDir string, hosting string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "POST":
			post(w, req, store, hosting)
		case "GET":
			get(w, templateDir, hosting)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func post(w http.ResponseWriter, req *http.Request, store storage.Store, hosting string) {
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
	http.Redirect(w, req, hosting, http.StatusSeeOther)
}

func get(w http.ResponseWriter, templateDir string, hosting string) {
	tmpl, err := template.ParseFiles(path.Join(templateDir, "base.html"), path.Join(templateDir, "new.html"))

	if err != nil {
		log.Panic(err.Error())
	}
	err = tmpl.Execute(w, pageData{Hosting: hosting, Join: path.Join})
	if err != nil {
		log.Panic(err)
	}
}

type pageData struct {
	Hosting string
	Join    func(elem ...string) string
}
