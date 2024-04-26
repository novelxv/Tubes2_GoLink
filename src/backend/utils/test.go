package utils

import (
	// "log"
	// "time"
	// "fmt"
	// "github.com/gin-contrib/cors"
	// "github.com/gin-gonic/gin"

	"github.com/angiekierra/Tubes2_GoLink/bfs"
	"github.com/angiekierra/Tubes2_GoLink/golink"
	"github.com/angiekierra/Tubes2_GoLink/ids"
	// "github.com/angiekierra/Tubes2_GoLink/scraper"
	// "github.com/angiekierra/Tubes2_GoLink/tree"
)

func mainC() {
	
		
	var stats *golink.GoLinkStats

	toggle := true
	startLink := "Joko Widodo"
	endLink := "General officer"

	if (toggle){
		stats = bfs.Bfsfunc(startLink,endLink,true)
	} else {
		stats = ids.Idsfunc(startLink,endLink,true)
	}
		
	stats.PrintStats()

	
}

// import(
// 	"log"
// 	"github.com/angiekierra/Tubes2_GoLink/scraper"
// )

// func main() {
// 	linkName := scraper.StringToWikiUrl("Joko Widodo") 
// 	links, err := scraper.Scraper(linkName)
// 	if err != nil {
// 		log.Fatal("Error scraping:", err)
// 	}
// 	scraper.PrintLink(links) 
// }
