package main

import (
	"fmt"
	"html/template"
	"net/http"
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
	myTasks := tasks{
		Day: "Wednesday",
		Todos: []todo{
			{TaskName: "Golang task dashboard.", Complete: false},
			{TaskName: "Eat breakfast.", Complete: true},
			{TaskName: "Finish first driving lesson.", Complete: true},
		},
	}
	template := template.Must(template.ParseFiles(`html\tasks.html`))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template.Execute(w, myTasks)
	})
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Starting server...")
	createTemplate()
}
