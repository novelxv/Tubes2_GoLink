package tree

import (
	"github.com/angiekierra/Tubes2_GoLink/scraper"
)

/* Tree Struct */
type Tree struct {
	Value    string // the link name
	Visited  bool   // check whether a node has already been visited or not
	MainRoute  bool  // main node to the goal
	Parent 	*Tree // save parent
	Children []*Tree // save child
	Depth 	int // save depth
}

/* Buat Node Baru */
func NewNode(value string) *Tree {
	return &Tree{
		Value:    value,
		Children: []*Tree{},
		Depth: 0,
	}
}

/* Buat Nodes untuk kumpulan link */ 
func (n *Tree) NewNodeLink(link []scraper.Link) {
	for _, link := range link {
		temp := NewNode(link.Name)
		n.AddChild(temp)
	}
}

/* Menambah Child dari Node Parent */
func (n *Tree) AddChild(child *Tree) {
	child.Parent = n
	child.Depth = n.Depth + 1
	n.Children = append(n.Children, child)
}

/* Get Depth */
func (n *Tree) GetDepth() int {
	return n.Depth
}

/* Ubah Node jadi Visited */
func (n *Tree) AddVisitedNode() {
	n.Visited = true
}

/* Menambahkan Node ke Main Route */
func (n *Tree) AddMainRoute() {
	n.MainRoute = true
}

/* Remove Node dari Main Route */
func (n *Tree) UndoMainRoute() {
	n.MainRoute = false
}

/* Cek apakah Node adalah Goal */
func IsGoalFound(s1 string, s2 string) bool {
	return s1 == s2
}

/* Mendapatkan Route ke Goal */
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

/* Mendapatkan Panjang Parent ke Child  */
func (n *Tree) ParentLength() int {
	length := 0
	for curr := n.Parent; curr != nil; curr = curr.Parent {
		length++
	}
	return length
}