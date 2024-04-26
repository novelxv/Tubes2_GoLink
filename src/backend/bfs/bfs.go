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
func Bfsfunc(value string, goal string) *golink.GoLinkStats {
	startTime := time.Now()

	// save the root
	root := tree.NewNode(value)
	stats := golink.NewGoLinkStats()

	// use BFS to search for the goal
	found := SearchForGoalBfsMT(root, goal, stats)

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

func BfsfuncM(value string, goal string) *golink.GoLinkStats {
    startTime := time.Now()

    // save the root
    root := tree.NewNode(value)
    stats := golink.NewGoLinkStats()

    // use BFS to search for the goal
    routes := SearchForGoalBfsM(root, goal, stats)

    elapsedTime := time.Since(startTime)
    stats.SetRuntime(elapsedTime)

    if (routes) {
        // for _, route := range routes {
        //     stats.AddRoute(route)
        //     stats.LinksTraversed += len(route)
        // }

        stats.PrintStats()
        // PrintTreeBfs(root)
        // stats.PrintStats()
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
		
		linkName := scraper.StringToWikiUrl(current.Value)
		links, _ := scraper.Scraper(linkName)
		current.NewNodeLink(links)

		for _, child := range current.Children {
			if !child.Visited {
				queue = append(queue, child)
			}
		}
	}
	return false
}

func SearchForGoalBfsM(n *tree.Tree, goal string, stats *golink.GoLinkStats) bool {
	queue := []*tree.Tree{n}
	for len(queue) > 0 {
		if (len(stats.Route) == 2){
			return true;
		}
		current := queue[0]
		queue = queue[1:]
		if !current.Visited {
			current.Visited = true
			stats.AddTraversed()
		}
		
		fmt.Printf("%s \n", current.Value)
		stats.AddChecked()
		
		linkName := scraper.StringToWikiUrl(current.Value)
		links, _ := scraper.Scraper(linkName)
		current.NewNodeLink(links)

		for _, child := range current.Children {
			if  (child.Value == goal){
				route := tree.GoalRoute(child)
				stats.AddRoute(route)
				fmt.Print("Found!!\n")
				stats.PrintStats()

			} else {
				if (child.ParentLength() < 2) && !child.Visited {
					queue = append(queue, child)
				}
			}
		}
	}
	return true
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

			links, _ := scraper.Scraper(scraper.StringToWikiUrl(node.Value))

			const maxGoroutines = 500 
			sem := make(chan bool, maxGoroutines)

			visitedNodes := make(map[string]*tree.Tree)

			for _, link := range links {
				if visitedNodes[link.Name] != nil {
					continue
				}
				child := tree.NewNode(link.Name)
				node.AddChild(child)
				visitedNodes[link.Name] = child

				if !child.Visited {
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
		}
	}()

	wg.Wait() // Wait for all goroutines to complete before finishing
	<-ctx.Done() // Ensure context is canceled
	return ctx.Err() == context.Canceled
}