package logic

import (
	"fmt"
	"github.com/angiekierra/Tubes2_GoLink/utils"
)

func idsfunc(value string, goal string) *Tree {
	level := 1
	found := false
	linkName := utils.StringToWikiUrl(value)
	links, _ := utils.Scraper(linkName)
	root := NewNode(value)
	root.NewNodeLink(links)

	for !found {
		fmt.Printf("%d\n", level)
		found = root.searchforgoal(false, found, goal)
		root.PrintTreeIds()
		level++
	}

	return root
}

func (n *Tree) searchforgoal(done bool, found bool, goal string) bool {
	found = isGoalFound(n.Value, goal)
	fmt.Printf("%s ", n.Value)
	if len(n.Children) == 0 && !done {
		links, _ := utils.Scraper(n.Value)
		n.NewNodeLink(links)
		done = true
	}
	if !done {
		for _, child := range n.Children {
			done = false
			child.searchforgoal(done, found, goal)
		}
	}
	return done
}

func isGoalFound(s1 string, s2 string) bool {
	return s1 == s2
}
