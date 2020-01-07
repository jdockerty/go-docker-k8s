package main

import (
	"fmt"
	"net/http"
)

func serveWebpage() {
	http.Handle("/tasks", http.StripPrefix("/tasks.html", http.FileServer(http.Dir("C:\\Users\\Jack\\Desktop\\Go\\src\\taskschedulerdashboard\\html\\tasks.html"))))
	http.ListenAndServe(":3000", nil)

}

func main() {
	fmt.Print("Starting server...")
	serveWebpage()

}
