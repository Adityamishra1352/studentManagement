package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func operations(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "<h1>Hello I am Student Management API</h1>")
	case "/delete":
		fmt.Fprint(w, "<h1>Delete Student</h1>")
	case "/add":
		fmt.Fprint(w, "<h1>Add Student</h1>")
		addStudent(w, r)
	case "/update":
		fmt.Fprint(w, "<h1>Update</h1>")
	default:
		fmt.Fprint(w, "Error")
	}
}
func addStudent(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	age := r.FormValue("age")
	enrollment := r.FormValue("enrolllment")
	var err error
	_, err = db.Exec("INSERT INTO students (name, age, enrollment) VALUES (?, ?, ?)", name, age, enrollment)
	if err != nil {
		fmt.Println("Error adding student to the database")
		return
	}
	if err == nil {
		fmt.Println("Student added successfully")
	}
}

func createTable() {
	newTable := `CREATE TABLE students(id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(255), age INTEGER, enrollment VARCHAR(255));`
	_, err := db.Exec(newTable)
	if err != nil {
		fmt.Println("error creating the database")
		return
	}
}

func main() {
	var err error
	db, err = sql.Open("sqlite3", "students.db")
	if err != nil {
		fmt.Println("error opening database")
		return
	}
	createTable()
	http.HandleFunc("/", operations)
	http.ListenAndServe(":4000", nil)
}
