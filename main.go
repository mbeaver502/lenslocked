package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
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

	usersC := controllers.Users{
		UserService: &models.UserService{
			DB: db,
		},
	}

	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "signup.gohtml"))
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", r)
}
