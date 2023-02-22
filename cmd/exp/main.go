package main

import (
	"html/template"
	"os"
)

// type User struct {
// 	Name string
// }

func main() {
	// file path is relative to the binary
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	// user := User{
	// 	Name: "John Doe",
	// }

	user := struct {
		Name string // hello.gohtml is looking for a "Name" field
	}{
		Name: "John Doe",
	}

	// We can give any type for `data`
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
