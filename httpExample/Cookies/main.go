package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main(){
	// Set up the route and handler
	http.HandleFunc("/", cookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	// Start the server on port 8080
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func cookie(w http.ResponseWriter, r *http.Request){
	// Initialize counter
	var counter int
	// Try to get the cookie from the request
	cookie, err := r.Cookie("counter-cookie")
	if err == http.ErrNoCookie {
		// If no cookie is found, set a new cookie with value 0
		http.SetCookie(w, &http.Cookie{
			Name: "counter-cookie",
			Value: "1",
		})
		counter = 1
	} else if err != nil {
		// If there is an error, write the error to the response
		w.Write([]byte(err.Error()))
		return
	} else {
		counter, _ = strconv.Atoi(cookie.Value)
		counter++
		fmt.Println("Incremented counter: ", counter)
		cookie.Value = strconv.Itoa(counter)
	
		// Set the updated cookie in the response
		http.SetCookie(w, cookie)
	}
	// Write the response to the client
	w.Write([]byte("You have visited this page " + strconv.Itoa(counter) + " times"))
}
