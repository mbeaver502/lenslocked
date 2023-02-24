package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}

	// Get the value from the `email` URL query param
	data.Email = r.FormValue("email")

	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	// err := r.ParseForm()
	// if err != nil {
	// 	http.Error(w, "Failed to parse form", http.StatusBadRequest)
	// 	return
	// }

	// email := r.PostForm.Get("email")
	// password := r.PostForm.Get("password")

	// We can use PostFormValue to automatically parse the form
	// before retrieving a value (if it exists), but we ignore
	// any parse errors that occur.
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	fmt.Fprint(w, email, password)
}
