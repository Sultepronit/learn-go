package main

import (
	"context"
	"fmt"
	"log"

	translate "cloud.google.com/go/translate/apiv3"
	"cloud.google.com/go/translate/apiv3/translatepb"
)

func main() {
	// 1. Параметри
	projectID := "" // Вставте свій ID
	location := "global"                  // Зазвичай використовується global
	text := "Hello, how are you?"
	targetLang := "uk"

	// 2. Виклик функції перекладу
	result, err := translateText(projectID, location, text, targetLang)
	if err != nil {
		log.Fatalf("Помилка перекладу: %v", err)
	}

	fmt.Println(result)
}

// Функція з документації (адаптована)
func translateText(projectID string, location string, text string, targetLanguageCode string) (string, error) {
	ctx := context.Background()
	client, err := translate.NewTranslationClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	req := &translatepb.TranslateTextRequest{
		Parent:             fmt.Sprintf("projects/%s/locations/%s", projectID, location),
		TargetLanguageCode: targetLanguageCode,
		Contents:           []string{text},
		MimeType:           "text/plain", // Можна змінити на text/html
	}

	resp, err := client.TranslateText(ctx, req)
	if err != nil {
		return "", err
	}

	// Повертаємо перший переклад із списку
	return resp.GetTranslations()[0].GetTranslatedText(), nil
}
