package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	muxer := http.NewServeMux()

	staticServer := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	muxer.HandleFunc("/", listHandler)
	muxer.Handle("/static/", staticServer)

	log.Printf("Server running at: %d", 8080)
	log.Fatal(http.ListenAndServe(":8080", muxer))
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
