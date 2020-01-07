package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	//"html/template"
	"path"
	"log"
)

func serveWebpage() {
	router := mux.NewRouter()
	router.HandleFunc("/tasks", serveHome)

	log.Fatal(http.ListenAndServe(":8080", router))

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	filepath := path.Dir(`C:\Users\Jack\Desktop\Go\src\taskschedulerdashboard\html\tasks.html`)
    w.Header().Set("Content-type", "text/html")
    http.ServeFile(w, r, filepath)
}

func main() {
	fmt.Print("Starting server...")
	serveWebpage()

}
