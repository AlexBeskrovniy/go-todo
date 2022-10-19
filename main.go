package main

import (
	"log"
	"net/http"
	"text/template"
)

var templ *template.Template
var port string

type Todo struct {
	Item        string
	IsComplited bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func getTodoList(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Test TODO List",
		Todos: []Todo{
			{Item: "Create HTTP Server", IsComplited: true},
			{Item: "Pass template to the browser", IsComplited: false},
			{Item: "Create new Todo without database", IsComplited: false},
		},
	}

	templ.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	templ = template.Must(template.ParseFiles("templates/index.gohtml"))
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", getTodoList)
	port = ":3001"

	log.Fatal(http.ListenAndServe(port, mux))
}
