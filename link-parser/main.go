package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	/*
		TODO:
		1. Accept multiple files
		2. Accept URLS
		3. Add testing for edge cases
	*/

	files := os.Args[1:]

	_, err := url.ParseRequestURI(files[0])
	if err != nil {
		fmt.Println(err)
	}
	
	// f, err := os.Open(files[0])
	// if err != nil {
	// 	fmt.Errorf("unable to open file: %v", err)
	// }

	// links := parser.Parse(f)

	// fmt.Printf("%+v", links)

}
