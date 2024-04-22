package ids

import (
	"fmt"
	"time"
	"github.com/angiekierra/Tubes2_GoLink/scraper"
	"github.com/angiekierra/Tubes2_GoLink/tree"
	"github.com/angiekierra/Tubes2_GoLink/golink"
)


// main IDS function
func Idsfunc(value string, goal string) *golink.GoLinkStats {
	// save the start time
	startTime := time.Now()
	
	level := 1 // note the current level
	found := false // is the goal alreay found

	// wiki link
	linkName := scraper.StringToWikiUrl(value)
	links, _ := scraper.Scraper(linkName)

	// save the root
	root := tree.NewNode(value)
	root.NewNodeLink(links)
	stats := golink.NewGoLinkStats(root)

	// add the root into the main route
	root.AddMainRoute()

	// do an iteration per level until the goal was found
	for !found {
		fmt.Printf("%d\n", level)
		// searching for goal
		found = SearhForGoal(root, &found, goal, level, stats)
		level++ // goal not found go to the next iteration
	}

	// goal found and set the runtime
	stats.SetRuntime(time.Since(startTime))

	if found {
		PrintTreeIds(root)
		stats.PrintStats()
	}

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
func SearhForGoal(n *tree.Tree, found *bool, goal string, level int, stats *golink.GoLinkStats) bool {
	// do recursive until reached level = 0
	if level>=0 { 
		stats.AddTraversed() // add node to traversed

		if !n.Visited {
			stats.AddChecked() // add node to checked
			n.AddVisitedNode() // add node to visited
			*found = tree.IsGoalFound(n.Value, goal) // check the current value with the goal
		}

		fmt.Printf("%s \n", n.Value)
	
		// Goal founded
		if *found {
			n.AddMainRoute() // add the goal into the main route
			fmt.Print("Found!!\n")
			return true
		}

		// search for another level if not founded yet
		if len(n.Children) == 0 {
			linkName := scraper.StringToWikiUrl(n.Value)
			links, _ := scraper.Scraper(linkName)
			n.NewNodeLink(links)
		}

		// if goal not found yet, go to the next children
		for _, child := range n.Children {
			n.AddMainRoute() // add node to be the temporary main route
			if(level>0){
				SearhForGoal(child, found, goal, level-1, stats)
			}
			if *found {
				break // break when the goal was found
			}
			n.UndoMainRoute() // goal was not founded undo from the main route
		}

	}

	return *found
}