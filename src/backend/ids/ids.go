package ids

import (
	"fmt"
	"sync"
	"time"
	"github.com/angiekierra/Tubes2_GoLink/scraper"
	"github.com/angiekierra/Tubes2_GoLink/tree"
	"github.com/angiekierra/Tubes2_GoLink/golink"
)

var (
	mu sync.Mutex
	found bool = false
)

// main IDS function
func Idsfunc(value string, goal string) *golink.GoLinkStats {
	// save the start time
	startTime := time.Now()
	
	level := 1 // note the current level

	// wiki link
	linkName := scraper.StringToWikiUrl(value)
	links, _ := scraper.Scraper(linkName)

	// save the root
	root := tree.NewNode(value)
	root.NewNodeLink(links)
	stats := golink.NewGoLinkStats()

	// add the root into the main route
	root.AddMainRoute()

	// do an iteration per level until the goal was found
	for !found{
		fmt.Printf("%d\n", level)
		// searching for goal
		SearchForGoal(root, goal, level, stats)
		level++ // goal not found, go to the next iteration
		
		if level==10 || found{
			break
		}
	}

	// goal found and set the runtime
	stats.SetRuntime(time.Since(startTime))

	if found {
		PrintTreeIds(root)
		stats.PrintStats()
	}

	found = false

	return stats
}

// print only the visited node
func PrintTreeIds(n *tree.Tree) {
	if n.Visited {
		fmt.Printf("%s", n.Value)

		if len(n.Children) > 0 {
			fmt.Printf("ANAKK (")

			for i, child := range n.Children {
				if child.Visited {
					PrintTreeIds(child)
					if i < len(n.Children)-1 {
						fmt.Printf(", ")
					}
				}
			}

			fmt.Printf(")")
		} else {
			fmt.Printf(" ()")
		}
	}
}

// function to search the word goal recursively in IDS
// function to search the word goal recursively in IDS
func SearchForGoal(node *tree.Tree, goal string, currentLevel int, stats *golink.GoLinkStats) {
	if currentLevel < 0 || found {
		return
	}

	var wg sync.WaitGroup
    defer wg.Wait()
	
	if !node.Visited && !found{
		stats.AddChecked()
		node.AddVisitedNode()
		found = tree.IsGoalFound(node.Value, goal)
	}
	fmt.Printf("%s \n", node.Value)

	if tree.IsGoalFound(node.Value, goal) {
		node.AddMainRoute()
		route := tree.GoalRoute(node)
		stats.AddRoute(route)
		fmt.Println("Found!!")
		return
	}

	if len(node.Children) == 0 && !found{
		mu.Lock()
		linkName := scraper.StringToWikiUrl(node.Value)
		links, _ := scraper.Scraper(linkName)
		node.NewNodeLink(links)
		mu.Unlock()
	}

	// if goal not found yet, go to the next children
	if(currentLevel>0){
		childSemaphore := make(chan struct{}, 1500) 

		for _, child := range node.Children {
			wg.Add(1)
			childSemaphore <- struct{}{}
			go func(childNode *tree.Tree) {
				defer wg.Done()
				defer func() { <-childSemaphore }() 
				SearchForGoal(childNode, goal, currentLevel-1, stats)
			}(child)
			
		}
		wg.Wait()

		if found{
			return
		}
	}

	if found{
		return
	}
}
