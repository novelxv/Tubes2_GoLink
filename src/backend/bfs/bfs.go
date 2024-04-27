package bfs

import (	
	"fmt"
	"time"
	"sync"
	"context"
	"github.com/angiekierra/Tubes2_GoLink/scraper"
	"github.com/angiekierra/Tubes2_GoLink/tree"
	"github.com/angiekierra/Tubes2_GoLink/golink"

	// _ "net/http/pprof"
	// "net/http"
	// "log"
)

// func init() {
//     go func() {
//         fmt.Println("Starting server for profiling at http://localhost:6060")
//         if err := http.ListenAndServe("localhost:6060", nil); err != nil {
//             log.Fatalf("HTTP server ListenAndServe: %v", err)
//         }
//     }()
// }

// main BFS function
func Bfsfunc(value string, goal string, multisol bool) *golink.GoLinkStats {
	startTime := time.Now()

	// save the root
	root := tree.NewNode(value)
	stats := golink.NewGoLinkStats()

	// use BFS to search for the goal
	found := SearchForGoalBfs(root, goal, stats, multisol)

	elapsedTime := time.Since(startTime)
	stats.SetRuntime(elapsedTime)

	if found {
		stats.PrintStats()
		// PrintTreeBfs(root)
		// stats.PrintStats()
	} else {
		fmt.Println("Goal not found")
	}

	return stats
}

// function to print the tree using BFS
func PrintTreeBfs(n *tree.Tree) {
	queue := []*tree.Tree{n}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.Visited {
			fmt.Printf("|%s", current.Value)
			queue = append(queue, current.Children...)
		}
	}
	fmt.Println()
}

var mu sync.Mutex

// function to search the word goal with BFS
func SearchForGoalBfs(root *tree.Tree, goal string, stats *golink.GoLinkStats, multisol bool) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	var wg sync.WaitGroup
	nodeQueue := make(chan *tree.Tree, 10000000)
	
	found := false
	var solutionDepth int

	wg.Add(1)
	go func() {
		defer wg.Done()
		nodeQueue <- root
	}()

	go func() {
		defer close(nodeQueue) // Ensure channel is closed when all processing is done
		for node := range nodeQueue {
			if node.Visited {
				continue
			}

			node.Visited = true
			stats.AddChecked()

			// if stats.LinksChecked > 2500 {
			// 	cancel()
			// 	return
			// }

			if tree.IsGoalFound(node.Value, goal) {
				fmt.Println("Found!!")
				route := tree.GoalRoute(node)
				stats.AddRoute(route)
				found = true
				solutionDepth = len(route) - 1
				if !multisol {
					cancel() // Notify to cancel all operations
					return
				}
			}

			if solutionDepth != 0 && node.GetDepth() > solutionDepth {
				cancel()
				return
			}

			mu.Lock()
			links, _ := scraper.Scraper(scraper.StringToWikiUrl(node.Value))

			const maxGoroutines = 10000
			sem := make(chan bool, maxGoroutines)

			// if len(links) > 150 {
			// 	links = links[:150]
			// }

			sem = make(chan bool, maxGoroutines)

			visitedNodes := make(map[string]struct{})

			for _, link := range links {
				if _, ok := visitedNodes[link.Name]; ok {
					continue
				}
				child := tree.NewNode(link.Name)
				node.AddChild(child)
				visitedNodes[link.Name] = struct{}{}


				if !child.Visited && child.GetDepth() <= 6 {
					wg.Add(1)
					go func(ch *tree.Tree) {
						sem <- true
						defer func() {
							<-sem
							wg.Done()
						}()
						select {
						case nodeQueue <- ch:
						case <-ctx.Done(): // Handle cancellation
						}
					}(child)
				}
			}
			mu.Unlock()
		}
	}()

	wg.Wait() // Wait for all goroutines to complete before finishing
	<-ctx.Done() // Ensure context is canceled
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Search timed out after 5 minutes")
	} else if ctx.Err() == context.Canceled {
        fmt.Println("Search stopped after reaching a level beyond the first solution")
    }
	return found
}