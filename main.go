package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRoot)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
