package main

import (
	"fmt"
	"strings"

	"github.com/AkifhanIlgaz/gophercises/link-parser"
)

var exampleHtml = `
<html>
  <body>
    <h1>Hello!</h1>
    <a href="/other-page">A link to another page</a>
    <a href="/second-page">A link to second page</a>
  </body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
