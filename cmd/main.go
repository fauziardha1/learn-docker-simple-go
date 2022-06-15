package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// get port from environment variable
	name := os.Getenv("NAME")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		// get value from environment variable
		fmt.Fprintf(w, "Hello, %s!", name)
	})

	http.ListenAndServe(":8080", nil)

}
