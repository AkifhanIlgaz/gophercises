package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"

	"github.com/AkifhanIlgaz/gophercises/link-parser"
)

func main() {
	urlFlag := flag.String("url", "https://www.calhoun.io", "the url that you want to build a sitemap for")
	flag.Parse()

	u, err := url.Parse(*urlFlag)
	if err != nil {
		panic(err)
	}

	domain := u.Host

	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	links, err := link.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	for _, link := range links {
		linkUrl, err := url.Parse(link.Href)
		if err != nil {
			panic(err)
		}
		if linkUrl.Host == domain {

			fmt.Println(link.Href)
		} else if linkUrl.Scheme == "" {
			fmt.Println(u.Scheme + "://" + domain + linkUrl.Path)
		}

	}
}
