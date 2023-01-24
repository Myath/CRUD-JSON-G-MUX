package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteStudent (w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	uID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	sl, err := getStudentsList()
	if err != nil {
		log.Fatal(err)
	}

	var newStudentList []Student
	for _, student := range sl.Students {
		if student.ID == uID {
			continue
		}
		newStudentList = append(newStudentList, student)
	}
	sl.Students = newStudentList

	if err := writeStudentToFile(sl); err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/students", http.StatusSeeOther)

}
