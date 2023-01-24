package handler

import (
	"html/template"
	"log"
	"net/http"
)

func pharseCreateStudent(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("templates/create-students.html")
	if err != nil {
		log.Fatalf("%v", err)
	}
	t.Execute(w, data)
}

func pharseEditStudent(w http.ResponseWriter, data any) {
	var err error
	t := template.Must(template.ParseFiles("templates/header.html", "templates/footer.html", "templates/edit-student.html"))
	if err != nil {
		log.Fatalf("%v", err)
	}
	t.ExecuteTemplate(w, "edit-student.html", data)
}
