package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/AkifhanIlgaz/gophercises/link-parser"
)

var visitedPages = map[string]bool{}

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")

	flag.Parse()

	bfs(*urlFlag)

	toXml := urlset{
		Xmlns: xmlns,
	}
	for page, _ := range visitedPages {
		toXml.Urls = append(toXml.Urls, loc{page})
	}
	
	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}
	fmt.Println()
}

func bfs(rootPage string) {
	if visitedPages[rootPage] {
		return
	}

	visitedPages[rootPage] = true

	pages := getPages(rootPage)

	for _, p := range pages {
		bfs(p)
	}
}

func getPages(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	return filter(base, hrefs(resp.Body, base))

}

func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)
	var hrefs []string

	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}

	return hrefs
}

func filter(base string, links []string) []string {
	var ret []string
	for _, l := range links {
		if strings.HasPrefix(l, base) {
			ret = append(ret, l)
		}
	}

	return ret
}

// func bfs(rootPage string) {
// 	if visitedPages[rootPage] {
// 		return
// 	}

// 	visitedPages[rootPage] = true

// 	resp, err := http.Get(rootPage)
// 	if err != nil {
// 		fmt.Errorf("cannot visit page: %v", rootPage)
// 		return
// 	}

// 	defer resp.Body.Close()

// 	linkedPages, _ := link.Parse(resp.Body)
// 	filteredPages := filterPagesWithDomain(linkedPages, domain, rootPage)

// 	for _, page := range filteredPages {

// 		bfs(page)
// 	}

// }

// func filterPagesWithDomain(pages []link.Link, domain string, baseUrl string) []string {
// 	pagesWithSameDomain := []string{}

// 	for _, page := range pages {
// 		pageUrl, err := url.Parse(page.Href)
// 		if err != nil {
// 			return nil
// 		}
// 		if pageUrl.Host == domain {
// 			pagesWithSameDomain = append(pagesWithSameDomain, page.Href)
// 		} else if pageUrl.Scheme == "" {

// 			pagesWithSameDomain = append(pagesWithSameDomain, baseUrl+pageUrl.Path)
// 		}
// 	}

// 	return pagesWithSameDomain
// }
