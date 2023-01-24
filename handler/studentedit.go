package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func StudentEdit (w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	uID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	sl, err := getStudentsList()
	if err != nil {
		log.Fatal(err)
	}

	var editStudent Student

	for _, student := range sl.Students {
		if student.ID == uID {
			editStudent = student
			break
		}
	}

	pharseEditStudent(w, editStudent)
}


func StudentUpdate (w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	uID, err := strconv.Atoi(id)
	if err!= nil {
		log.Fatal(err)
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	roll, _ := strconv.Atoi(r.FormValue("roll"))
	eng, _ := strconv.Atoi(r.FormValue("eng"))
	ban, _ := strconv.Atoi(r.FormValue("ban"))
	math, _ := strconv.Atoi(r.FormValue("math"))

	if name == "" {
		pharseEditStudent(w, Student{ID: uID, Name: name, Email: email, Roll: roll, English: eng, Bangla: ban, Mathematics: math, FormError: FormError{NameError: "All field is required."}})
		return
	} else if email == "" {
		pharseEditStudent(w, Student{ID: uID, Name: name, Email: email, Roll: roll, English: eng, Bangla: ban, Mathematics: math, FormError: FormError{NameError: "All field is required."}})
		return
	} else if roll == 0 {
		pharseEditStudent(w, Student{ID: uID, Name: name, Email: email, Roll: roll, English: eng, Bangla: ban, Mathematics: math, FormError: FormError{NameError: "All field is required."}})
		return
	} else if r.FormValue("eng") == "" {
		pharseEditStudent(w, Student{ID: uID, Name: name, Email: email, Roll: roll, English: eng, Bangla: ban, Mathematics: math, FormError: FormError{NameError: "All field is required."}})
		return
	} else if r.FormValue("ban") == "" {
		pharseEditStudent(w, Student{ID: uID, Name: name, Email: email, Roll: roll, English: eng, Bangla: ban, Mathematics: math, FormError: FormError{NameError: "All field is required."}})
		return
	} else if r.FormValue("math") == ""{
		pharseEditStudent(w, Student{ID: uID, Name: name, Email: email, Roll: roll, English: eng, Bangla: ban, Mathematics: math, FormError: FormError{NameError: "All field is required."}})
		return
	}

	sl, err := getStudentsList()
	if err!= nil {
		log.Fatal(err)
	}

	grade, gpa := Grade(eng,ban,math)

	for i, student := range sl.Students{
		if student.ID == uID {
			sl.Students[i].Name = name
			sl.Students[i].Email = email
			sl.Students[i].Roll = roll
			sl.Students[i].English = eng
			sl.Students[i].Bangla = ban
			sl.Students[i].Mathematics = math
			sl.Students[i].Grade = grade
			sl.Students[i].GPA = gpa
            break
		}
		
	}

	if err := writeStudentToFile(sl); err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/students", http.StatusSeeOther)
}
