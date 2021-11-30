package main

import (
	"PortofolioWeb/Handler"
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ini adalah index!")
}

func main() {

	http.HandleFunc("/input", Handler.Input)
	http.HandleFunc("/result", Handler.Result)
	http.HandleFunc("/", root)

	fileserver := http.FileServer(http.Dir("assets"))
	http.Handle("/static/", http.StripPrefix("/static", fileserver))

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
