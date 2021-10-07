package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mohamed-samir907/todo-go/views"
)

func ping() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			response := views.Response{
				Code: 200,
				Body: "Pong",
			}

			json.NewEncoder(res).Encode(response)
		}
	}
}
