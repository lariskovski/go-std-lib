package main

import (
	"os"
	"text/template"
)

type User struct {
	Username string
	Name     string
	IsAdmin  bool
}

func main(){

	user := User{
		"lariskovski",
		"larissa",
		true,
	}

	template := template.Must(template.ParseGlob("templates/*"))
	// template := template.Must(template.ParseFiles("templates/index.gohtml"))
	err := template.ExecuteTemplate(os.Stdout, "index.gohtml", user)
	if err != nil {
		panic(err)
	}
}