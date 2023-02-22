package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", pathHandler)
	fmt.Println("starting server on :3000...")
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello, World!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>No contact info available at this time...</h1>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
		//http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}
