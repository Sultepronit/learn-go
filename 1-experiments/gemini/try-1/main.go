package main

// go get google.golang.org/genai

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/genai"
)

func main() {
	model := "gemini-2.5-flash-lite"
	model = "gemini-flash-latest"
	instruction := "Input should somehow hint at some japanese kanji. Generate suggestions of kanji that: A) contains same or similar element(s); B) is result of adding/removing/replacing some element(s) to/from/of the input; C) looks similar. Just answer, no additional text."
	instruction = "Input somehow hints at some japanese kanji. Even if it contains several signs they all are representing a single one! Generate suggestions of kanji that: A) contains same or similar element(s); B) is result of adding/removing/replacing some element(s) to/from/of the input; C) looks similar."
	instruction = "You are helping someone with limited knowledge of japanese. He has a kanji that somehow resembles him the input he provides. Generate suggestions of kanji that: A) contains same or similar element(s); B) is result of adding/removing/replacing some element(s) to/from/of the input; C) looks similar."
	instruction = "You are helping someone with limited knowledge of japanese. They have a kanji they don't know. But the input they provide: a) represents the elements/radicals (same or just similar) the kanji consists of; b) is kanji that looks similar to the needed. Try to combine the input if it's more than one character (all of the hints should be used!). Example: 心門文' -> 憫. Try to add/remove/replace some elements of the input, and/or suggest visually similar kanji, if the input is a single character. Example: 周 <-> 彫; 殴 <-> 投."
    instruction = "You are helping someone with limited knowledge of japanese. They have a kanji they don't know, but they do visually(!) associate it with something they provide as the input. Example: 心門文' -> 憫; 亦心 -> 恋; 周 <-> 彫; 殴 <-> 投; 疲 -> 痩. Please provide most suitable suggestions. Only the list of kanji to chose from, no additional text!"
    instruction = "You are helping someone with limited knowledge of Japanese. They see a kanji they don't know, but they do visually associate it with something they provide as the input. Examples of input-> expected: 心門文 -> 憫; 亦心 -> 恋; 道ホ -> 述; 周 <-> 彫; 殴 <-> 投; 疲 <-> 痩. Please, provide rather long list of suggestions. Guessing is ok. No explanations, just the list of kanji to chose from."
	
    req := ""
	// req := "心門文"
	// req = "亦心"
    // req = "木公"
    // req = "湖"
    // req = "論"
    req = "格"
    req = "巻" // 圏
    // req = "秋" // 称
    // req = "秋ホ" // 称

	ctx := context.Background()
	// The client gets the API key from the environment variable `GEMINI_API_KEY`.
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	config := &genai.GenerateContentConfig{
		SystemInstruction: &genai.Content{
			Parts: []*genai.Part{
				{Text: instruction},
			},
		},
	}

	result, err := client.Models.GenerateContent(ctx, model, genai.Text(req), config)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Text())
}
