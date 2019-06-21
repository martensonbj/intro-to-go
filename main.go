package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// http.HandleFunc("/", home)
	http.HandleFunc("/todos", getAllTodos)
	http.HandleFunc("/add-todo", addTodo)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

type Todo struct {
	Title   string
	Content string
}

type PageVariables struct {
	PageTitle string
	PageTodos []Todo
}

// get request to get all existing todos from the db
var todos []Todo

func addTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo := Todo{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}

	todos = append(todos, todo)
	log.Print(todos)
	http.Redirect(w, r, "/todos", http.StatusSeeOther)

}

func getAllTodos(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		PageTitle: "Intro To Go",
		PageTodos: todos,
	}

	t, err := template.ParseFiles("todos.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template execution error: ", err)
		return
	}

	err = t.Execute(w, pageVariables)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template execution error: ", err)
		return
	}
}
