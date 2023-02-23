package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/mbeaver502/lenslocked/views"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")

	t, err := views.Parse(tplPath)
	if err != nil {
		log.Printf("%v", err)
		http.Error(w, "Error while parsing template.", http.StatusInternalServerError)
		return
	}

	t.Execute(w, r)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")

	t, err := views.Parse(tplPath)
	if err != nil {
		log.Printf("%v", err)
		http.Error(w, "Error while parsing template.", http.StatusInternalServerError)
		return
	}

	t.Execute(w, r)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "faq.gohtml")

	t, err := views.Parse(tplPath)
	if err != nil {
		log.Printf("%v", err)
		http.Error(w, "Error while parsing template.", http.StatusInternalServerError)
		return
	}

	t.Execute(w, r)
}
