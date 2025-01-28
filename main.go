package main

import (
	"fmt"
	"log"
	"net/http"
)

var globalCounter int

// This is the first attempt at creating a go web server
// I'm using this as a understanding exercise
// before I attempt to build an actual audio streaming server

func main() {
	http.HandleFunc("/counter", counterServer)
	fmt.Println("Server running on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func counterServer(w http.ResponseWriter, r *http.Request) {
	globalCounter++
	fmt.Fprintf(w, "Counter: %d", globalCounter)
}
