package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/angiekierra/Tubes2_GoLink/bfs"
	"github.com/angiekierra/Tubes2_GoLink/golink"
	"github.com/angiekierra/Tubes2_GoLink/ids"
	"github.com/angiekierra/Tubes2_GoLink/scraper"

	"log"

)

type InputData struct {
    StartLink   string `json:"startLink"`
    EndLink     string `json:"endLink"`
    UseToggle   bool   `json:"useToggle"`
    IsChecked bool   `json:"isChecked"` 
}

type ResponseData struct {
	Articles [][]string      `json:"articles"`
	ArticlesVisited   int  `json:"articlesVisited"`
	ArticlesSearched int   `json:"articlesSearched"`
	TimeNeeded int64		`json:"timeNeeded"`	
}

// testing
func main() {

	scraper.LoadFromJSON("./scraper/final.json")
	
	r := gin.Default()

	r.Use(cors.Default())

	r.POST("/api/input", func(c *gin.Context) {
		var input InputData


		// / Bind the request body to the inputData struct
        if err := c.BindJSON(&input); err != nil {
            c.JSON(400, gin.H{"error": "Invalid request payload"})
            return
        }
		
		//initialize the inputs
        startLink := input.StartLink
		endLink := input.EndLink
		toggle := input.UseToggle
		check := input.IsChecked

		log.Println(startLink,endLink,toggle,check)

	
		var stats *golink.GoLinkStats
		
		if (!check) {
			if (toggle){
				stats = bfs.Bfsfunc(startLink,endLink,false)
			} else {
				stats = ids.Idsfunc(startLink,endLink,false)
			}
		} else {
			if (toggle){
				stats = bfs.Bfsfunc(startLink,endLink,true)
			} else {
				stats = ids.Idsfunc(startLink,endLink,true)
			}
		}
		
		runTime := stats.Runtime.Milliseconds()
		

		c.JSON(200, ResponseData{
			Articles:          stats.Route,
			ArticlesVisited:  stats.LinksChecked,
			ArticlesSearched:  stats.LinksTraversed,
			TimeNeeded:        runTime,
		})


	})
	r.Run() 
}
