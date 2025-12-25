package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// func checkNode(n *html.Node, tag string) {
// 	return n.Data == tag
// }

func collectNodes(n *html.Node, result []*html.Node, tag string) []*html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.ElementNode {
			continue
		}

		if c.Data == tag {
			result = append(result, c)
		} else {
			result = collectNodes(c, result, tag)
		}
	}

	return result
}

func isHeaderRight(header string, expected string) bool {
	spl := regexp.MustCompile(`[0-9]|\s\(-`)
	header = spl.Split(header, 2)[0]
	repl := strings.NewReplacer("|", "", "́", "")
	header = repl.Replace(header)
	header = strings.TrimSpace(header)
	// fmt.Println("[", header, "]")
	// fmt.Println("[", expected, "]")
	return strings.EqualFold(header, expected)
}

func isArticleMain(article *html.Node, query string) bool {
	b := findNode(article, "b")
	if b == nil {
		return false
	}

	header := getTextContent(b)
	fmt.Println(header)

	return isHeaderRight(header, query)
}

func parseE2u() {
	// file, err := os.Open("e2u-2.html")
	file, err := os.Open("e2u.html")
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

	tds := make([]*html.Node, 0, 5)
	tds = collectNodes(doc, tds, "td")
	// fmt.Println(tds)
	// html := nodesToHtml(tds)
	// fmt.Println(html)

	articles := map[string][]*html.Node{
		"main":    {},
		"other":   {},
		"context": {},
	}
	for _, t := range tds {
		// fmt.Println(i, t)
		if checkAttribute(t, "class", "result_row") {
			articles["context"] = append(articles["context"], t)
		} else if isArticleMain(t, "apple") {
			articles["main"] = append(articles["main"], t)
		} else {
			articles["other"] = append(articles["other"], t)
		}
	}
	// fmt.Println(articles)
	for i, t := range articles {
		fmt.Println(i, t)
	}

	// saveToFile("parsed2.html", html)
}
