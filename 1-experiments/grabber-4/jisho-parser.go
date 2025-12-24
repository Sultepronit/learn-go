package main

import (
	// "bytes"
	"os"

	"golang.org/x/net/html"
)

func findTheNode(n *html.Node) *html.Node {
	// n.RemoveChild()
	if n.Data == "div" {
		// fmt.Println(n.Type, n.Data, n.Attr)
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == "primary" {
				return n
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.ElementNode {
			continue
		}

		r := findTheNode(c)
		if r != nil {
			return r
		}
	}

	return nil
}

// func toHtml(n *html.Node) string {
// 	var b bytes.Buffer
// 	html.Render(&b, n)
// 	s := b.String()
// 	// fmt.Println(s)
// 	return s

// 	// for _, node := range foundNodes {
// 	// 	html.Render(&buf, node)
// 	// }
// }

// func saveToFile(name string, text string) {
// 	err := os.WriteFile(name, []byte(text), 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func parseJisho() {
	file, err := os.Open("jisho.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// fmt.Println(file)

	doc, err := html.Parse(file)
	if err != nil {
		panic(err)
	}

	// fmt.Println(doc)

	node := findTheNode(doc)
	// cleanHtml(node)
	html := toHtml(node)

	saveToFile("parsed1.html", html)
}
