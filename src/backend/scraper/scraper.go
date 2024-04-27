package scraper

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

type Link struct {
	Name string
}

// Cache untuk menyimpan hasil scraping
var (
	LinkCache = make(map[string][]Link)
	LinkCache = make(map[string][]Link)
	cacheMutex = sync.Mutex{}
)

// Printing the Link Struct
func PrintLink(links []Link) {
	for _, item := range links {
		fmt.Printf("Title: %s", item.Name)
		fmt.Printf("Title: %s", item.Name)
	}
}

// Concatenating the input into a url and change space into underscore
func StringToWikiUrl(name string) string {
	url := "https://en.wikipedia.org/wiki/" + strings.ReplaceAll(name, " ", "_")
	return url
}

func UrlToString(url string) string {
	title := strings.TrimPrefix(url, "https://en.wikipedia.org/wiki/")
	title = strings.ReplaceAll(title, "_", " ")
	return title
}
func UrlToString(url string) string {
	title := strings.TrimPrefix(url, "https://en.wikipedia.org/wiki/")
	title = strings.ReplaceAll(title, "_", " ")
	return title
}

// Link scrapper with cache
func Scraper(linkName string) ([]Link, error) {
	// Cek cache terlebih dahulu
	cacheMutex.Lock()
	if links, found := LinkCache[linkName]; found {
	if links, found := LinkCache[linkName]; found {
		cacheMutex.Unlock()
		fmt.Println("Cache hit for:", linkName)
		return links, nil
	}
	cacheMutex.Unlock()

	c := colly.NewCollector()
	var links []Link

	urlPattern := regexp.MustCompile(`^/wiki/..*`)
	avoid := regexp.MustCompile(`^/wiki/(Special:|Talk:|User:|Portal:|Wikipedia:|File:|Category:|Help:|Template:|Template_talk:).*`)

	c.OnHTML("div#mw-content-text a[title]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		if (urlPattern.MatchString(link) && !avoid.MatchString(link)) {
			item := Link{}
			item.Name = h.Attr("title")
			links = append(links, item)	
			
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

	err := c.Visit(linkName)
	if err != nil {
		fmt.Println("Error scraping link:", err)
	}

	// Simpan hasil ke cache sebelum return
	cacheMutex.Lock()
	LinkCache[linkName] = links
	cacheMutex.Unlock()

	return links, nil
}
