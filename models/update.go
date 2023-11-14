package models

import (
	"context"
	"crud/views"
	"fmt"
	"os"
)

func UpdateStudent(update *views.Students) error {
	updateStudent := `UPDATE students SET first_name=$1, last_name=$2  WHERE roll_number=$3;`

	_, err := con.Exec(context.Background(), updateStudent, update.FirstName, update.LastName, update.RollNumber)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error updating student: %v\n", err)
		return err
	}

	fmt.Println("Student updated successfully")
	return nil
}

func UpdatParents(update *views.Parents) error {
	updateParents := `UPDATE parents SET parent_first_name=$1, parent_last_name=$2 WHERE student_id =$3;`

	_, err := con.Exec(context.Background(), updateParents, update.ParentFirstName, update.ParentLastName, update.ParentID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error updating parents: %v\n", err)
		return err
	}

	fmt.Println("Parents updated successfully")
	return nil
}
