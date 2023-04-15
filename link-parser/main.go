package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/AkifhanIlgaz/gophercises/link-parser/parser"
)

func main() {
	/*
		TODO:
		1. Accept multiple files
		2. Accept URLS
		3. Add testing for edge cases // Testing for parsing a website
		4. Create a new struct with filename and links
	*/

	for _, source := range os.Args[1:] {
		var sourceReader io.Reader

		if isValidUrl(source) {
			resp, err := http.Get(source)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			sourceReader = resp.Body
		} else {
			f, err := os.Open(source)
			if err != nil {
				fmt.Println(err)
			}
			defer f.Close()
			sourceReader = f
		}

		links := parser.Parse(sourceReader)
		fmt.Printf("%+v\n", links)
	}

}

func isValidUrl(input string) bool {
	_, err := url.ParseRequestURI(input)

	return err == nil
}
