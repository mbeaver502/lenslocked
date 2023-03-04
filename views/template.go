package views

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
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

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	tpl := template.New(patterns[0])

	// Add placeholder func to allow templates to parse successfully
	tpl = tpl.Funcs(template.FuncMap{
		"csrfField": func() (template.HTML, error) {
			return "", fmt.Errorf("csrfField is not implemented")
		},
	})

	tpl, err := tpl.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	return Template{
		htmlTmpl: tpl,
	}, nil
}

// Parse the given `filepath` into an HTML template.
// func Parse(filepath string) (Template, error) {
// 	tpl, err := template.ParseFiles(filepath)
// 	if err != nil {
// 		return Template{}, fmt.Errorf("parsing template: %w", err)
// 	}

// 	return Template{
// 		htmlTmpl: tpl,
// 	}, nil
// }

// Execute a template by writing any `data` into `w`.
func (t Template) Execute(w http.ResponseWriter, r *http.Request, data any) {
	// Clone allows multiple users to access the same template without overwriting others' data
	tpl, err := t.htmlTmpl.Clone()
	if err != nil {
		log.Printf("cloning template: %v", err)
		http.Error(w, "There was an error cloning the template.", http.StatusInternalServerError)
		return
	}

	// Provide the actual implementation for the csrfField template func
	tpl = tpl.Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return csrf.TemplateField(r)
		},
	})

	w.Header().Set("Content-Type", "text/html")

	var buf bytes.Buffer

	err = tpl.Execute(&buf, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}

	io.Copy(w, &buf)
}
