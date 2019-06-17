package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is neat and simple!")
}

func main() {
	fmt.Println("Starting go server at port:8000")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
