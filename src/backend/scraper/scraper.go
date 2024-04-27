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
	cacheMutex = sync.Mutex{}
)

// Printing the Link Struct
func PrintLink(links []Link) {
	for _, item := range links {
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

// Link scrapper with cache
func Scraper(linkName string) ([]Link, error) {
	// Cek cache terlebih dahulu
	cacheMutex.Lock()
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

// BfsScrapper is a breadth-first scrapper that scrapes links from a given start link
func BfsScrapper(startLink string) ([]Link, error) {
    visited := make(map[string]bool)
    queue := []string{startLink}
    var links []Link
    var mu sync.Mutex
	defer func() {
		err := SaveToJSON("final28.json")
		if err != nil {
			fmt.Println("Error saving data to JSON:", err)
		}
	}()
	
	count := 0
    for len(queue) > 0 {
		if count % 100 == 0 {
			err := SaveToJSON("testing.json")
			if err != nil {
				fmt.Println("Error saving data to JSON:", err)
			}
		}
		if (count == 10000){
			return links,nil
		}
        currentLink := queue[0]
        queue = queue[1:]

        // Check if the link has been visited before
        mu.Lock()
        if visited[currentLink] {
            mu.Unlock()
            continue
        }
        visited[currentLink] = true
        mu.Unlock()

        // Check if the link exists in the LinkCache
        cacheMutex.Lock()
        _, exists := LinkCache[currentLink]
        cacheMutex.Unlock()
        if exists {
            continue
        }

        // Scrape the current link
        scrapedLinks, err := Scraper(currentLink)
        if err != nil {
            fmt.Println("Error scraping link:", err)
        }

        // Append the scraped links to the result
        mu.Lock()
        links = append(links, scrapedLinks...)
        mu.Unlock()

        // Append the scraped links to the LinkCache
        cacheMutex.Lock()
        LinkCache[currentLink] = scrapedLinks
        cacheMutex.Unlock()

        // Enqueue newly discovered links to the queue
        mu.Lock()
        for _, l := range scrapedLinks {
            if !visited[l.Name] {
                queue = append(queue, StringToWikiUrl(l.Name))
            }
        }
        mu.Unlock()
		count++
    }

    return links, nil
}

// SaveToJSON saves the LinkCache map to a JSON file
func SaveToJSON(filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "    ")
    return encoder.Encode(LinkCache)
}


// LoadFromJSON loads the LinkCache map from a JSON file
func LoadFromJSON(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    err = decoder.Decode(&LinkCache)
    if err != nil {
        return err
    }

    return nil
}

// PrintLinkCache prints the contents of the LinkCache map
func PrintLinkCache() {
    for k, v := range LinkCache {
        fmt.Println(k, ":", v)
    }
}

// func main() {
//     // Attempt to load previously saved data
//     // if err := LoadFromJSON("testing.json"); err != nil {
//     //     fmt.Println("Starting from scratch, error loading data:", err)
//     // }

//     links, err := BfsScrapper(StringToWikiUrl("Chicken"))
//     if err != nil {
//         fmt.Println("Error scraping links:", err)
//     }

//     fmt.Println(links)
// }
