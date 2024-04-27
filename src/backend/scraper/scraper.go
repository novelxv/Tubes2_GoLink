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

/* Link Struct */
type Link struct {
	Name string
}

/* Variabel untuk menyimpan cache */
var (
	LinkCache = make(map[string][]Link)
	cacheMutex = sync.Mutex{}
)

/* Print Link */
func PrintLink(links []Link) {
	for _, item := range links {
		fmt.Printf("Title: %s", item.Name)
	}
}

/* String to Wiki URL */
func StringToWikiUrl(name string) string {
	url := "https://en.wikipedia.org/wiki/" + strings.ReplaceAll(name, " ", "_")
	return url
}

/* URL to String */
func UrlToString(url string) string {
	title := strings.TrimPrefix(url, "https://en.wikipedia.org/wiki/")
	title = strings.ReplaceAll(title, "_", " ")
	return title
}

/* Scraper Main Function */
func Scraper(linkName string) ([]Link, error) {
	// Cek cache terlebih dahulu
	cacheMutex.Lock()
	if links, found := LinkCache[linkName]; found {
		cacheMutex.Unlock()
		// fmt.Println("Cache hit for:", linkName) // debug
		return links, nil
	}
	cacheMutex.Unlock()

    // Jika tidak ada di cache, lakukan scraping
	c := colly.NewCollector()
	var links []Link
    // Regex untuk mengecek apakah link valid
	urlPattern := regexp.MustCompile(`^/wiki/..*`)
	avoid := regexp.MustCompile(`^/wiki/(Special:|Talk:|User:|Portal:|Wikipedia:|File:|Category:|Help:|Template:|Template_talk:).*`)
    
    // Batasi domain yang di-scrape
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
        // fmt.Println("Visiting", r.URL) // debug
    })

	// When received a response
    c.OnResponse(func(r *colly.Response) {
        // fmt.Println("Got a response from", r.Request.URL) // debug
    })

	// When encountering an error
    c.OnError(func(r *colly.Response, e error) {
        // fmt.Println("Error:", e) // debug
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

/* BFS Scrapper dari start link */
func BfsScrapper(startLink string) ([]Link, error) {
    visited := make(map[string]bool)
    queue := []string{startLink}
    var links []Link
    var mu sync.Mutex
	defer func() {
		err := SaveToJSON("final.json")
		if err != nil {
			fmt.Println("Error saving data to JSON:", err)
		}
	}()
	
	count := 0
    for len(queue) > 0 {
        // simpan data ke JSON setiap 100 link
		if count % 100 == 0 {
			err := SaveToJSON("testing.json")
			if err != nil {
				fmt.Println("Error saving data to JSON:", err)
			}
		}
        // limit jumlah link yang di-scrape
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

/* Simpan link yang sudah di-scrape ke dalam file JSON */
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


/* Load file JSON ke dalam LinkCache */
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

/* Print Link Cache */
func PrintLinkCache() {
    for k, v := range LinkCache {
        fmt.Println(k, ":", v)
    }
}

// Untuk melakukan scraping, uncomment kode di bawah ini

// func main() {
//     // Attempt to load previously saved data
//     // if err := LoadFromJSON("testing.json"); err != nil {
//     //     fmt.Println("Starting from scratch, error loading data:", err)
//     // }

//     links, err := BfsScrapper(StringToWikiUrl("Inauguration of Joko Widodo"))
//     if err != nil {
//         fmt.Println("Error scraping links:", err)
//     }

//     fmt.Println(links)
// }
