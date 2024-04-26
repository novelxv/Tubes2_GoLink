package ids

import (
	"fmt"
	"sync"
	"time"
	"context"
	"github.com/angiekierra/Tubes2_GoLink/scraper"
	"github.com/angiekierra/Tubes2_GoLink/tree"
	"github.com/angiekierra/Tubes2_GoLink/golink"
)

var (
	mu sync.Mutex
	found bool = false
	single bool = false
)

// main IDS function
func Idsfunc(value string, goal string, multisol bool) *golink.GoLinkStats {
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// do an iteration per level until the goal was found
	for !found{
		fmt.Printf("%d\n", level)
		// searching for goal
		if (multisol) { // search for multiple solution
			SearchForGoal(ctx,root, goal, level, stats, false)
		} else { // search for only one solution
			SearchForGoal(ctx,root, goal, level, stats, true)
		}
		level++ // goal not found, go to the next iteration
		
		if level==10 || found{ // break if found or reach level 10
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
	single = false

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
func SearchForGoal(ctx context.Context, node *tree.Tree, goal string, currentLevel int, stats *golink.GoLinkStats, onesolution bool) {
	select {
		case <-time.After(5*time.Minute):
			found = true
			return
		case <-ctx.Done():
			return
		default:
	}

	// recursive base: stop until the currentlevel=0 or found the goal (if single solution)
	if currentLevel < 0 || single { 
		return
	}

	var wg sync.WaitGroup
    defer wg.Wait()
	
	// check the goal if not found yet and make the node to visited
	if !node.Visited && !found{
		stats.AddChecked()
		node.AddVisitedNode()
		found = tree.IsGoalFound(node.Value, goal)
		if found && onesolution {
			single = true
		}
	}

	fmt.Printf("%s \n", node.Value)

	// add the goal to main route
	if tree.IsGoalFound(node.Value, goal) {
		node.AddMainRoute()
		route := tree.GoalRoute(node)
		stats.AddRoute(route)
		fmt.Println("Found!!")
		return
	}

	// if node doesnt have a children do scrapping
	if len(node.Children) == 0 && currentLevel>0 {
		mu.Lock() // lock when getting a request
		linkName := scraper.StringToWikiUrl(node.Value)
		links, _ := scraper.Scraper(linkName)
		node.NewNodeLink(links)
		mu.Unlock()
	}

	// if goal not found yet, go to the next children
	if(currentLevel>0){
		childSemaphore := make(chan struct{}, 10000) // do coccurent process

		for _, child := range node.Children {
			wg.Add(1)
			childSemaphore <- struct{}{}
			go func(childNode *tree.Tree) {
				defer wg.Done()
				defer func() { <-childSemaphore }() 
				SearchForGoal(ctx, childNode, goal, currentLevel-1, stats, onesolution)
			}(child)
			
		}
		wg.Wait()

		if found{ // goal found
			return
		}
	}

	if found{ // goal found
		return
	}
}
