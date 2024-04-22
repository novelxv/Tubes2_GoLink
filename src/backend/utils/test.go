package main

import(
	"log"
	"github.com/angiekierra/Tubes2_GoLink/scraper"
)

func main() {
	linkName := scraper.StringToWikiUrl("Joko Widodo") 
	links, err := scraper.Scraper(linkName)
	if err != nil {
		log.Fatal("Error scraping:", err)
	}
	scraper.PrintLink(links) 
}
