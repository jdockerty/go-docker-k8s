package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
)

type todo struct {
	TaskName string
	Complete bool
}

type tasks struct {
	Day   string
	Todos []todo
}

func runCommand(command, args string) {
	cmd := exec.Command(command, args)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(out.String(), "\n")
}

func createTemplate() {
	myTasks := tasks{
		Day: "Wednesday",
		Todos: []todo{
			{TaskName: "Golang task dashboard.", Complete: true},
			{TaskName: "Eat breakfast.", Complete: false},
			{TaskName: "Finish first driving lesson.", Complete: true},
		},
	}

	template := template.Must(template.ParseFiles(`tasks.html`))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template.Execute(w, myTasks)
	})
	http.ListenAndServe(":8081", nil)
}

func main() {
	//runCommand("ls", "")
	fmt.Println("Starting server...")
	createTemplate()
}
