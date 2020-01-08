package main

import (
	"fmt"
	"net/http"

	//"github.com/gorilla/mux"
	"html/template"
)

type todo struct {
	TaskName string
	Complete bool
}

type tasks struct {
	Day   string
	Todos []todo
}

func createTemplate() {
	template := template.Must(template.ParseFiles(`html\tasks.html`))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		myTasks := tasks{
			Day: "Wednesday",
			Todos: []todo{
				{TaskName: "Golang task dashboard.", Complete: false},
				{TaskName: "Eat breakfast.", Complete: true},
				{TaskName: "Finish first driving lesson.", Complete: true},
			},
		}
		template.Execute(w, myTasks)
	})

	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Starting server...")
	createTemplate()

}
