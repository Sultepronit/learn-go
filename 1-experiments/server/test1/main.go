package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
		fmt.Fprintf(w, "Query: %s\n", r.URL.RawQuery)
		fmt.Fprintf(w, "expression: %s\n", r.URL.Query().Get("expression"))
		fmt.Println(r.Method, r.URL.Path)
	})

	fmt.Println("Listening port 8080?")
	http.ListenAndServe(":8080", nil)
}
