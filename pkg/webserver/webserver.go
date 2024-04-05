package webserver

import (
	"html/template"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Define the data to be passed into the template
	data := struct {
		Title   string
		Message string
	}{
		Title:   "Welcome",
		Message: "This is a simple example of using ParseFiles in Go.",
	}

	// Parse the template file
	tmpl, err := template.ParseFiles("pkg/webserver/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with the data and write the result to the response writer
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Start() {

	http.HandleFunc("/", IndexHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("http can't start", err)
	}
}
