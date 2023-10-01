package utility

import (
	"encoding/json"
	"fmt"
	"os"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Intro struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

func ParseJson() map[string]Intro {
	file, err := os.ReadFile("./gopher.json")
	if err != nil {
		fmt.Println("error reading file", err)
	}
	parsedFile := make(map[string]Intro)
	if err := json.Unmarshal(file, &parsedFile); err != nil {
		fmt.Println("error parsing json ", err)
	}

	return parsedFile
}
