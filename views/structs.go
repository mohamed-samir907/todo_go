package views

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

type Todo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r Response) SendJson(res http.ResponseWriter) {
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(r.Code)
	json.NewEncoder(res).Encode(r.Body)
}
