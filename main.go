package main

import (
	"log"
	"net/http"

	"github.com/holgerspexet/holger-quotes/create"
	"github.com/holgerspexet/holger-quotes/list"
	"github.com/holgerspexet/holger-quotes/storage"
)

var store storage.Store

func main() {
	muxer := http.NewServeMux()
	store = storage.NewMemoryStorage()

	staticServer := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	muxer.HandleFunc("/", list.ListHandler(store))
	muxer.HandleFunc("/ny", create.CreateHandler(store))
	muxer.Handle("/static/", staticServer)

	log.Printf("Server running at: %d", 8080)
	log.Fatal(http.ListenAndServe(":8080", muxer))
}
