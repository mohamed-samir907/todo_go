package controllers

import (
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", ping())
	mux.HandleFunc("/", CreateTodo())

	return mux
}
