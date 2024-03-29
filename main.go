package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/holgerspexet/holger-quotes/config"
	"github.com/holgerspexet/holger-quotes/create"
	"github.com/holgerspexet/holger-quotes/list"
	"github.com/holgerspexet/holger-quotes/static"
	"github.com/holgerspexet/holger-quotes/storage"
)

// This should be overridden with the build flag '-ldflags "-X main.version=<version>"'
var version = "unset"

func main() {
	versionFlag := flag.Bool("version", false, "Prints binary version")
	flag.Parse()
	if *versionFlag {
		fmt.Print(version)
		return
	}

	startServer()
}

func requestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(
			"%s %s",
			r.Method,
			r.RequestURI,
		)
		targetMux.ServeHTTP(w, r)
	})
}

func startServer() {
	config := config.LoadConfig()
	muxer := http.NewServeMux()
	store := storage.CreateStorage(config)

	basePath := config.Hosting
	newPath := path.Join(basePath, "ny")
	staticPath := path.Join(basePath, "static") + "/"

	muxer.HandleFunc(basePath, list.Handler(store, config.Hosting))
	muxer.HandleFunc(newPath, create.Handler(store, config.Hosting))
	muxer.Handle(staticPath, http.StripPrefix(staticPath, http.FileServer(static.Assets)))

	log.Printf("Server running at: %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), requestLogger(muxer)))
}
