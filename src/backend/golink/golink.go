package golink

import (
	"fmt"
	"github.com/angiekierra/Tubes2_GoLink/tree"
	"time"
)

// fungsi untuk menyimpan statistik pencarian link Wiki
type GoLinkStats struct {
	Tree              *tree.Tree  // Pohon pencarian
	ArticlesTraversed int         // Jumlah artikel yang telah dilalui
	ArticlesChecked   int         // Jumlah artikel yang telah diperiksa
	Runtime           time.Duration // Durasi runtime pencarian
}

// fungsi untuk menginisialisasi GoLinkStats dengan pohon root
func NewGoLinkStats(root *tree.Tree) *GoLinkStats {
	return &GoLinkStats {
		Tree:              root,
		ArticlesTraversed: 0,
		ArticlesChecked:   0,
		Runtime:           0,
	}
}

// fungsi untuk menambahkan jumlah artikel yang telah dilalui
func (g *GoLinkStats) AddTraversed() {
	g.ArticlesTraversed++
}

// fungsi untuk menambahkan jumlah artikel yang telah diperiksa
func (g *GoLinkStats) AddChecked() {
	g.ArticlesChecked++
}

// fungsi untuk mengatur runtime pencarian
func (g *GoLinkStats) SetRuntime(duration time.Duration) {
	g.Runtime = duration
}

// fungsi untuk mencetak statistik pencarian
func (g *GoLinkStats) PrintStats() {
	fmt.Printf("Total Articles Traversed: %d\n", g.ArticlesTraversed)
	fmt.Printf("Total Articles Checked: %d\n", g.ArticlesChecked)
	fmt.Printf("Total Runtime: %v\n", g.Runtime)
}
