package imdb

import "encoding/json"

type SearchResult struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

type SearchResponse struct {
	SearchType string         `json:"searchType"`
	Results    []SearchResult `json:"results"`
}

func (client *IMDbClient) Search(expression string) (output *SearchResponse, err error) {
	res, err := client.Call("/Search", expression)
	if err != nil {
		return
	}
	json.Unmarshal(res, &output)
	return
}

func (client *IMDbClient) SearchTitle(expression string) {
}
