package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/tox/", http.StripPrefix("/tox/", http.FileServer(http.Dir("tox"))))
	port := os.Getenv("PORT")
	if port == "" {
		port = "80" // Default port if not specified
	}
	http.ListenAndServe(":"+port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Irene")
	name := Name{"Irene"}
	template, _ := template.ParseFiles("index.html")
	template.Execute(w, name)
}

type Name struct {
	FName string
}
