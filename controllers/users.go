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
	UserService    *models.UserService
	SessionService *models.SessionService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}

	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, r, data)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}

	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(w, r, data)
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

	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	setCookie(w, CookieSession, session.Token)

	http.Redirect(w, r, "/users/me", http.StatusFound)
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

	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	setCookie(w, CookieSession, session.Token)

	http.Redirect(w, r, "/users/me", http.StatusFound)
}

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	token, err := readCookie(r, CookieSession)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	user, err := u.SessionService.User(token)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	fmt.Fprintf(w, "Current User: %+v\n", user)
}

func (u Users) ProcessSignOut(w http.ResponseWriter, r *http.Request) {
	token, err := readCookie(r, CookieSession)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	err = u.SessionService.Delete(token)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	deleteCookie(w, CookieSession)

	http.Redirect(w, r, "/signin", http.StatusFound)
}
