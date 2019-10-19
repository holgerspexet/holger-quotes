package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/holgerspexet/holger-quotes/config"
	"github.com/holgerspexet/holger-quotes/create"
	"github.com/holgerspexet/holger-quotes/list"
	"github.com/holgerspexet/holger-quotes/storage"
)

func main() {
	config := config.LoadConfig()
	muxer := http.NewServeMux()
	store := storage.CreateStorage(config)

	base := config.Hosting
	new := path.Join(base, "ny")
	static := path.Join(base, "static") + "/"

	staticServer := http.StripPrefix(static, http.FileServer(http.Dir(config.StaticDir)))
	muxer.HandleFunc(base, list.Handler(store, config.TemplateDir, config.Hosting))
	muxer.HandleFunc(new, create.Handler(store, config.TemplateDir, config.Hosting))
	muxer.Handle(static, staticServer)

	log.Printf("Server running at: %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), muxer))
}
