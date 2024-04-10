package logic

import (
	"fmt"

	"github.com/angiekierra/Tubes2_GoLink/utils"
)

type Tree struct {
	Value    string
	Children []*Tree
}

func NewNode(value string) *Tree {
	return &Tree{
		Value:    value,
		Children: []*Tree{},
	}
}

func (n *Tree) NewNodeLink(link []utils.Link) {
	for _, link := range link {
		temp := NewNode(link.Name)
		n.AddChild(temp)
	}
}

func (n *Tree) AddChild(child *Tree) {
	n.Children = append(n.Children, child)
}

// DepthFirstTraversal performs a depth-first traversal of the tree starting from this node
func (n *Tree) PrintTreeIds() {
	fmt.Printf("%s ", n.Value)
	fmt.Printf("(")
	for _, child := range n.Children {
		child.PrintTreeIds()
	}
	fmt.Printf(")")
}
