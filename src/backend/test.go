// package main

// import (
// 	"fmt"

// 	"github.com/angiekierra/Tubes2_GoLink/bfs"
// 	// "github.com/angiekierra/Tubes2_GoLink/ids"
// 	// "github.com/angiekierra/Tubes2_GoLink/ids"
// 	"github.com/angiekierra/Tubes2_GoLink/golink"
// 	"github.com/angiekierra/Tubes2_GoLink/scraper"
// )

// func main() {
// 	// load the json file
// 	scraper.LoadFromJSON("./scraper/final2.json")
// 	scraper.LoadFromJSON("./scraper/final.json")
// 	scraper.LoadFromJSON("./scraper/testing.json")
// 	// scraper.LoadFromJSON("./scraper/1.json")
// 	// scraper.LoadFromJSON("./scraper/2.json")
// 	// scraper.PrintLinkCache()
// 	fmt.Println("Cache is loaded")

// 	// Call idsfunc to create the tree
// 	var test *golink.GoLinkStats

// 	// test = bfs.Bfsfunc("Chicken", "Duck", false)
// 	test = bfs.Bfsfunc("Inauguration of Joko Widodo", "Indonesia", false)
// 	// test = bfs.Bfsfunc("Inauguration of Joko Widodo", "Joko Widodo", true)
// 	// test = bfs.Bfsfunc("Joko Widodo", "Philosophy", false)
// 	// test = bfs.Bfsfunc("Joko Widodo", "Philosophy", true)
// 	// test = bfs.Bfsfunc("Bandung Institute of Technology", "Philosophy", false)
// 	// test = ids.Idsfunc("Joko Widodo", "Inauguration of Joko Widodo", false)
// 	// test = bfs.Bfsfunc("Genetic", "Human", false)
// 	// test = bfs.Bfsfunc("Joko Widodo", "Inauguration of Joko Widodo", false)
// 	// test = bfs.Bfsfunc("Joko Widodo", "Prabowo Subianto", false)

// 	// Print the tree structure using the PrintTreeIds method
// 	// test.PrintStats()

// 	// test = bfs.Bfsfunc("Tetangga Masa Gitu?", "Situation Comedy (album)")
// 	// test = bfs.Bfsfunc("Joko Widodo", "Jusuf Kalla")
// 	test = bfs.Bfsfunc("Joko Widodo", "Indonesia", false)
// 	// test = bfs.Bfsfunc("Vector", "Mathematics", false)
// 	// test = bfs.Bfsfunc("Vector", "Euclidean vector")
// 	test.PrintStats()

// }

