package cyoa

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func init() {
	templ = template.Must(template.New("").Parse(defaultHandleTemplate))
}

var templ *template.Template

var defaultHandleTemplate = `
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
        <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
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

type HandleOption func(h *handler)

func WithTemplate(t *template.Template) HandleOption {
	return func(h *handler) {
		h.t = t
	}
}

func NewHandler(s Story, opts ...HandleOption) http.Handler {
	h := handler{s, templ}
	for _, opt := range opts {
		opt(&h)
	}

	return h
}

type handler struct {
	s Story
	t *template.Template
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}

	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Chapter not found.", http.StatusNotFound)
	}

}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text,omitempty"`
	Chapter string `json:"arc,omitempty"`
}

func JSONGetStory(r *os.File) (Story, error) {
	d := json.NewDecoder(r)
	var storyArc Story

	if err := d.Decode(&storyArc); err != nil {
		return nil, err
	}

	return storyArc, nil
}
