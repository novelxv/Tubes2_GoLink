package scraper

import (
	"fmt"
	"regexp"
	"strings"
	"time"
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

// link title containing on unused title
func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// Link scrapper
func Scraper(linkName string) ([]Link, error) {
	c := colly.NewCollector(
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	notUsed := [...]string{
		"Visit the main page [z]",
		"Guides to browsing Wikipedia",
		"Articles related to current events",
		"Visit a randomly selected article [x]",
		"Learn about Wikipedia and how it works",
		"Guidance on how to use and edit Wikipedia",
		"Learn how to edit Wikipedia",
		"The hub for editors",
		"A list of recent changes to Wikipedia [r]",
		"Add images or other media for use on Wikipedia",
		"Search Wikipedia [f]",
		"A list of edits made from this IP address [y]",
		"Discussion about edits from this IP address [n]",
		"View the content page [c]",
		"Discuss improvements to the content page [t]",
		"List of all English Wikipedia pages containing links to this page [j]",
		"Recent changes in pages linked from this page [k]",
		"Upload files [u]",
		"A list of all special pages [q]",
	}

	var links []Link

	urlPattern := regexp.MustCompile(`^/wiki/..*`)

	// Only selects a element with title attributes
	c.OnHTML("div#mw-content-text a[title]", func(h *colly.HTMLElement) {
		link := h.Attr("href")

		if urlPattern.MatchString(link) {
			item := Link{}
			item.Name = h.Attr("title")
			if !contains(notUsed[:], item.Name) {
				item.Url = "https://en.wikipedia.org" + link
				links = append(links, item)	
			}
		}
	})

	// When first requested
    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL)
    })

	// When received a response
    c.OnResponse(func(r *colly.Response) {
        fmt.Println("Got a response from", r.Request.URL)
    })

	// When encountering an error
    c.OnError(func(r *colly.Response, e error) {
        fmt.Println("Error:", e)
    })

	// Visiting the link and scraping
    err := c.Visit(linkName)
    if err != nil {
        return nil, err
    }

    return links, nil
}