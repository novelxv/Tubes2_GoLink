package ids

import (
	"fmt"

	"github.com/angiekierra/Tubes2_GoLink/scraper"
	"github.com/angiekierra/Tubes2_GoLink/tree"
)


// main IDS function
func Idsfunc(value string, goal string) *tree.Tree {
	level := 1 // note the current level
	found := false // is the goal alreay found

	// wiki link
	linkName := scraper.StringToWikiUrl(value)
	links, _ := scraper.Scraper(linkName)

	// save the root
	root := tree.NewNode(value)
	root.NewNodeLink(links)

	// do an iteration per level until the goal was found
	for !found {
		fmt.Printf("%d\n", level)
		root.SearhForGoal(&found, goal, level)
		level++
	}

	return root // return tree
}
