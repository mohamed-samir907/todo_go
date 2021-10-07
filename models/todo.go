package models

import "github.com/mohamed-samir907/todo-go/views"

func Create(name, description string) error {
	query, err := conn.Query(
		"INSERT INTO todos (name, description) VALUES(?, ?)",
		name,
		description,
	)

	if err != nil {
		return err
	}

	defer query.Close()

	return nil
}

func ReadAll() ([]views.Todo, error) {
	query, err := conn.Query("SELECT * FROM todos")

	if err != nil {
		return nil, err
	}

	todos := []views.Todo{}

	for query.Next() {
		todo := views.Todo{}
		query.Scan(&todo.Name, &todo.Description)

		todos = append(todos, todo)
	}

	defer query.Close()

	return todos, nil
}

func ReadByName(name string) ([]views.Todo, error) {
	query, err := conn.Query("SELECT * FROM todos WHERE name = ?", name)

	if err != nil {
		return nil, err
	}

	todos := []views.Todo{}

	for query.Next() {
		todo := views.Todo{}
		query.Scan(&todo.Name, &todo.Description)

		todos = append(todos, todo)
	}

	defer query.Close()

	return todos, nil
}

func Delete(name string) error {
	query, err := conn.Query("DELETE FROM todos WHERE name = ?", name)

	if err != nil {
		return err
	}

	defer query.Close()

	return nil
}
