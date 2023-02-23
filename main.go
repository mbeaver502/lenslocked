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

	r.Get("/", createHandler("home.gohtml"))
	r.Get("/contact", createHandler("contact.gohtml"))
	r.Get("/faq", createHandler("faq.gohtml"))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", r)
}

func createHandler(filename string) http.HandlerFunc {
	tpl := views.Must(views.Parse(filepath.Join("templates", filename)))
	return controllers.StaticHandler(tpl)
}
