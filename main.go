package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"github.com/mbeaver502/lenslocked/controllers"
	"github.com/mbeaver502/lenslocked/models"
	"github.com/mbeaver502/lenslocked/templates"
	"github.com/mbeaver502/lenslocked/views"
)

func main() {
	r := chi.NewRouter()

	r.Get("/",
		controllers.StaticHandler(
			views.Must(
				views.ParseFS(templates.FS, "tailwind.gohtml", "home.gohtml"))))

	r.Get("/contact",
		controllers.StaticHandler(
			views.Must(
				views.ParseFS(templates.FS, "tailwind.gohtml", "contact.gohtml"))))

	r.Get("/faq",
		controllers.FAQ(
			views.Must(
				views.ParseFS(templates.FS, "tailwind.gohtml", "faq.gohtml"))))

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = models.Migrate(db, "migrations")
	if err != nil {
		panic(err)
	}

	userService := models.UserService{
		DB: db,
	}
	sessionService := models.SessionService{
		DB: db,
	}

	usersC := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}

	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "signup.gohtml"))
	usersC.Templates.SignIn = views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "signin.gohtml"))

	r.Get("/signup", usersC.New)
	r.Get("/signin", usersC.SignIn)
	r.Post("/users", usersC.Create)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Post("/signout", usersC.ProcessSignOut)
	r.Get("/users/me", usersC.CurrentUser)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMW := csrf.Protect([]byte(csrfKey), csrf.Secure(false)) // TODO: remove csrf.Secure(false)

	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", csrfMW(r))
}
