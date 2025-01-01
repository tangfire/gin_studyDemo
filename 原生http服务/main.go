package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.String())
	if r.Method != "GET" {
		byteData, _ := io.ReadAll(r.Body)
		fmt.Println(string(byteData))
	}

	fmt.Println(r.Header)

	byteData, _ := json.Marshal(Response{Code: 0, Message: "success", Data: map[string]any{}})

	//w.Write([]byte("Hello world"))

	w.Write(byteData)
}

func main() {

	http.HandleFunc("/index", Index)

	fmt.Println("http server running 127.0.0.1:8080")

	//http.ListenAndServe(":8080", nil)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
