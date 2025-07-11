package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	fmt.Println("Start?")
	c := colly.NewCollector(
		colly.AllowedDomains("www.scrapingcourse.com", "uk.glosbe.com", "jisho.org"),
		// colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
		// 	"AppleWebKit/537.36 (KHTML, like Gecko) "+
		// 	"Chrome/122.0.0.0 Safari/537.36"),
	)

	// c.OnRequest(func(r *colly.Request) {
	// 	r.Headers.Set("Accept", "text/html,application/xhtml+xml")
	// 	r.Headers.Set("Accept-Language", "en-US,en;q=0.9")
	// 	r.Headers.Set("Connection", "keep-alive")
	// 	r.Headers.Set("Upgrade-Insecure-Requests", "1")
	// 	r.Headers.Set("Referer", "https://example.com")
	// })

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Here goes response!")
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	// c.OnHTML("li", func(e *colly.HTMLElement) {
	// 	// fmt.Println(e.Text)
	// 	fmt.Println(e.DOM.Html())
	// 	// url := e.ChildAttr("a", "href")
	// 	// fmt.Println(url)
	// })

	// c.Visit("https://www.scrapingcourse.com/ecommerce")

	// c.OnHTML("div.w-1/2 dir-aware-pr-1", func(e *colly.HTMLElement) {
	// c.OnHTML("div.py-2", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Text)
	// 	// fmt.Println(e.DOM.Html())
	// 	// url := e.ChildAttr("a", "href")
	// 	// fmt.Println(url)
	// })

	// c.Visit("https://uk.glosbe.com/en/uk/apple")

	c.OnHTML("#primary", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		// fmt.Println(e.DOM.Html())
		// url := e.ChildAttr("a", "href")
		// fmt.Println(url)
	})

	c.Visit("https://jisho.org/search/%E5%BF%83")
}
