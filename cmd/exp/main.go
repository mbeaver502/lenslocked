package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string // template.HTML will NOT be encoded automatically
}

func main() {
	// file path is relative to the binary
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "John Doe",
		// html/template will encode this -- UNLESS the type is template.HTML
		// text/template will NOT encode this
		Bio: `<script>alert("Haha, you have been h4x0r3d!");</script>`,
	}

	// We can give any type for `data`
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
