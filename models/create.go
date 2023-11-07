package models

import (
	"context"
	"crud/views"
	"fmt"
	"os"

	"github.com/gofrs/uuid"
)

func InsertStudent(stud *views.Students) (uuid.UUID, error) {
	fmt.Println("created entry :", stud.FirstName, stud.LastName, stud.RollNumber)
	insertStudent := `INSERT INTO students (first_name, last_name, rollNumber)
        VALUES ($1, $2, $3)
        ON CONFLICT (first_name, last_name) DO NOTHING
        RETURNING id;`

	var entryID uuid.UUID

	err := con.QueryRow(context.Background(), insertStudent, stud.FirstName, stud.LastName, stud.RollNumber).Scan(&entryID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error inserting the entry: %v\n", err)
		return uuid.Nil, err
	}

	fmt.Printf("Entry inserted with id %d\n", entryID)
	return entryID, nil
}

func InsertParent(studentID uuid.UUID, parent *views.Parents) error {
	insert := `	INSERT INTO parents (student_id, parent_first_name, parent_last_name)
	VALUES ($1, $2, $3)
	ON CONFLICT (student_id, parent_first_name, parent_last_name) DO NOTHING;`

	_, err := con.Exec(context.Background(), insert, studentID, parent.ParentFirstName, parent.ParentLastName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error inserting the entry: %v\n", err)
		return err
	}

	fmt.Println("Entry inserted successfully")
	return nil
}
