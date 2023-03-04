package controllers

import (
	"html/template"
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML // Use the text verbatim as HTML -- Insecure for untrusted HTML text!
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes, but only for students.",
		},
		{
			Question: "What are your support hours?",
			Answer:   "Our support staff are available 9-5, M-F.",
		},
		{
			Question: "How do I contact support?",
			Answer:   `Email us - <a href="mailto:somebody@example.com">somebody@example.com</a>`,
		},
		{
			Question: "Where is your office located?",
			Answer:   "On the surface of the sun.",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, questions)
	}
}
