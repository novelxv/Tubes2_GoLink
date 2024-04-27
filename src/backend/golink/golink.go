package golink

import (
	"fmt"
	"time"
	"reflect"
)

// fungsi untuk menyimpan statistik pencarian link Wiki
type GoLinkStats struct {
	Route              [][]string  // list rute pencarian
	LinksTraversed int         // Jumlah artikel yang telah dilalui
	LinksChecked   int         // Jumlah artikel yang telah diperiksa
	Runtime           time.Duration // Durasi runtime pencarian
}

// fungsi untuk menginisialisasi GoLinkStats dengan pohon root
func NewGoLinkStats() *GoLinkStats {
	return &GoLinkStats{
		Route:         [][]string{},
		LinksTraversed: 0,
		LinksChecked:   0,
		Runtime:        0,
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

// fungsi buat menambahkan rute
func (g *GoLinkStats) AddRoute(rute []string) {
	if(!SameList(g.Route,rute)){
		g.Route = append(g.Route, rute)
		g.LinksTraversed = len(rute)
	}
	
}

// fungsi untuk mengatur runtime pencarian
func (g *GoLinkStats) SetRuntime(duration time.Duration) {
	g.Runtime = duration
}

// fungsi untuk memeriksa apakah list sudah ada di dalam Route
func SameList(listOfLists [][]string, checkList []string) bool {
	// Iterasi melalui setiap list di Route
	for _, list := range listOfLists {
		// Memeriksa apakah list sudah ada di dalam Route
		if reflect.DeepEqual(list, checkList) {
			return true
		}
	}
	return false
}

// fungsi untuk mencetak statistik pencarian
func (g *GoLinkStats) PrintStats() {
	fmt.Printf("Total Links Traversed: %d\n", g.LinksTraversed)
	fmt.Printf("Total Links Checked: %d\n", g.LinksChecked)
	fmt.Printf("Total Runtime: %v\n", g.Runtime)
	fmt.Printf("Route: ")
	fmt.Println(g.Route)
}
