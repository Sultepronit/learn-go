package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// n.RemoveChild()

func checkAttribute(n *html.Node, attr string, value string) bool {
	for _, a := range n.Attr {
		if a.Key == attr && a.Val == value {
			return true
		}
	}

	return false
}

func findNode(n *html.Node, tag string) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.ElementNode {
			continue
		}

		if c.Data == tag {
			return c
		} else {
			return findNode(c, tag)
		}
	}

	return nil
}

func getShallowText(n *html.Node) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			fmt.Println(c.Data)
		}
	}
}

func getTextContent(n *html.Node) string {
	var sb strings.Builder

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			switch c.Type {
			case html.TextNode:
				sb.WriteString(c.Data)
			case html.ElementNode:
				traverse(c)
			}
		}
	}
	traverse(n)

	return sb.String()
}

func toHtml(n *html.Node) string {
	var b bytes.Buffer
	html.Render(&b, n)
	s := b.String()
	return s
}

func nodesToHtml(nodes []*html.Node) string {
	var b bytes.Buffer
	for _, n := range nodes {
		html.Render(&b, n)
	}

	return b.String()
}

func saveToFile(name string, text string) {
	err := os.WriteFile(name, []byte(text), 0644)
	if err != nil {
		panic(err)
	}
}
