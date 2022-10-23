package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/rs/xid"
)

var templ *template.Template
var port string

type Todo struct {
	Id          string `json:"id"`
	Item        string `json:"item"`
	IsComplited bool   `json:"is_complited"`
}

type PageData struct {
	Title string
	Todos []Todo
}

func getTodoList(w http.ResponseWriter, r *http.Request) {
	var todos []Todo
	path := "./todos.json"

	byteValue, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(byteValue, &todos)
	if err != nil {
		panic(err)
	}

	data := PageData{
		Title: "Test TODO List",
		Todos: todos,
	}

	templ.Execute(w, data)
}

func createNewTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод запрещен!", http.StatusMethodNotAllowed)

		return
	}

	todo := r.FormValue("todo")
	guid := xid.New().String()

	newTodo := Todo{
		Id:          guid,
		Item:        todo,
		IsComplited: false,
	}

	var todos []Todo
	path := "./todos.json"

	byteValue, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(byteValue, &todos)
	if err != nil {
		panic(err)
	}
	todos = append(todos, newTodo)
	fmt.Println(newTodo, todos)

	jsonTodos, err := json.MarshalIndent(&todos, "", "    ")
	if err != nil {
		panic(err)
	}
	os.WriteFile("todos.json", jsonTodos, 0666)
	getTodoList(w, r)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	todoId := r.FormValue("id")

	var todos []Todo
	path := "./todos.json"

	byteValue, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(byteValue, &todos)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(todos); i++ {
		if todos[i].Id == todoId {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}

	jsonTodos, err := json.MarshalIndent(&todos, "", "    ")
	if err != nil {
		panic(err)
	}
	os.WriteFile("todos.json", jsonTodos, 0666)
	getTodoList(w, r)
}

func changeStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	r.ParseMultipartForm(0)
	todoId := r.FormValue("id")

	fmt.Println(todoId)
}

func main() {
	mux := http.NewServeMux()
	templ = template.Must(template.ParseFiles("templates/index.gohtml"))
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", getTodoList)
	mux.HandleFunc("/todo/create", createNewTodo)
	mux.HandleFunc("/todo/delete", deleteTodo)
	mux.HandleFunc("/todo/status", changeStatus)
	port = ":3001"

	log.Fatal(http.ListenAndServe(port, mux))
}
