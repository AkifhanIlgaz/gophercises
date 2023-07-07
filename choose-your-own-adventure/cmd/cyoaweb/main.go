package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/AkifhanIlgaz/gophercises/choose-your-own-adventure/cyoa"
)

func main() {
	port := flag.Int("port", 8080, "The port to start the CYOA web application on")
	fileName := flag.String("file", "gopher.json", "Path to the story file")
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JSONGetStory(f)
	if err != nil {
		panic(err)
	}

	storyPrefixTmpl := template.Must(template.New("").Parse(storyTemplate))
	storyPrefixHandler := cyoa.NewHandler(story, cyoa.WithPathFunc(pathFn), cyoa.WithTemplate(storyPrefixTmpl))

	defaultHandler := cyoa.NewHandler(story)

	mux := http.NewServeMux()
	mux.Handle("/story/", storyPrefixHandler)
	mux.Handle("/", defaultHandler)

	addr := fmt.Sprintf(":%d", *port)
	fmt.Printf("Starting the server on port: %d", *port)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story/" || path == "/story" {
		path = "/story/intro"
	}

	return path[len("/story/"):]
}

const storyTemplate = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
      <p>{{.}}</p>
      {{end}}

      <ul>
        {{range .Options}}
        <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
      </ul>
    </section>
    <style>
      body {
        font-family: Arial, Helvetica, sans-serif;
      }
      h1 {
        text-align: center;
        position: relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #fffcf6;
        border: 1px solid #eeee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }

      li {
        padding-top: 10px;
      }

      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }

      a:active,
      a:hover {
        color: #7792a2;
      }

      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>
`
