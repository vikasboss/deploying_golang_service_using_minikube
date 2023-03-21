package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		fmt.Print("Just doing testing\n")
		fmt.Fprintf(w, "Hello, world!")
	})

	http.ListenAndServe(":8080", nil)
}
