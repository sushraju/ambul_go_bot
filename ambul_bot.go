package main

import (
	"fmt"
	"log"
)

func main() {
	bc := botConfig{}
	err := bc.InitializeNewsBot()

	if err != nil {
		log.Fatal("Error in initializing bot %s", err.Error)
	}

	newsArticles := new(NewsArticles)
	newsArticles, err = bc.NewsAPIConfig.GetEverything()

	if err != nil {
		log.Fatal("Error in fetching articles %s", err.Error)
	} else {
		for _, article := range newsArticles.Articles {
			err = bc.UpdateTwitterStatus(article.URL)
			if err != nil {
				fmt.Println("Error in updating the status in twitter %s", err.Error)
			}
		}
	}
}
