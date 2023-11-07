package controllers

import (
	"crud/models"
	"crud/views"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var insert struct {
			FirstName  string `json:"firstname"`
			LastName   string `json:"lastname"`
			RollNumber int    `json:"rollnumber"`
			PFirstName string `json:"fathername"`
			PLastName  string `json:"Mothername"`
		}
		var err error
		if err = json.NewDecoder(r.Body).Decode(&insert); err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}

		stud := views.Students{
			FirstName:  insert.FirstName,
			LastName:   insert.LastName,
			RollNumber: insert.RollNumber,
		}

		// Insert the student and get the student ID
		studID, err := models.InsertStudent(&stud)
		if err != nil {
			http.Error(w, "Failed to create student", http.StatusBadRequest)
			return
		}

		fmt.Println(studID)
		parent := views.Parents{
			ParentFirstName: insert.PFirstName,
			ParentLastName:  insert.PLastName,
		}

		// Insert the parent, associating it with the student using studID
		err = models.InsertParent(studID, &parent)
		if err != nil {
			http.Error(w, "Failed to create parent", http.StatusBadRequest)
			return
		}

	}
}

func ReadAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		students, err := models.ReadAllStudents()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		parents, err := models.ReadAllParents()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			Students []views.Students
			Parents  []views.Parents
		}{
			Students: students,
			Parents:  parents,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
