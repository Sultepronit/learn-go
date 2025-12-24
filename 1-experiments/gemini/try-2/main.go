package main

// go get cloud.google.com/go/vertexai/genai

import (
	"context"
	"fmt"
	"log"

	// "github.com/google/generative-ai-go/genai"
	// "google.golang.org/api/option"
    "cloud.google.com/go/vertexai/genai"
)

func main() {
	ctx := context.Background()

	projectID := "hazel-lyceum-448211-t3" 
	location := "global"                 
    client, err := genai.NewClient(ctx, projectID, location) 
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Обираємо модель (наприклад, вашу улюблену 3-flash)
	model := client.GenerativeModel("gemini-1.5-flash")

	// Налаштування моделі (температура, ліміти тощо) - це "дорослий" контроль
	model.SetTemperature(0.7)

	// Запит
	resp, err := model.GenerateContent(ctx, genai.Text("Поясни ідіому 'down the road' одним реченням."))
	if err != nil {
		log.Fatal(err)
	}

	// Вивід результату
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
}