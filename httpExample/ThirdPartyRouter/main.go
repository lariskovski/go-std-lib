package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	
	router.GET("/", Index)
	router.GET("/images", ListImages)
	router.GET("/images/:id", GetImage)
	router.POST("/images", CreateImage)
	router.DELETE("/images/:id", DeleteImage)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Welcome to the Image Blog!")
}

func ListImages(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Listing all images...")
}

func GetImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	fmt.Fprintf(w, "Getting image with ID: %s\n", id)
}

func CreateImage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Creating a new image...")
}

func DeleteImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	fmt.Fprintf(w, "Deleting image with ID: %s\n", id)
}