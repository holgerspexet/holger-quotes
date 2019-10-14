package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	muxer := http.NewServeMux()

	staticServer := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	muxer.HandleFunc("/", listHandler)
	muxer.HandleFunc("/ny", createHandler)
	muxer.Handle("/static/", staticServer)

	log.Printf("Server running at: %d", 8080)
	log.Fatal(http.ListenAndServe(":8080", muxer))
}

func createHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			http.Error(w, "Unable to parse the form", http.StatusUnprocessableEntity)
			return
		}

		fmt.Printf("%s", req.Form)
		http.Redirect(w, req, "/", http.StatusSeeOther)
	case "GET":
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

func listHandler(w http.ResponseWriter, req *http.Request) {
	// fmt.Fprint(w, "Hello World")
	tmpl, err := template.ParseFiles("./templates/base.html", "./templates/list.html")
	if err != nil {
		log.Panic(err.Error())
	}

	quotes := []QuoteInfo{
		{CreatedAt: "Igår", CreatedBy: "Ingen", Quote: "Hejsan", Quoted: "Johan"},
		{CreatedAt: "Idag", CreatedBy: "Jag", Quote: "Det var en gång", Quoted: "Astrid"},
		{CreatedAt: "Imon", CreatedBy: "Någon annan", Quote: "All makt åt Tengil", Quoted: "Jonatan"},
	}

	err = tmpl.Execute(w, ListPageData{Quotes: quotes})
	if err != nil {
		log.Panic(err)
	}
}

type ListPageData struct {
	Quotes []QuoteInfo
}

type QuoteInfo struct {
	CreatedAt string
	CreatedBy string
	Quote     string
	Quoted    string
}
