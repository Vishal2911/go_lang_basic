package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello world</h1>")
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>its vishal</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("server started")
	http.ListenAndServe(":3000", nil)
}
