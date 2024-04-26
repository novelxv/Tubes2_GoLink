package tree

import (
	"github.com/angiekierra/Tubes2_GoLink/scraper"
)

type Tree struct {
	Value    string // the link name
	Visited  bool   // check whether a node has already been visited or not
	MainRoute  bool  // main node to the goal
	Parent 	*Tree // save parent
	Children []*Tree // save child
	Depth 	int // save depth
}

// function to make a New Node from string
func NewNode(value string) *Tree {
	return &Tree{
		Value:    value,
		Children: []*Tree{},
		Depth: 0,
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
	child.Parent = n
	child.Depth = n.Depth + 1
	n.Children = append(n.Children, child)
}

// function to get the depth of the node
func (n *Tree) GetDepth() int {
	return n.Depth
}

// fuunction to make node visited
func (n *Tree) AddVisitedNode() {
	n.Visited = true
}

// fuunction to add node to main route
func (n *Tree) AddMainRoute() {
	n.MainRoute = true
}

// fuunction to remove node from main route
func (n *Tree) UndoMainRoute() {
	n.MainRoute = false
}

// function to check if the current value is the same with the goal
func IsGoalFound(s1 string, s2 string) bool {
	return s1 == s2
}

func GoalRoute (n *Tree) [] string{
	if n == nil {
		return nil
	}
	var route []string

	// save route into a list
	for curr := n; curr != nil; curr = curr.Parent {
		linkName := scraper.StringToWikiUrl(curr.Value)
		route = append([]string{linkName}, route...)
	}

	return route
}