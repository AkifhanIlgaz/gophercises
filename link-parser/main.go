package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AkifhanIlgaz/gophercises/link-parser/parser"
)

func main() {
	/*
		TODO:
		1. File flag to specify the file to parse
			Accept multiple files
			Accept links
	*/
	file := flag.String("file", "tests/ex1.html", "File to parse")
	flag.Parse()

	f, err := os.Open(*file)
	if err != nil {
		fmt.Errorf("unable to open file: %v", err)
	}

	links, err := parser.Parse(f)
	if err != nil {
		fmt.Errorf("unable to parse file: %v", err)
	}

	fmt.Printf("%+v", links)

}
