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
	StartLink string        `json:"startLink"`
	EndLink   string        `json:"endLink"`
	UseToggle bool          `json:"useToggle"`
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
		endLink := input.EndLink
		useToggle := input.UseToggle
		
		linkName := scraper.StringToWikiUrl(startLink) 
		links, err := scraper.Scraper(linkName)
		if err != nil {
			log.Fatal("Error scraping:", err)
		}
		scraper.PrintLink(links) 
		// Send a response back to the frontend
		c.JSON(200, ResponseData {
			StartLink: startLink,
			EndLink:   endLink,
			UseToggle: useToggle,
		})


	})
	r.Run() 
}
