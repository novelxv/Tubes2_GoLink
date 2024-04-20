package bfs

import (	
	"github.com/angiekierra/Tubes2_GoLink/scraper"
	"github.com/angiekierra/Tubes2_GoLink/tree"
)


// main BFS function
func Bfsfunc(value string, goal string) *tree.Tree {
	// wiki link
	linkName := scraper.StringToWikiUrl(value)
	links, _ := scraper.Scraper(linkName)

	// save the root
	root := tree.NewNode(value)
	root.NewNodeLink(links)

	// use BFS to search for the goal
	found := root.SearhForGoalBfs(goal)

	if found {
		root.PrintTreeBfs()
	}

	return root // return tree
}
