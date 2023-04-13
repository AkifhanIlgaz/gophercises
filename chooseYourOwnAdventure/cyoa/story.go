package cyoa

import (
	"encoding/json"
	"os"
)

type Story map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text,omitempty"`
	Arc  string `json:"arc,omitempty"`
}

func GetChapters(storyFilePath string) (Story, error) {
	storyFile, err := os.ReadFile(storyFilePath)
	if err != nil {
		return nil, err
	}

	var storyArc Story

	err = json.Unmarshal(storyFile, &storyArc)
	if err != nil {
		return nil, err
	}

	return storyArc, nil
}
