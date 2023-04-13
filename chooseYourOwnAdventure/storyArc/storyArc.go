package storyarc

import (
	"encoding/json"
	"os"
)

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func GetStoryArcs(storyFilePath string) (map[string]Arc, error) {
	storyFile, err := os.ReadFile(storyFilePath)
	if err != nil {
		return nil, err
	}

	var storyArc map[string]Arc

	err = json.Unmarshal(storyFile, &storyArc)
	if err != nil {
		return nil, err
	}

	return storyArc, nil
}
