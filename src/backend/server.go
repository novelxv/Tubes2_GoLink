package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors" 
	"github.com/angiekierra/Tubes2_GoLink/scraper"
	"log"
	// "fmt"
	// "net/http"
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
	TimeNeeded float64		`json:"timeNeeded"`	
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
		
        startLink := input.StartLink
		
		
		linkName := scraper.StringToWikiUrl(startLink) 
		links, err := scraper.Scraper(linkName)
		if err != nil {
			log.Fatal("Error scraping:", err)
		}
		scraper.PrintLink(links) 
		// Send a response back to the frontend

		// testing
		articlesVisited := 2
		articlesSearched := 100

		var timeNeeded float64 = 3.14
		articles := [][]string{
			{"apple", "banana", "orange"},
			{"grape", "melon", ""},
		}

		c.JSON(200, ResponseData{
			Articles:          articles,
			ArticlesVisited:   articlesVisited,
			ArticlesSearched:  articlesSearched,
			TimeNeeded:        timeNeeded,
		})


	})
	r.Run() 
}
