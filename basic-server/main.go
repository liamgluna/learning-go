package main

import (
	"log"
	"net/http"
)

func main() {
	srv := http.NewServeMux()

	srv.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	log.Print("Starting server on localhost:4000")

	err := http.ListenAndServe(":4000", srv)
	log.Fatal(err)
}
