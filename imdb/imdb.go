package imdb

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type IMDbClientConfig struct {
	API    string
	APIKey string
}

type IMDbClient struct {
	config  *IMDbClientConfig
	request *http.Client
}

const (
	API = "https://imdb-api.com/API"
)

func NewClient(config *IMDbClientConfig) (client *IMDbClient) {
	if config.API == "" {
		config.API = API
	}
	request := http.DefaultClient
	client = &IMDbClient{config, request}
	return
}

func (client *IMDbClient) Request(method string, url string, body io.Reader) ([]byte, error) {
	log.Println(url)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "imdb-client")
	res, err := client.request.Do(req)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(res.Body)
}

func (client *IMDbClient) Call(api string, params any) ([]byte, error) {
	url := client.config.API + fmt.Sprintf("%s/%s/%s", api, client.config.APIKey, params)
	return client.Request("GET", url, nil)
}
