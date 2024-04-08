package main

import(
	"log"
	"github.com/angiekierra/Tubes2_GoLink/utils"
)

func main() {
	linkName := utils.StringToWikiUrl("Fourth Extraordinary Session of the Islamic Summit Conference") 
	links, err := utils.Scraper(linkName)
	if err != nil {
		log.Fatal("Error scraping:", err)
	}

	utils.PrintLink(links) 
}