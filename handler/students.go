package handler

import (
	"html/template"
	"log"
	"net/http"
)

type FormError struct {
	NameError string
}
type Student struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Roll        int    `json:"roll"`
	English     int    `json:"english"`
	Bangla      int    `json:"bangla"`
	Mathematics int    `json:"mathematics"`
	Grade       string `json:"grade"`
	GPA float64 `json:"gpa"`
	FormError
}

type StudentList struct {
	Students []Student `json:"students"`
}

func StudentsList (w http.ResponseWriter, r *http.Request) {
	sl, err := getStudentsList()
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles("templates/students-list.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, sl)
}

func CreateStudent (w http.ResponseWriter, r *http.Request) {
	pharseCreateStudent(w, nil)
}
