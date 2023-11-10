package models

import (
	"context"
	"crud/views"
)

func ReadAllStudents() ([]views.Students, error) {
	rows, err := con.Query(context.Background(), "SELECT first_name,last_name,roll_number FROM students")

	if err != nil {

		return nil, err
	}

	all := []views.Students{}
	for rows.Next() {
		data := views.Students{}
		rows.Scan(&data.FirstName, &data.LastName, &data.RollNumber)
		all = append(all, data)
	}
	return all, nil

}

func ReadAllParents() ([]views.Parents, error) {
	rows, err := con.Query(context.Background(), "SELECT parent_first_name,parent_last_name from parents")

	if err != nil {
		return nil, err
	}

	parents := []views.Parents{}
	for rows.Next() {
		data := views.Parents{}
		rows.Scan(&data.ParentFirstName, &data.ParentLastName)
		parents = append(parents, data)
	}
	return parents, nil
}

func GetUser(username string) (*views.User, error) {
	var user views.User

	err := con.QueryRow(context.Background(), "SELECT username, password FROM users WHERE username = $1", username).
		Scan(&user.UserName, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
