package main

import (
	"bytes"
	"os"
    "fmt"

	"golang.org/x/net/html"
)

func checkAttribute(n *html.Node, attr string, value string) bool {
    for _, a := range n.Attr {
        if a.Key == attr && a.Val == value {
            return true
        }
    }
    
    return false
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
        fmt.Println(n)
        html.Render(&b, n)
    }

    return  b.String()
}

func saveToFile(name string, text string) {
	err := os.WriteFile(name, []byte(text), 0644)
	if err != nil {
		panic(err)
	}
}
