package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Username string
	Email    string
	Password string
}

var newUser User

func main() {
	http.HandleFunc("/signup", showSignupForm)

	http.HandleFunc("/signup/process", processSignupForm)

	http.ListenAndServe(":8080", nil)
}

func showSignupForm(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("signup.html"))

	tpl.Execute(w, nil)
}

func processSignupForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

}
