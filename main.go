package main

import (
	"net/http"

	"golang_example/pkg/db"
	"golang_example/src/authentication"
	"golang_example/src/home"
)

func main() {
	db.Open()

	http.HandleFunc("/authorization", authentication.Handle)
	http.HandleFunc("/home", home.Handle)

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}
