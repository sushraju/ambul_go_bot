package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// init bot
	bc := botConfig{}
	err := bc.InitializeNewsBot()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing bot: %v\n", err)
		os.Exit(1)
	}

	// set loggers
	f, err := os.OpenFile("ambul_bot.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error setting up log: %v\n", err)
		os.Exit(1)
	}

	log.SetOutput(f)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	defer f.Close()

	// fetch news articles
	newsArticles := new(NewsArticles)
	newsArticles, err = bc.NewsAPIConfig.GetEverything()

	// update twitter status with article urls
	if err != nil {
		log.Fatal("Error in fetching articles: ", err)
	} else {
		for _, article := range newsArticles.Articles {
			log.Println("Twitter status set with ", article.URL)
			err = bc.UpdateTwitterStatus(article.URL)
			if err != nil {
				log.Println("Error in updating the status in twitter: ", err)
			}
		}
	}
}
