package main

import (
	"log"
	"os"
)

func main() {
	bc := botConfig{}
	err := bc.InitializeNewsBot()

	if err != nil {
		log.Fatal("Error in initializing bot %s", err.Error)
	}

	file, err := os.OpenFile("ambul_bot.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error setting up log: ", err)
	}

	log.SetOutput(file)
	defer file.Close()

	newsArticles := new(NewsArticles)
	newsArticles, err = bc.NewsAPIConfig.GetEverything()

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
