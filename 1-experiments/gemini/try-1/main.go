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
    // The client gets the API key from the environment variable `GEMINI_API_KEY`.
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	model := "gemini-2.5-flash-lite"
	result, err := client.Models.GenerateContent(
		ctx,
		model,
		genai.Text("One-sentence joke, please."),
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Text())
}
