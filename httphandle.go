package main

import (
	"fmt"
	"net/http"
)

func main() {

	Handler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}
	http.HandleFunc("/", Handler)

	fmt.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
