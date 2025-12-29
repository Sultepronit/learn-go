package main

import (
	"fmt"
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

func parseJisho() {
	file, err := os.Open("jisho.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		panic(err)
	}

	// doc, err := grab2("https://jisho.org/search/%E5%BF%83", false)
	// if err != nil {
	// 	panic(err)
	// }

	// node := findTheNode(doc)
	// // cleanHtml(node)
	// html := nodeToHtml(node)

	// saveToFile("parsed-j0.html", html)

	blocks := collectNodes(doc, Tag{"div", "class", "concept_light"})
	fmt.Println(blocks)
	for _, block := range blocks {
		removeTags(block, []Tag{
			{"div", "class", "concept_light-status"},
			{"a", "class", "light-details_link"},
		})
	}

	// fmt.Println(nodesToHtml(blocks))
	saveToFile("parsed-j3.html", nodesToHtml(blocks))
}
