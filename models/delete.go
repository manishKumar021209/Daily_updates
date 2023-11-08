package models

import (
	"context"
	"fmt"
	"os"

	"github.com/gofrs/uuid"
)

func DeleteStudent(studentID uuid.UUID) error {
	delete := `DELETE FROM students where student_id=$1;`

	_, err := con.Exec(context.Background(), delete, studentID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting student : %v\n", err)
		return err
	}

	fmt.Println("Entry deleted successfully")
	return nil
}
