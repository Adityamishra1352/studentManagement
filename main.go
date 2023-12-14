package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func operations(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "<h1>Hello I am Student Management API</h1>")
	case "/delete":
		fmt.Fprint(w, "<h1>Delete Student</h1>")
		deleteStudent(w, r)
	case "/add":
		fmt.Fprint(w, "<h1>Add Student</h1>")
		addStudent(w, r)
	case "/update":
		fmt.Fprint(w, "<h1>Update Student Details</h1>")
		updateStudent(w, r)
	case "/view":
		fmt.Fprint(w, "<h1>View Students</h1>")
		viewStudents(w, r)
	default:
		fmt.Fprint(w, "Error")
	}
}
func updateStudent(w http.ResponseWriter, r *http.Request) {
	var err error
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("Invalid Id Type")
	}
	name := r.FormValue("name")
	ageString := r.FormValue("age")
	age, _ := strconv.Atoi(ageString)
	enrollment := r.FormValue("enrollment")
	_, err = db.Exec("UPDATE students SET name=?,age=?, enrollment=? where id=?", name, age, enrollment, id)
	if err != nil {
		fmt.Println("Error updating the student record")
	}
}
func viewStudents(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.Query("SELECT * FROM students")
	defer rows.Close()
	var students []string
	for rows.Next() {
		var id int
		var name string
		var age int
		var enrollment string

		err := rows.Scan(&id, &name, &age, &enrollment)
		if err != nil {
			fmt.Println("Error scanning student:", err)
			continue
		}

		studentInfo := fmt.Sprintf("ID: %d, Name: %s, Age: %d, Enrollment: %s", id, name, age, enrollment)
		students = append(students, studentInfo)
	}
	for _, student := range students {
		fmt.Fprintln(w, student)
	}
}
func addStudent(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	ageString := r.FormValue("age")
	age, _ := strconv.Atoi(ageString)
	enrollment := r.FormValue("enrollment")
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

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("Invalid ID type")
	}
	_, err = db.Exec("DELETE FROM students WHERE id=?", id)
	if err != nil {
		fmt.Println("Error seleteing student details")
	}
	if err == nil {
		fmt.Println("Student details deleted successfully")
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
