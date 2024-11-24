package main

import "net/http"


type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}


func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}