package parser

import (
	"fmt"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) []Link {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		fmt.Println(err)

	}

	links := []Link{}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			href = ""
		}

		text := strings.Join(strings.Fields(s.Text()), " ")

		links = append(links, Link{href, text})

	})

	return links
}
