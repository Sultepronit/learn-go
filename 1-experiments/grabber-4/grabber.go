package main

import (
	"io"
	"net/http"
	"os"
)

func grab(slyMode bool) {
	url := "https://jisho.org/search/%E5%BF%83"
	filename := "jisho.html"

	// sly!
	url = "https://e2u.org.ua/s?w=apple&dicts=all&highlight=on&filter_lines=on"
	filename = "e2u.html"

	url = "https://e2u.org.ua/s?w=apple+pie&dicts=all&highlight=on&filter_lines=on"
	filename = "e2u-2.html"

	url = "https://slovnyk.ua/index.php?swrd=%D1%8F%D0%B1%D0%BB%D1%83%D1%87%D0%BD%D0%B8%D0%B9"
	filename = "slovnyk.html"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	if slyMode {
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		req.Header.Set("Accept-Language", "en-US,en;q=0.9")
		req.Header.Set("Connection", "keep-alive")
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// html, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(html))

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}
}
