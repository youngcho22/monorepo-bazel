package main

import (
	"log"
	"net/http"

	hello_world "github.com/youngcho22/monorepo-bazel/projects/go_hello_world/src/go/src/instabase/hello_world"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request")
	w.Write([]byte(hello_world.HelloWorld()))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)

	// Bind to a port and pass our router in
	log.Println("Going to listen on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
