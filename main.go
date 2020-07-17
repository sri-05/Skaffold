package main

import (
	"fmt"
	"net/http"
	"os"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello "+os.Getenv("NAME")+" Welcome to Skaffold Page")
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Version is V0.0.1")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/Version", versionHandler)
	http.ListenAndServe(":8080", nil)
}
