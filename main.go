package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/mbeaver502/lenslocked/controllers"
	"github.com/mbeaver502/lenslocked/views"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(parseTemplate("home.gohtml")))
	r.Get("/contact", controllers.StaticHandler(parseTemplate("contact.gohtml")))
	r.Get("/faq", controllers.StaticHandler(parseTemplate("faq.gohtml")))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", r)
}

func parseTemplate(filename string) views.Template {
	return views.Must(views.Parse(filepath.Join("templates", filename)))
}
