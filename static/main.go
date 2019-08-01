package main

import (
	"log"
	"net/http"
)

func main() {
	//create fileserver
	fileserver := http.FileServer(http.Dir("static/"))

	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			fileserver.ServeHTTP(w, r)
		},
	)
	//starting HTTP server on port 8000.
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
