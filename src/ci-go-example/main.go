package main

import (
	"net/http"
	"encoding/json"
)

type (
	Request struct {
		Code int `json:"code"`
		Data []string `json:"data"`
		Description string `json:"description"`
	}
)

var (
	request *Request
)

func init() {
	//test
	request = &Request{
		Code: 100,
		Data: []string{"Hello", "World"},
		Description: "No description",
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8082", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	encoder := json.NewEncoder(w)
	encoder.Encode(request)
}

func sum(d1, d2 int) int {
	return d1 + d2
}
