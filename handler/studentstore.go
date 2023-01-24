package handler

import (
	"log"
	"net/http"
	"strconv"
)

func StudentStore (w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	name := r.FormValue("name")
	email := r.FormValue("email")
	roll, _ := strconv.Atoi(r.FormValue("roll"))
	eng, _ := strconv.Atoi(r.FormValue("eng"))
	ban, _ := strconv.Atoi(r.FormValue("ban"))
	math, _ := strconv.Atoi(r.FormValue("math"))

	if name == "" {
		pharseCreateStudent(w, FormError{NameError: "All field is required."})
		return
	} else if email == "" {
		pharseCreateStudent(w, FormError{NameError: "All field is required."})
		return
	} else if r.FormValue("roll") == "" {
		pharseCreateStudent(w, FormError{NameError: "All field is required."})
		return
	} else if r.FormValue("eng") == "" {
		pharseCreateStudent(w, FormError{NameError: "All field is required."})
		return
	} else if r.FormValue("ban") == "" {
		pharseCreateStudent(w, FormError{NameError: "All field is required."})
		return
	} else if r.FormValue("math") == "" {
		pharseCreateStudent(w, FormError{NameError: "All field is required."})
		return
	}

	sl, err := getStudentsList()
	if err != nil {
		log.Fatal(err)
	}

	for _, student := range sl.Students {
		if student.Name == name {
			pharseCreateStudent(w, FormError{NameError: "Student already exists."})
			return
		} else if student.Email == email {
			pharseCreateStudent(w, FormError{NameError: "Student already exists."})
			return
		} else if student.Roll == roll {
			pharseCreateStudent(w, FormError{NameError: "Student already exists."})
			return
		}
	}

	lastStudent := Student{}
	if len(sl.Students) >= 1 {
		lastStudent = sl.Students[len(sl.Students)-1]
	}

	grade, gpa := Grade(eng,ban,math)

	newStudent := Student{
		ID:          lastStudent.ID + 1,
		Name:        name,
		Email:       email,
		Roll:        roll,
		English:     eng,
		Bangla:      ban,
		Mathematics: math,
		Grade:       grade,
		GPA:         gpa,
	}

	sl.Students = append(sl.Students, newStudent)

	if err := writeStudentToFile(sl); err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/students", http.StatusSeeOther)
}
