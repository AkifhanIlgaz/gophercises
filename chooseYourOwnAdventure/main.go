package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	storyArc "github.com/AkifhanIlgaz/gophercises/chooseYourOwnAdventure/storyArc"
)

var storyArcs map[string]storyArc.Arc

func main() {
	storyPath := flag.String("story", "./gopher.json", "Path to the story file")
	if storyPath == nil {
		fmt.Errorf("Story file path is required")
		os.Exit(1)
	}
	flag.Parse()

	st, err := storyArc.GetStoryArcs(*storyPath)
	if err != nil {
		fmt.Errorf("Unable to read story file: %v", err)
	}

	storyArcs = st

	mux := http.NewServeMux()
	mux.HandleFunc("/", arcHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mux)

}

func arcHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	chosenArc := r.URL.Path
	arc := storyArcs[chosenArc[1:]]
	if chosenArc == "/" {
		arc = storyArcs["intro"]
	}
	// send response with html template
	fmt.Fprint(w, arc)
}
