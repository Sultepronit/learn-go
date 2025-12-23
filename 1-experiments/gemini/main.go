package main

// go get google.golang.org/genai

import (
    "context"
    "fmt"
    "log"

    "google.golang.org/genai"
)

func main() {
    ctx := context.Background()
    client, err := genai.NewClient(ctx, &genai.ClientConfig{
		HTTPOptions: genai.HTTPOptions{APIVersion: "v1"},
	})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(client)
}