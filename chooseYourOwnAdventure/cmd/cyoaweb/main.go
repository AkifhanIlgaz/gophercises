package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AkifhanIlgaz/gophercises/chooseYourOwnAdventure/cyoa"
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

	h := cyoa.NewHandler(story)
	addr := fmt.Sprintf(":%d", *port)

	fmt.Printf("Starting the server on port: %d", *port)
	log.Fatal(http.ListenAndServe(addr, h))
}
