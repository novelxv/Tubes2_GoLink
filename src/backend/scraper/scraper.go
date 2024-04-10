package scraper

import (
	"github.com/gocolly/colly"
	"fmt"
	"strings"
)

type Link struct {
	Name string
	Url string
}

// Printing the Link Struct
func PrintLink(links []Link){
	for _,item := range links{
		fmt.Printf("Title: %s, Url: %s \n", item.Name,item.Url)
	}
}

// Concatenating the input into a url and change space into underscore
func StringToWikiUrl(name string) (string){
	url := "https://en.wikipedia.org/wiki/" + strings.ReplaceAll(name," ","_")
	return url
}

// Link scrapper
func Scraper(linkName string ) ([]Link, error){
	c := colly.NewCollector()
	
	var links []Link
	
	// Only selects a element with title attributes
	c.OnHTML("a[title]",func(h *colly.HTMLElement) {
		item := Link {}
		item.Name = h.Attr("title")
		item.Url = "https://en.wikipedia.org" + h.Attr("href") // Concatenating into string
		links = append(links, item)
	})
	
	// // When first requested
    // c.OnRequest(func(r *colly.Request) {
    //     fmt.Println("Visiting", r.URL)
    // })

	// // When received a response
    // c.OnResponse(func(r *colly.Response) {
    //     fmt.Println("Got a response from", r.Request.URL)
    // })

	// // When encountering an error
    // c.OnError(func(r *colly.Response, e error) {
    //     fmt.Println("Error:", e)
    // })

	// Visiting the link and scraping
    err := c.Visit(linkName)
    if err != nil {
        return nil, err
    }

    return links, nil
}


