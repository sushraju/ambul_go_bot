package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// TwitterAPI struct def.
type TwitterAPI struct {
	TwitterAPIKey         string `json:"api_key"`
	TwitterAPISecKey      string `json:"api_secret_key"`
	TwitterAccessToken    string `json:"access_token"`
	TwitterAccessTokenSec string `json:"access_token_secret"`
	client                *twitter.Client
}

// NewsAPI struct def.
type NewsAPI struct {
	NewsAPIKey  string `json:"api_key"`
	NewsSources string `json:"sources"`
	httpClient  http.Client
}

// botConfig uses twitterAPI and newsAPI
type botConfig struct {
	TwitterConfig TwitterAPI `json:"auth"`
	NewsAPIConfig NewsAPI    `json:"news"`
}

// InitializeNewsBot with twitter and news api config
func (bc *botConfig) InitializeNewsBot() error {

	jsonConfigName, err := os.Hostname()
	data, err := ioutil.ReadFile(jsonConfigName + ".json")

	if err == nil {
		err = json.Unmarshal(data, bc)
		if err == nil {
			config := oauth1.NewConfig(bc.TwitterConfig.TwitterAPIKey, bc.TwitterConfig.TwitterAPISecKey)
			token := oauth1.NewToken(bc.TwitterConfig.TwitterAccessToken, bc.TwitterConfig.TwitterAccessTokenSec)
			httpClient := config.Client(oauth1.NoContext, token)

			// Twitter client
			bc.TwitterConfig.client = twitter.NewClient(httpClient)
		}
	}

	return err
}

// UpdateTwitterStatus sets the status tweet with the URL
func (bc *botConfig) UpdateTwitterStatus(tweet string) error {
	_, _, err := bc.TwitterConfig.client.Statuses.Update(tweet, nil)

	return err
}
