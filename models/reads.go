package models

import "api/views"

func ReadAll() ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM TODO")
	if err != nil {
		return nil, err
	}

	todo := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todo = append(todo, data)
	}
	return todo, nil
}

func ReadByName(name string) ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM TODO WHERE name = $1", name)
	if err != nil {
		return nil, err
	}

	todo := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todo = append(todo, data)
	}
	return todo, nil
}
