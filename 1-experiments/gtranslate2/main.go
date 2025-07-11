package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Translation struct {
	Text string `json:"translatedText"`
}

type Data struct {
	Translations []Translation `json:"translations"`
}

type Response struct {
	Data Data `json:"data"`
}

func main() {
	// expression := "do it now!"
	expression := "never again!"
	key := ""
	url := "https://translation.googleapis.com/language/translate/v2?key=" + key

	data := map[string]interface{}{
		"q":      expression,
		"source": "en",
		"target": "uk",
		"format": "text",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(body))

	var resDetails Response
	err = json.Unmarshal(body, &resDetails)
	if err != nil {
		panic(err)
	}

	result := resDetails.Data.Translations[0].Text
	fmt.Println(result)
}
