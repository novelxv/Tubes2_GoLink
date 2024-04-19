package utils

import(
	"log"
)

func main() {
	linkName := StringToWikiUrl("Fourth Extraordinary Session of the Islamic Summit Conference") 
	links, err := Scraper(linkName)
	if err != nil {
		log.Fatal("Error scraping:", err)
	}
	PrintLink(links) 
}
