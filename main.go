package main

import (
	"fmt"
	"net/http"

	"github.com/mohamed-samir907/todo-go/controllers"
	"github.com/mohamed-samir907/todo-go/models"
)

func main() {
	mux := controllers.RegisterRoutes()
	db := models.Connect()
	defer db.Close()
	fmt.Println("Starting server on port 8000")
	http.ListenAndServe(":8000", mux)
}
