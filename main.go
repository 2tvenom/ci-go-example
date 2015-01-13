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


	resp, err := http.Get("http://localhost:8000/api/list")
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
	http.ListenAndServe(":8080", nil)
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
