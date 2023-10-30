package models

import "fmt"

func CreateTodo(name, todo string) error {
	fmt.Println("created entry:", name, todo)
	insertQ, err := con.Query("INSERT INTO TODO VALUES($1, $2)", name, todo)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer func() {
		if insertQ != nil {
			insertQ.Close()
		}
	}()

	return nil
}

func DeleteTodo(name string) error {
	fmt.Println("deleted :", name)
	insertQ, err := con.Query("DELETE FROM TODO WHERE name=$1", name)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// func DeleteTodo(name string) error {
// 	_, err := con.Exec("DELETE FROM TODO WHERE name=$1", name)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	return nil
// }

// func DeleteTodo(name string) error {
// 	tx, err := con.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	defer tx.Rollback()

// 	_, err = tx.Exec("DELETE FROM TODO WHERE name = $1", name)
// 	if err != nil {
// 		return err
// 	}

// 	// Commit the transaction to save the changes
// 	if err := tx.Commit(); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func DeleteTodo(name string) error {
// 	deleteQ, err := con.Prepare("DELETE FROM TODO WHERE name = $1")
// 	if err != nil {
// 		return err
// 	}
// 	defer deleteQ.Close()

// 	_, err = deleteQ.Exec(name)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
