package grabber

import (
	"io"
	"net/http"
	"os"
)

func main() {
	// resp, err := http.Get("https://example.com")
	resp, err := http.Get("https://uk.glosbe.com/en/uk/apple")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// html, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(html))

	file, err := os.Create("page2.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}
}
