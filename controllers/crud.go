package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mohamed-samir907/todo-go/models"
	"github.com/mohamed-samir907/todo-go/views"
)

func CreateTodo() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		switch req.Method {
		case http.MethodPost:
			create(res, req)
		case http.MethodDelete:
			deleteByName(res, req)
		case http.MethodGet:
			if req.URL.Query().Has("name") {
				getByName(res, req)
			} else {
				getAll(res, req)
			}
		default:
			res.WriteHeader(http.StatusMethodNotAllowed)
			res.Write([]byte("Method not allowed"))
		}
	}
}

func create(res http.ResponseWriter, req *http.Request) {
	todo := views.Todo{}

	json.NewDecoder(req.Body).Decode(&todo)

	err := models.Create(todo.Name, todo.Description)

	if err != nil {
		res.Write([]byte("An error occured"))
		log.Fatal(err.Error())
	}

	res.Write([]byte("Todo created\n"))
}

func getAll(res http.ResponseWriter, req *http.Request) {
	todos, err := models.ReadAll()

	if err != nil {
		res.Write([]byte("An error occured"))
		log.Fatal(err.Error())
	}

	response := views.Response{Code: http.StatusOK, Body: todos}
	response.SendJson(res)
}

func getByName(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	todos, err := models.ReadByName(name)

	if err != nil {
		res.Write([]byte("An error occured"))
		log.Fatal(err.Error())
	}

	response := views.Response{Code: http.StatusOK, Body: todos}
	response.SendJson(res)
}

func deleteByName(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	if name == "" {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("Name paramter not found"))
		return
	}

	err := models.Delete(name)

	if err != nil {
		res.Write([]byte("An error occured"))
		log.Fatal(err.Error())
	}

	response := views.Response{Code: http.StatusOK, Body: "Deleted successfully"}
	response.SendJson(res)
}
