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

	// Створюємо клієнт. 
	// Якщо ви зробили `export GOOGLE_APPLICATION_CREDENTIALS`, 
	// SDK сам знайде ваші права без жодних ключів у коді!
	// client, err := genai.NewClient(ctx, option.WithCredentialsFile("keys.json")) // Або без опції, якщо є export
    client, err := genai.NewClient(ctx) // Або без опції, якщо є export
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