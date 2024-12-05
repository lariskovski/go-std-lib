package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	
	http.HandleFunc("/", UploadImage)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func UploadImage(w http.ResponseWriter, r *http.Request) {
	var s string

	if r.Method == http.MethodPost {
		// Parse the form data
		r.ParseMultipartForm(10 << 20) // 10 MB
		file, header, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()
		data, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Printf("File name: %s\nFile size: %d bytes\n", header.Filename, len(data))
		s = string(data)
	}

	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<form method="post" enctype="multipart/form-data">
		<input type="file" name="image">
		<input type="submit">
		</form>`+s)
}
