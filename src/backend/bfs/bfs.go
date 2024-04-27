package bfs

import (	
	"fmt"
	"time"
	"sync"
	"context"
	"github.com/angiekierra/Tubes2_GoLink/scraper"
	"github.com/angiekierra/Tubes2_GoLink/tree"
	"github.com/angiekierra/Tubes2_GoLink/golink"
)


// main BFS function
func Bfsfunc(value string, goal string, multisol bool) *golink.GoLinkStats {
	startTime := time.Now()

	// save the root
	root := tree.NewNode(value)
	stats := golink.NewGoLinkStats()

	var found bool

	// use BFS to search for the goal
	if (multisol){
		found = SearchForGoalBfsMTMS(root, goal, stats)
	} else {
		found = SearchForGoalBfsMT(root, goal, stats)
	}

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
func SearhForGoalBfs(n *tree.Tree, goal string, stats *golink.GoLinkStats) bool {
	queue := []*tree.Tree{n}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if !current.Visited {
			current.Visited = true
			stats.AddTraversed()
		}
		
		fmt.Printf("%s \n", current.Value)
		
		if tree.IsGoalFound(current.Value, goal) {
			fmt.Print("Found!!\n")
			route := tree.GoalRoute(current)
			for i := 0; i < len(route)-1; i++ {
				stats.AddChecked()
			}
			stats.AddRoute(route)
			return true
		}
		
		mu.Lock()
		linkName := scraper.StringToWikiUrl(current.Value)
		links, _ := scraper.Scraper(linkName)
		current.NewNodeLink(links)
		mu.Unlock()

		for _, child := range current.Children {
			if !child.Visited {
				queue = append(queue, child)
			}
		}
	}
	return false
}

func SearchForGoalBfsMT(root *tree.Tree, goal string, stats *golink.GoLinkStats) bool {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	nodeQueue := make(chan *tree.Tree, 10) 

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
			stats.AddTraversed()

			if tree.IsGoalFound(node.Value, goal) {
				fmt.Println("Found!!")
				route := tree.GoalRoute(node)
				for i := 0; i < len(route)-1; i++ {
					stats.AddChecked()
				}
				stats.AddRoute(route)
				cancel() // Notify to cancel all operations
				return
			}

			mu.Lock()
			links, _ := scraper.Scraper(scraper.StringToWikiUrl(node.Value))

			const maxGoroutines = 10000 
			sem := make(chan bool, maxGoroutines)

			// if len(links) > 100 {
			// 	links = links[:100]
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


				if !child.Visited && child.GetDepth() <= 9 {
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
	return ctx.Err() == context.Canceled
}

func SearchForGoalBfsMTMS(root *tree.Tree, goal string, stats *golink.GoLinkStats) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	var wg sync.WaitGroup
	nodeQueue := make(chan *tree.Tree, 10)
	
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
			stats.AddTraversed()

			if tree.IsGoalFound(node.Value, goal) {
				fmt.Println("Found!!")
				route := tree.GoalRoute(node)
				for i := 0; i < len(route)-1; i++ {
					stats.AddChecked()
				}
				stats.AddRoute(route)
				found = true
				solutionDepth = len(route) - 1
				// cancel() // Notify to cancel all operations
				// return
			}

			if solutionDepth != 0 && node.GetDepth() > solutionDepth {
				cancel()
				return
			}

			mu.Lock()
			links, _ := scraper.Scraper(scraper.StringToWikiUrl(node.Value))

			const maxGoroutines = 1500 
			sem := make(chan bool, maxGoroutines)

			if len(links) > 100 {
				links = links[:100]
			}

			sem = make(chan bool, maxGoroutines)

			visitedNodes := make(map[string]struct{})

			for _, link := range links {
				if _, ok := visitedNodes[link.Name]; ok {
					continue
				}
				child := tree.NewNode(link.Name)
				node.AddChild(child)
				visitedNodes[link.Name] = struct{}{}


				if !child.Visited && child.GetDepth() <= 9 {
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