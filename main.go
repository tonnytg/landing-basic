package main

import (
	"fmt"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "teste")
}

func main() {
	log.Println("start webserver")

	http.HandleFunc("/", IndexHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("http can't start", err)
	}
}
