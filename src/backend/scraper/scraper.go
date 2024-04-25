package scraper

import (
	"fmt"
	"regexp"
	"strings"
	"sync/atomic"
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
	c := colly.NewCollector()

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
	avoid := regexp.MustCompile(`^/wiki/(Special:|Talk:|User:|Portal:|Wikipedia:|File:|Category:|Help:|Template:|Template_talk:).*`)


	// Only selects a element with title attributes
	c.OnHTML("div#mw-content-text a[title]", func(h *colly.HTMLElement) {
		link := h.Attr("href")

		if (urlPattern.MatchString(link) && !avoid.MatchString(link)) {
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

var shouldContinue int32 = 1 // Atomic flag
// Scrapper with cancellation
func Scraper2(linkName string) ([]Link, error) {
    c := colly.NewCollector(
        colly.Async(true), // Enable asynchronous request
    )
    
    var links []Link
    urlPattern := regexp.MustCompile(`^/wiki/..*`)
    avoid := regexp.MustCompile(`^/wiki/(Special:|Talk:|User:|Portal:|Wikipedia:|File:|Category:|Help:|Template:|Template_talk:).*`)

    c.OnHTML("div#mw-content-text a[title]", func(h *colly.HTMLElement) {
        if atomic.LoadInt32(&shouldContinue) == 0 {
            return // Skip processing if cancellation flag is set
        }
        link := h.Attr("href")
        if urlPattern.MatchString(link) && !avoid.MatchString(link) {
            links = append(links, Link{
                Name: h.Attr("title"),
                Url:  "https://en.wikipedia.org" + link,
            })
        }
    })

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL.String())
        if atomic.LoadInt32(&shouldContinue) == 0 {
            r.Abort() // Abort the request if cancellation flag is set
        }
    })

    c.OnResponse(func(r *colly.Response) {
        fmt.Println("Got a response from", r.Request.URL.String())
    })

    c.OnError(func(r *colly.Response, e error) {
        fmt.Println("Error:", e)
    })

    err := c.Visit(linkName)
    if err != nil {
        return nil, err
    }

    c.Wait() // Wait for all jobs to finish

    return links, nil
}

func CancelScraping() {
    atomic.StoreInt32(&shouldContinue, 0) // Set flag to 0 to indicate stopping
}