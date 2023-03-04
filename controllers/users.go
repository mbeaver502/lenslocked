package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mbeaver502/lenslocked/models"
)

type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}
	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}

	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}

	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(w, data)
}

func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}

	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")

	user, err := u.UserService.Authenticate(data.Email, data.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     "email",
		Value:    user.Email,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	log.Println("Authenticated user:", user)
	fmt.Fprintf(w, "Authenticated user: %+v", user)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	// We can use PostFormValue to automatically parse the form
	// before retrieving a value (if it exists), but we ignore
	// any parse errors that occur.
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	user, err := u.UserService.Create(email, password)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	log.Println("User created:", email)
	fmt.Fprintf(w, "User created: %+v", user)
}

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	email, err := r.Cookie("email")
	if err != nil {
		log.Println(err)
		http.Error(w, "The email cookie could not be read.", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Email cookie: %s\n", email.Value)
}
