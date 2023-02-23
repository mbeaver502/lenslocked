package views

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"text/template"
)

type Template struct {
	htmlTmpl *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}

	return t
}

func ParseFS(fs fs.FS, pattern string) (Template, error) {
	tpl, err := template.ParseFS(fs, pattern)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	return Template{
		htmlTmpl: tpl,
	}, nil
}

// Parse the given `filepath` into an HTML template.
func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	return Template{
		htmlTmpl: tpl,
	}, nil
}

// Execute a template by writing any `data` into `w`.
func (t Template) Execute(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "text/html")

	err := t.htmlTmpl.Execute(w, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}
