package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello, fa no s3 wagyimi na suro mbaa! </h1>")
	log.Printf("Received request for path: %s", r.URL.Path)
}

func main() {
	var addr string = ":8180"
	handler := http.HandlerFunc(HelloServer)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Could not listen on port %s %v", addr, err)
	}
}
