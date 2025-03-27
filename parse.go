package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="...">) in an HTML
// document.
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and will return a
// slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	// 1. find <a> nodes in document
	nodes := linkNodes(doc)
	for _, node := range nodes {
		fmt.Println(node)
	}
	// 2. for search link node...
	// 2.a build a Link
	// 3. return the Links
	return nil, nil
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var r []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		r = append(r, linkNodes(c)...)
	}
	return r
}
