package controllers

import (
	"crud/models"
	"crud/views"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
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

func updateS(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {

		var insert struct {
			FirstName  string `json:"firstname"`
			LastName   string `json:"lastname"`
			RollNumber int    `json:"rollnumber"`
			// PFirstName string `json:"fathername"`
			// PLastName  string `json:"Mothername"`
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

		err = models.UpdateStudent(&stud)
		if err != nil {
			http.Error(w, "Failed to update student", http.StatusBadRequest)
			return
		}

	}
}

func updateP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {

		var insert struct {
			PID        string `json:"pid"`
			PFirstName string `json:"parentfirstname"`
			PLastName  string `json:"parentlastname"`
		}

		var err error
		if err = json.NewDecoder(r.Body).Decode(&insert); err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}
		fmt.Println(insert.PID)
		fmt.Println(insert.PFirstName)
		fmt.Println(insert.PLastName)

		parent := views.Parents{
			StudentID:       insert.PID,
			ParentFirstName: insert.PFirstName,
			ParentLastName:  insert.PLastName,
		}

		err = models.UpdatParents(&parent)
		if err != nil {
			http.Error(w, "Failed to update parents", http.StatusBadRequest)
			return
		}

	}
}

func DeleteByID(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodDelete {
		params := mux.Vars(r)
		studentid := params["studentid"]
		studentid = r.URL.Query().Get("id")

		studentUID, err := uuid.FromString(studentid)

		if err != nil {
			http.Error(w, "Invalid student ID", http.StatusBadRequest)
			return
		}

		if err := models.DeleteStudent(studentUID); err != nil {
			http.Error(w, "Failed to delete student", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Student deleted successfully")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
