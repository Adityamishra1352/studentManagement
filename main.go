package main

import (
	"fmt"
	"net/http"
)

func operations(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "<h1>Hello I am default</h1>")
	case "/delete":
		fmt.Fprint(w, "<h1>Delete Student</h1>")
	case "/add":
		fmt.Fprint(w, "<h1>Add Student</h1>")
	case "/update":
		fmt.Fprint(w, "<h1>Update</h1>")
	default:
		fmt.Fprint(w, "Error")
	}
}
func main() {
	http.HandleFunc("/", operations)
	http.ListenAndServe("", nil)
}
