package main

import (
	"os"
	"strings"

	"golang.org/x/net/html"
)

func parseSlovnyk() {
	file, err := os.Open("slovnyk.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		panic(err)
	}

	// fmt.Println(doc)

	blocks := collectNodes(doc, Tag{"div", "class", "toggle-content"})
	// unwrapTags(blocks[0], "a")
	// fmt.Println(blocks)
	// fmt.Println(nodesToHtml(blocks))
	var sb strings.Builder
	for _, b := range blocks {
		unwrapTags(b, "a")
		sb.WriteString(nodeToHtml(b))
	}

	saveToFile("parsed-s1.html", sb.String())
}
