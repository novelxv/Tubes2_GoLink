package bfs

import (	
	"fmt"
	"time"
	"sync"
	"context"
	"github.com/angiekierra/Tubes2_GoLink/scraper"
	"github.com/angiekierra/Tubes2_GoLink/tree"
	"github.com/angiekierra/Tubes2_GoLink/golink"
)

/* Main BFS Function */
func Bfsfunc(value string, goal string, multisol bool) *golink.GoLinkStats {
	// mulai track waktu
	startTime := time.Now()
	// buat root dan stats	
	root := tree.NewNode(value)
	stats := golink.NewGoLinkStats()
	// cari goal
	found := SearchForGoalBfs(root, goal, stats, multisol)
	// set runtime
	elapsedTime := time.Since(startTime)
	stats.SetRuntime(elapsedTime)
	// print stats
	if found {
		stats.PrintStats()
	} else {
		fmt.Println("Goal not found")
	}
	return stats
}

/* Print Tree */
func PrintTreeBfs(n *tree.Tree) {
	queue := []*tree.Tree{n}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.Visited {
			fmt.Printf("|%s", current.Value)
			queue = append(queue, current.Children...)
		}
	}
	fmt.Println()
}

// mutex untuk lock goroutine
var mu sync.Mutex

/* Cari Goal */
func SearchForGoalBfs(root *tree.Tree, goal string, stats *golink.GoLinkStats, multisol bool) bool {
	// buat context, cancel jika sudah 5 menit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	// buat waitgroup
	var wg sync.WaitGroup
	nodeQueue := make(chan *tree.Tree, 10000000)
	
	found := false
	var solutionDepth int
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		nodeQueue <- root
	}()

	go func() {
		// pastikan channel ditutup ketika selesai
		defer close(nodeQueue) 
		for node := range nodeQueue {
			if node.Visited {
				continue
			}
			// set node jadi visited
			node.Visited = true
			stats.AddChecked()

			// cek apakah goal sudah ditemukan
			if tree.IsGoalFound(node.Value, goal) {
				fmt.Println("Found!!")
				route := tree.GoalRoute(node)
				stats.AddRoute(route)
				found = true
				solutionDepth = len(route) - 1
				// jika tidak multi solution, cancel semua operasi
				if !multisol {
					cancel() // Notify to cancel all operations
					return
				}
			}

			// jika depth lebih dari solution depth, berhenti cari
			if solutionDepth != 0 && node.GetDepth() > solutionDepth {
				cancel()
				return
			}

			// jika belum ditemukan, lakukan scrapping
			mu.Lock()
			links, _ := scraper.Scraper(scraper.StringToWikiUrl(node.Value))

			// batasi jumlah goroutine
			const maxGoroutines = 5000
			sem := make(chan bool, maxGoroutines)

			visitedNodes := make(map[string]struct{})

			for _, link := range links {
				// jika link sudah pernah dikunjungi, skip
				if _, ok := visitedNodes[link.Name]; ok {
					continue
				}

				// buat node baru untuk tiap child
				child := tree.NewNode(link.Name)
				node.AddChild(child)
				visitedNodes[link.Name] = struct{}{}

				// jika belum pernah dikunjungi dan depth kurang dari 6, tambahkan ke queue
				if !child.Visited && child.GetDepth() <= 6 {
					wg.Add(1)
					go func(ch *tree.Tree) {
						sem <- true
						defer func() {
							<-sem
							wg.Done()
						}()
						select {
						case nodeQueue <- ch:
						case <-ctx.Done(): // Handle cancellation
						}
					}(child)
				}
			}
			mu.Unlock()
		}
	}()

	wg.Wait() 
	<-ctx.Done() 

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Search timed out after 5 minutes")
	}
	return found
}