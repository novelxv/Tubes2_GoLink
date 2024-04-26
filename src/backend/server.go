package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "fmt"
	// "github.com/angiekierra/Tubes2_GoLink/bfs"
	"github.com/angiekierra/Tubes2_GoLink/golink"
	"github.com/angiekierra/Tubes2_GoLink/ids"
	"log"
	// "github.com/angiekierra/Tubes2_GoLink/scraper"
	// "github.com/angiekierra/Tubes2_GoLink/tree"
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
		// janlup ganti functionnya nanti
		if (!check) {
			if (toggle){
				stats = ids.Idsfunc(startLink,endLink)
				log.Println("bfs normal")
			} else {
				stats = ids.Idsfunc(startLink,endLink)
				log.Println("ids normal")
			}
		} else {
			if (toggle){
				stats = ids.Idsfunc(startLink,endLink)
				log.Println("bfs multi")
			} else {
				stats = ids.Idsfunc(startLink,endLink)
				log.Println("ids multi")
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
