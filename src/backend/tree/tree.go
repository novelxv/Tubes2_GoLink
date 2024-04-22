package tree

import (
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