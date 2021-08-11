package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world<h1>")
}

func main() {
	fmt.Println("Server Starting...")
	http.HandleFunc("/", index)

	http.ListenAndServe(":3000", nil)
}