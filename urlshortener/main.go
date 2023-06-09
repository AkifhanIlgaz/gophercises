package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/AkifhanIlgaz/gophercises/urlshortener/urlshort"
)

func main() {
	yamlFile := flag.String("yaml", "", "Path to yaml file of path to urls")
	yaml, err := os.ReadFile(*yamlFile)
	if err != nil {
		fmt.Errorf("Unable to read yaml file")
	}
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
