package main

import (
	// "log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/angiekierra/Tubes2_GoLink/bfs"
	"github.com/angiekierra/Tubes2_GoLink/golink"
	"github.com/angiekierra/Tubes2_GoLink/ids"
	// "github.com/angiekierra/Tubes2_GoLink/scraper"
	// "github.com/angiekierra/Tubes2_GoLink/tree"
)

type InputData struct {
	StartLink string `json:"startLink"`
	EndLink   string `json:"endLink"`
	UseToggle bool   `json:"useToggle"`
}

type ResponseData struct {
	Articles [][]string      `json:"articles"`
	ArticlesVisited   int  `json:"articlesVisited"`
	ArticlesSearched int   `json:"articlesSearched"`
	TimeNeeded time.Duration		`json:"timeNeeded"`	
}

// testing
func main() {
	r := gin.Default()

	r.Use(cors.Default())

	r.POST("/api/input", func(c *gin.Context) {
		var input InputData


		// / Bind the request body to the inputData struct
        if err := c.BindJSON(&input); err != nil {
            c.JSON(400, gin.H{"error": "Invalid request payload"})
            return
        }
		
		// initialize the inputs
        startLink := input.StartLink
		endLink := input.EndLink
		toggle := input.UseToggle

		stats := golink.NewGoLinkStats()

		if (!toggle){
			bfsStats := bfs.Bfsfunc(startLink,endLink)
			stats.Route = bfsStats.Route
			stats.LinksChecked = bfsStats.LinksChecked
			stats.LinksTraversed = bfsStats.LinksTraversed
			stats.Runtime = bfsStats.Runtime
		} else {
			idsStats := ids.Idsfunc(startLink,endLink)
			stats.Route = idsStats.Route
			stats.LinksChecked = idsStats.LinksChecked
			stats.LinksTraversed = idsStats.LinksTraversed
			stats.Runtime = idsStats.Runtime
		}
		
		

		c.JSON(200, ResponseData{
			Articles:          stats.Route,
			ArticlesVisited:  stats.LinksChecked,
			ArticlesSearched:  stats.LinksTraversed,
			TimeNeeded:        stats.Runtime,
		})


	})
	r.Run() 
}
