package list

import (
	"net/http"

	"github.com/holgerspexet/holger-quotes/storage"
	"github.com/holgerspexet/holger-quotes/templates"
)

// Handler returns a http.HandlerFunc which creates a list of all quotes
func Handler(store storage.Store, hosting string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		templates.ExecuteListTemplate(w, hosting, store.Get())
	}
}
