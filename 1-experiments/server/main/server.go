package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		if query != "" {
			fmt.Println(r.Method, r.URL.Path, query)
			resp := Gtranslate(query)
			fmt.Fprintln(w, resp)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Listening on port: %s\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
