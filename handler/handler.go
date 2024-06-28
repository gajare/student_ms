package handler

import (
	"encoding/json"
	"net/http"
	"student_ms/models"

	"github.com/gorilla/mux"
)

var students = []models.Student{
	{ID: "1", Name: "Amol Gahare", Age: 20, Grade: "A"},
	{ID: "2", Name: "Naina gajare", Age: 22, Grade: "B"},
	{ID: "3", Name: "Sandip Gajare", Age: 25, Grade: "C"},
	{ID: "4", Name: "Rohit bhoi", Age: 27, Grade: "D"},
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range students {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Student not found", http.StatusNotFound)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	students = append(students, student)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range students {
		if item.ID == params["id"] {
			students = append(students[:index], students[index+1:]...)
			var student models.Student
			_ = json.NewDecoder(r.Body).Decode(&student)
			student.ID = params["id"]
			students = append(students, student)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(student)
			return
		}
	}
	http.Error(w, "Student not found", http.StatusNotFound)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range students {
		if item.ID == params["id"] {
			students = append(students[:index], students[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(students)
			return
		}
	}
	http.Error(w, "Student not found", http.StatusNotFound)
}
