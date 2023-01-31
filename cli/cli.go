package cli

import (
	"log"

	"github.com/song940/imdb/imdb"
)

func Run() {
	imdb := imdb.NewClient(&imdb.IMDbClientConfig{
		APIKey: "k_xkqzfxfq",
	})
	result, err := imdb.Search("Saw")
	log.Println(result.SearchType, err)
	for _, result := range result.Results {
		log.Println(result.Title)
	}
}
