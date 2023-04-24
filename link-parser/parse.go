package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="...">) in an HTML document.
type Link struct {
	Href string
	Text string
}

// Parse parses r as an HTML document and returns the links in the document.
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

func buildLink(node *html.Node) Link {
	var ret Link

	for _, attr := range node.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}

	ret.Text = text(node)

	return ret

}

func text(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}

	if node.Type != html.ElementNode {
		return ""
	}

	var ret string

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c) + " "
	}

	return strings.Join(strings.Fields(ret), " ")
}

func linkNodes(root *html.Node) []*html.Node {

	if root.Type == html.ElementNode && root.Data == "a" {
		return []*html.Node{root}
	}

	var ret []*html.Node

	for c := root.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}

	return ret
}
