package create

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/holgerspexet/holger-quotes/storage"
	"github.com/holgerspexet/holger-quotes/templates"
)

// Handler returns a http.HandlerFunc
// which retuns a form for creating quotes on GET
// and creates adds a new quote to the storage system on POST
func Handler(store storage.Store, hosting string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "POST":
			post(w, req, store, hosting)
		case "GET":
			get(w, hosting)
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

	newQuote := storage.QuoteInfo{
		Quote: strings.Join(req.Form["quote"], ""),
		Who:   strings.Join(req.Form["who"], ""),
		Where: strings.Join(req.Form["where"], ""),
		When:  time.Now(),
	}
	store.Store(newQuote)
	log.Printf("New quote: %+v", newQuote)
	http.Redirect(w, req, hosting, http.StatusSeeOther)
}

func get(w http.ResponseWriter, hosting string) {
	templates.ExecuteNewTemplate(w, hosting)
}
