package main

import (
	"log"
	"net/http"
	"text/template"
)


type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// Handle error
		log.Fatal(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}
var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}