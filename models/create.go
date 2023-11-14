package models

import (
	"context"
	"crud/views"
	"fmt"
	"os"

	"github.com/gofrs/uuid"
)

func InsertStudent(stud *views.Students) error {
	fmt.Println("created entry :", stud.FirstName, stud.LastName, stud.RollNumber)
	insertStudent := `INSERT INTO students (first_name, last_name, roll_Number, parent_id)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (roll_number) DO NOTHING
        RETURNING student_id`

	var entryID uuid.UUID

	err := con.QueryRow(context.Background(), insertStudent, stud.FirstName, stud.LastName, stud.RollNumber, stud.ParentID).Scan(&entryID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error inserting the entry: %v\n", err)
		return err
	}

	fmt.Printf("Entry inserted with id %d\n", entryID)
	return nil
}

func InsertParent(parent *views.Parents) error {
	insert := `	with insert AS(
		INSERT INTO parents (parent_first_name, parent_last_name)
	VALUES ($1, $2 )
	ON CONFLICT  DO NOTHING RETURNING parent_id
	) 
	select parent_id from insert 
	UNION 
	select parent_id from parents where parent_first_name=$1 AND parent_last_name=$2`

	err := con.QueryRow(context.Background(), insert, parent.ParentFirstName, parent.ParentLastName).Scan(&parent.ParentID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error inserting the entry: %v\n", err)
		return err
	}

	fmt.Println("Entry inserted successfully")
	return nil
}

func InsertUser(user *views.User) error {
	insert := `INSERT INTO USERS (username, password)
	VALUES($1,$2);`

	_, err := con.Query(context.Background(), insert, user.UserName, user.Password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error inserting the user: %v\n", err)

	}

	fmt.Println("user inserted successfully")
	return nil
}
