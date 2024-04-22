package tree

import (
	"fmt"	
	"github.com/angiekierra/Tubes2_GoLink/scraper"
)

type Tree struct {
	Value    string // the link name
	Visited  bool   // check whether a node has already been visited or not
	Children []*Tree
}

// function to make a New Node from string
func NewNode(value string) *Tree {
	return &Tree{
		Value:    value,
		Children: []*Tree{},
	}
}

// function to make a New Node from link
func (n *Tree) NewNodeLink(link []scraper.Link) {
	for _, link := range link {
		temp := NewNode(link.Name)
		n.AddChild(temp)
	}
}

// function to add child to the parent
func (n *Tree) AddChild(child *Tree) {
	n.Children = append(n.Children, child)
}

// fuunction to make node visited
func (n *Tree) AddVisitedNode() {
	n.Visited = true
}

// function to check if the current value is the same with the goal
func IsGoalFound(s1 string, s2 string) bool {
	return s1 == s2
}

/* IDS */

// print only the visited node
func (n *Tree) PrintTreeIds() {
	if n.Visited {
		// print the current value
		fmt.Printf("%s", n.Value)

		if len(n.Children) > 0 {
			fmt.Printf("ANAKK (")

			// print all the children
			visitedChildren := make([]*Tree, 0)
			// saving only the visited children
			for _, child := range n.Children {
				if child.Visited {
					visitedChildren = append(visitedChildren, child)
				}
			}

			for i, child := range visitedChildren {
				child.PrintTreeIds()
				if i < len(visitedChildren)-1 {
					fmt.Printf(", ")
				}
			}

			fmt.Printf(")")
		} else {
			fmt.Printf(" ()")
		}
	}
}

// function to search the word goal recursively in IDS
func (n *Tree) SearhForGoal(found *bool, goal string, level int) bool {
	if level>=0 {
		*found = IsGoalFound(n.Value, goal) // check the current value with the goal
		fmt.Printf("%s \n", n.Value)
		n.AddVisitedNode() // make the node visited

		// Goal founded
		if *found {
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
			if(level>0){
				child.SearhForGoal(found, goal, level-1)
			}
			if *found {
				break // break when the goal was found
			}
		}

	}

	return *found
}

/* BFS */

// // function to print the tree using BFS
// func (n *Tree) PrintTreeBfs() {
// 	queue := []*Tree{n}
// 	for len(queue) > 0 {
// 		current := queue[0]
// 		queue = queue[1:]
// 		if current.Visited {
// 			fmt.Printf("%s", current.Value)
// 			queue = append(queue, current.Children...)
// 		}
// 	}
// 	fmt.Println()
// }

// // function to search the word goal recursively in BFS
// func (n *Tree) SearhForGoalBfs(goal string, stats *golink.GoLinkStats) bool {
// 	queue := []*Tree{n}
// 	for len(queue) > 0 {
// 		current := queue[0]
// 		queue = queue[1:]
// 		if !current.Visited {
// 			current.Visited = true
// 			stats.AddTraversed()
// 		}

// 		fmt.Printf("%s \n", current.Value)
// 		stats.AddChecked()
		
// 		if isGoalFound(current.Value, goal) {
// 			fmt.Print("Found!!\n")
// 			return true
// 		}
		
// 		linkName := scraper.StringToWikiUrl(current.Value)
// 		links, _ := scraper.Scraper(linkName)
// 		current.NewNodeLink(links)

// 		for _, child := range current.Children {
// 			if !child.Visited {
// 				queue = append(queue, child)
// 			}
// 		}
// 	}
// 	return false
// }