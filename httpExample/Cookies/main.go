package main

import (
	"net/http"
	"strconv"
)

func main(){
	// Set up the route and handler
	http.HandleFunc("/", cookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}

func cookie(w http.ResponseWriter, r *http.Request){
	// Initialize counter
	counter := 0
	// Try to get the cookie from the request
	cookie, err := r.Cookie("counter-cookie")
	if err == http.ErrNoCookie {
		// If no cookie is found, set a new cookie with value 0
		http.SetCookie(w, &http.Cookie{
			Name: "counter-cookie",
			Value: "0",
		})
	}

	// Convert cookie value to integer
	counter, _ = strconv.Atoi(cookie.Value)
	// Increment the counter
	counter++
	// Update the cookie value
	cookie.Value = strconv.Itoa(counter)

	// Set the updated cookie in the response
	http.SetCookie(w, cookie)
	// Write the response to the client
	w.Write([]byte("You have visited this page " + strconv.Itoa(counter) + " times"))
}
