package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type RequestData struct {
	Method string
	URL    string
	Proto  string
	Header http.Header
	Form   map[string][]string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func index(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func request(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// Handle error
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "text/html")

	data := RequestData{
		Method: r.Method,
		URL:    r.URL.String(),
		Proto:  r.Proto,
		Header: r.Header,
		Form:   r.Form,
	}

	err = tpl.ExecuteTemplate(w,"request.gohtml" ,data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/request", request)
	http.ListenAndServe(":8080", nil)
}