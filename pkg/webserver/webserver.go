package webserver

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type Post struct {
	Title   string
	Content string
}

var templatesDir = "pkg/webserver/templates/"

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
	// tmpl := template.Must(template.ParseFiles("pkg/webserver/templates/main.html"))
	// tmpl, err := template.ParseFiles("pkg/webserver/templates/index.html")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// Obtém todos os arquivos .html no diretório de templates
	files, err := filepath.Glob(filepath.Join(templatesDir, "*.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse os templates
	tmpl := template.Must(template.ParseFiles(files...))

	// Execute the template with the data and write the result to the response writer
	// if err := tmpl.Execute(w, data); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	// err := tmpl.Execute(w, data)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	err = tmpl.ExecuteTemplate(w, "main.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	// Define a blog post
	post := Post{
		Title:   "Hello, World!",
		Content: "This is my first blog post. Welcome to my blog!",
	}

	// Parse the post template file
	tmpl, err := template.ParseFiles("templates/post.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with the blog post data and write the result to the response writer
	if err := tmpl.Execute(w, post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Start() {

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/post", PostHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("http can't start", err)
	}
}
