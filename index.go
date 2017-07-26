package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func main() {
    http.HandleFunc("/", handler)
	http.HandleFunc("/home", homeHandler)
    http.ListenAndServe(":8080", nil)
}