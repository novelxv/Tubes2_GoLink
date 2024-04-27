package golink

import (
	"fmt"
	"time"
	"reflect"
)

/* Struct GoLinkStats */
type GoLinkStats struct {
	Route              [][]string  // list rute pencarian
	LinksTraversed int         // jumlah artikel yang telah dilalui
	LinksChecked   int         // jumlah artikel yang telah diperiksa
	Runtime           time.Duration // durasi runtime pencarian
}

/* Inisialisasi Stats */
func NewGoLinkStats() *GoLinkStats {
	return &GoLinkStats{
		Route:         [][]string{},
		LinksTraversed: 0,
		LinksChecked:   0,
		Runtime:        0,
	}
}

/* Menambahkan Jumlah Artikel yang Telah Dilalui */
func (g *GoLinkStats) AddTraversed() {
	g.LinksTraversed++
}

/* Menambahkan Jumlah Artikel yang Telah Diperiksa */
func (g *GoLinkStats) AddChecked() {
	g.LinksChecked++
}

/* Menambahkan Rute */
func (g *GoLinkStats) AddRoute(rute []string) {
	if(!SameList(g.Route,rute)){
		g.Route = append(g.Route, rute)
		g.LinksTraversed = len(rute)
	}
	
}

/* Set Durasi Runtime */
func (g *GoLinkStats) SetRuntime(duration time.Duration) {
	g.Runtime = duration
}

/* Memeriksa Apakah List Sudah Ada di dalam Route */
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

/* Print Stats */
func (g *GoLinkStats) PrintStats() {
	fmt.Printf("Total Links Traversed: %d\n", g.LinksTraversed)
	fmt.Printf("Total Links Checked: %d\n", g.LinksChecked)
	fmt.Printf("Total Runtime: %v\n", g.Runtime)
	fmt.Printf("Route: ")
	fmt.Println(g.Route)
}
