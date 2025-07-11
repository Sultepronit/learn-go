package main

// go get golang.org/x/net/html

import (
	"bytes"
	"os"

	"golang.org/x/net/html"
)

func findNode(n *html.Node) *html.Node {
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

		r := findNode(c)
		if r != nil {
			return r
		}
	}

	return nil
}

func removeMe(n *html.Node, c *html.Node) *html.Node {
	// next := c.NextSibling
	prev := c.PrevSibling
	n.RemoveChild(c)
	// return next
	return prev
}

// func sliceToSet(list []string) map[string]struct{} {
//     result := make(map[string]struct{})
//     for _, item := range list {
//         result[item] = struct{}{}
//     }
//     return result
// }

var removeByAttrS = [...]string{
	"concept_light-status",
	"light-details_link",
}

var removeByAttr = map[string]struct{}{
	"concept_light-status": {},
	"light-details_link":   {},
}

func cleanHtml(p *html.Node) {
	for n := p.FirstChild; n != nil; n = n.NextSibling {
		switch n.Type {
		case html.CommentNode:
			n = removeMe(p, n)
			continue
		case html.ElementNode:
			if n.Data == "h4" {
				n = removeMe(p, n)
				continue
			} else {
				// fmt.Println(c.Type, c.Data, c.Attr)
				var newAttrs []html.Attribute
				toRemove := false
				for _, a := range n.Attr {
					// if a.Key == "id" && a.Val == "primary"
					for i := range removeByAttrS {
						if a.Val == removeByAttrS[i] {
							toRemove = true
							break
						}
					}
					if toRemove {
						break
					}
				}
				if toRemove {
					n = removeMe(p, n)
					continue
				} else {
					// n.Attr = newAttrs
				}

				cleanHtml(n)
			}
		}
	}
}

func toHtml(n *html.Node) string {
	var b bytes.Buffer
	html.Render(&b, n)
	s := b.String()
	// fmt.Println(s)
	return s

	// for _, node := range foundNodes {
	// 	html.Render(&buf, node)
	// }
}

func saveToFile(name string, text string) {
	err := os.WriteFile(name, []byte(text), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
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

	node := findNode(doc)
	cleanHtml(node)
	html := toHtml(node)

	saveToFile("parsed5.html", html)
}
