package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// NewsAPIOptions struct for options sent to the NewsApi endpoint
type NewsAPIOptions struct {
	Sources  string
	From     string
	To       string
	Language string
	SortBy   string
	Page     int
}

// NewsArticles list of Articles
type NewsArticles struct {
	Articles []Article `json:"articles"`
}

// Article URLs
type Article struct {
	URL string `json:"url"`
}

// GetEverything is used to fetch the news
func (na *NewsAPI) GetEverything() (*NewsArticles, error) {
	if len(na.NewsAPIKey) == 0 {
		log.Fatal("News API access has not been configured.")
	}

	const (
		NewsAPIEndPoint = "https://newsapi.org/v2/everything"
		APIKeyReqHeader = "X-Api-Key"
	)

	rand.Seed(time.Now().UnixNano())
	var (
		sourcesList = strings.Split(na.NewsSources, string(","))
		sourcesLen  = len(sourcesList)
		sources     = sourcesList[rand.Intn(sourcesLen-0)] + string(',') + sourcesList[rand.Intn(sourcesLen-0)] + string(',') + sourcesList[rand.Intn(sourcesLen-0)]
	)

	log.Println("Sources to fetch news from: ", sources)

	time.LoadLocation("America/Los_Angeles")
	var (
		dt       = time.Now().Format("2006-01-02")
		lang     = "en"
		sortBy   = "relevancy"
		numPages = 1
	)

	options := NewsAPIOptions{
		Sources:  sources,
		From:     dt,
		To:       dt,
		Language: lang,
		SortBy:   sortBy,
		Page:     numPages,
	}

	// make URL arguments from options given above
	urlArgs := makeURLArgs(options)

	// request and response here
	url := fmt.Sprintf("%s?%s", NewsAPIEndPoint, urlArgs.Encode())
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add(APIKeyReqHeader, na.NewsAPIKey)
	resp, err := na.httpClient.Do(req)

	if err != nil {
		log.Println("Error while making a request: ", err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		log.Printf("Error while requesting %s", url)
		log.Printf("Response code %s ", resp.Status)
		return nil, errors.New(resp.Status)
	}

	defer resp.Body.Close()

	// instantiate NewsArticles for unmarshalling json data
	newsArticles := new(NewsArticles)

	err = json.NewDecoder(resp.Body).Decode(newsArticles)

	if err != nil {
		log.Println("Error while decoding json: ", err)
		return nil, err
	}

	return newsArticles, nil
}

// makeURLArgs converts options to URL args
func makeURLArgs(options NewsAPIOptions) url.Values {
	urlArgs := url.Values{}
	values := reflect.ValueOf(options)

	for i := 0; i < values.NumField(); i++ {
		argName := values.Type().Field(i).Name
		argValue := values.Field(i)
		argType := argValue.Kind()

		switch argType {

		case reflect.String:
			if len(argValue.String()) > 0 {
				urlArgs.Add(strings.ToLower(string(argName)), argValue.String())
			}

		case reflect.Int:
			intString := strconv.FormatInt(argValue.Int(), 10)
			if len(intString) > 0 {
				urlArgs.Add(strings.ToLower(string(argName)), intString)
			}
		}
	}

	return urlArgs
}
