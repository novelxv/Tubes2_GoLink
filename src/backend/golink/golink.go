package golink

import (
	"fmt"
	"github.com/angiekierra/Tubes2_GoLink/tree"
	"time"
)

// fungsi untuk menyimpan statistik pencarian link Wiki
type GoLinkStats struct {
	Tree              *tree.Tree  // Pohon pencarian
	LinksTraversed int         // Jumlah artikel yang telah dilalui
	LinksChecked   int         // Jumlah artikel yang telah diperiksa
	Runtime           time.Duration // Durasi runtime pencarian
}

// fungsi untuk menginisialisasi GoLinkStats dengan pohon root
func NewGoLinkStats(root *tree.Tree) *GoLinkStats {
	return &GoLinkStats {
		Tree:              root,
		LinksTraversed: 0,
		LinksChecked:   0,
		Runtime:           0,
	}
}

// fungsi untuk menambahkan jumlah artikel yang telah dilalui
func (g *GoLinkStats) AddTraversed() {
	g.LinksTraversed++
}

// fungsi untuk menambahkan jumlah artikel yang telah diperiksa
func (g *GoLinkStats) AddChecked() {
	g.LinksChecked++
}

// fungsi untuk mengatur runtime pencarian
func (g *GoLinkStats) SetRuntime(duration time.Duration) {
	g.Runtime = duration
}

// fungsi untuk mencetak statistik pencarian
func (g *GoLinkStats) PrintStats() {
	fmt.Printf("Total Links Traversed: %d\n", g.LinksTraversed)
	fmt.Printf("Total Links Checked: %d\n", g.LinksChecked)
	fmt.Printf("Total Runtime: %v\n", g.Runtime)
}
