package scraper

import (
	"fmt"
	"regexp"
	"strings"
	"github.com/gocolly/colly"
)

type Link struct {
	Name string
	Url  string
}

// Printing the Link Struct
func PrintLink(links []Link) {
	for _, item := range links {
		fmt.Printf("Title: %s, Url: %s \n", item.Name, item.Url)
	}
}

// Concatenating the input into a url and change space into underscore
func StringToWikiUrl(name string) string {
	url := "https://en.wikipedia.org/wiki/" + strings.ReplaceAll(name, " ", "_")
	return url
}


// Link scrapper
func Scraper(linkName string) ([]Link, error) {
	c := colly.NewCollector()


	var links []Link

	urlPattern := regexp.MustCompile(`^/wiki/..*`)
	avoid := regexp.MustCompile(`^/wiki/(Special:|Talk:|User:|Portal:|Wikipedia:|File:|Category:|Help:|Template:|Template_talk:).*`)


	// Only selects a element with title attributes
	c.OnHTML("div#mw-content-text a[title]", func(h *colly.HTMLElement) {
		link := h.Attr("href")

		if (urlPattern.MatchString(link) && !avoid.MatchString(link)) {
			item := Link{}
			item.Name = h.Attr("title")
			links = append(links, item)	
			
		}
	})

	// When first requested
    // c.OnRequest(func(r *colly.Request) {
    //     fmt.Println("Visiting", r.URL)
    // })

	// When received a response
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