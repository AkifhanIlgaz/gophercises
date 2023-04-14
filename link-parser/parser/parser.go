package parser

import (
	"fmt"
	"io"

	"github.com/PuerkitoBio/goquery"
)

type Link struct {
	Href string
	Text string
}

// Parse takes path to a HTML document and returns a slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	// Create new document from r using goquery
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		fmt.Println("Unable to create new document from reader: ")
		return nil, err
	}

	// Find all a tags
	links := []Link{}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// For each a tag, get the href and text
		href, exists := s.Attr("href")
		if !exists {
			href = ""
		}
		text := s.Text()

		links = append(links, Link{href, text})

	})

	return links, nil
}
