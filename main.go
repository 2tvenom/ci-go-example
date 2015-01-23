package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
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
	request = &Request{
		Code: 100,
		Data: []string{"Hello", "World"},
		Description: "No description",
	}
	return
	resp, err := http.Get("http://127.0.0.1:8080/api/list")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		println(err.Error())
		return
	}

	buff := bytes.NewBuffer(body)

	decoder := json.NewDecoder(buff)

	request = &Request{}

	err = decoder.Decode(request)

	if err != nil {
		println(err.Error())
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
