package link

import (
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
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

func buildLink(n *html.Node) Link {
	var r Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			r.Href = attr.Val
			break
		}
	}
	r.Text = "TODO: Parse the text..."

	return r
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
