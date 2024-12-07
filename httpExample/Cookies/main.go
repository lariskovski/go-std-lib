package main

import "net/http"

func main(){
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.ListenAndServe(":8080", nil)
}


func set(w http.ResponseWriter, r *http.Request){
	http.SetCookie(w, &http.Cookie{
		Name: "my-cookie",
		Value: "some value",
	})
	w.Write([]byte("Cookie set"))
}

func read(w http.ResponseWriter, r *http.Request){
	c, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Write([]byte("Your cookie: " + c.Value))
}