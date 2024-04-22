package bfs

import (	
	"fmt"
	"time"
	"github.com/angiekierra/Tubes2_GoLink/scraper"
	"github.com/angiekierra/Tubes2_GoLink/tree"
	"github.com/angiekierra/Tubes2_GoLink/golink"
)


// main BFS function
func Bfsfunc(value string, goal string) *golink.GoLinkStats {
	startTime := time.Now()
	
	// wiki link
	linkName := scraper.StringToWikiUrl(value)
	links, _ := scraper.Scraper(linkName)

	// save the root
	root := tree.NewNode(value)
	root.NewNodeLink(links)
	stats := golink.NewGoLinkStats(root)

	// use BFS to search for the goal
	found := SearhForGoalBfs(root, goal, stats)

	elapsedTime := time.Since(startTime)
	stats.SetRuntime(elapsedTime)

	if found {
		PrintTreeBfs(root)
		stats.PrintStats()
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

// function to search the word goal recursively in BFS
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
		stats.AddChecked()
		
		if tree.IsGoalFound(current.Value, goal) {
			fmt.Print("Found!!\n")
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
