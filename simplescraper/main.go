package main

import (
	"simplescraper/api"
	"simplescraper/scraper"
	"time"
)

func runScraper() {

	for {
		scraper.Scrap()
		time.Sleep(30 * time.Minute)
	}

}

func main() {
	// go scraper.Scrap()
	api.InitAPI()
}
