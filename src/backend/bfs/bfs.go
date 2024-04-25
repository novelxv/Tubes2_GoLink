package bfs

import (	
	"fmt"
	"time"
	// "context"
	"sync"
	"github.com/angiekierra/Tubes2_GoLink/scraper"
	"github.com/angiekierra/Tubes2_GoLink/tree"
	"github.com/angiekierra/Tubes2_GoLink/golink"
)


// main BFS function
func Bfsfunc(value string, goal string) *golink.GoLinkStats {
	startTime := time.Now()

	// save the root
	root := tree.NewNode(value)
	stats := golink.NewGoLinkStats()

	// use BFS to search for the goal
	// found := SearhForGoalBfs(root, goal, stats)
	found := SearchForGoalBfsMT(root, goal, stats)

	elapsedTime := time.Since(startTime)
	stats.SetRuntime(elapsedTime)

	if found {
		PrintTreeBfs(root)
		stats.PrintStats()
	}

	return stats
}

// function to print the tree using BFS
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

// function to search the word goal recursively in BFS
func SearhForGoalBfs(n *tree.Tree, goal string, stats *golink.GoLinkStats) bool {
	queue := []*tree.Tree{n}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if !current.Visited {
			current.Visited = true
			stats.AddTraversed()
		}

		fmt.Printf("%s \n", current.Value)
		stats.AddChecked()
		
		if tree.IsGoalFound(current.Value, goal) {
			fmt.Print("Found!!\n")
			route := tree.GoalRoute(current)
			stats.AddRoute(route)
			return true
		}
		
		linkName := scraper.StringToWikiUrl(current.Value)
		links, _ := scraper.Scraper(linkName)
		current.NewNodeLink(links)

		for _, child := range current.Children {
			if !child.Visited {
				queue = append(queue, child)
			}
		}
	}
	return false
}

// searchForGoalBfsMT handles the BFS using goroutines
func SearchForGoalBfsMT(root *tree.Tree, goal string, stats *golink.GoLinkStats) bool {
	var wg sync.WaitGroup
	nodeQueue := make(chan *tree.Tree)
	results := make(chan bool)
	done := make(chan struct{})

	go func() {
		nodeQueue <- root
	}()

	wg.Add(1)
	go func() {
		defer close(nodeQueue)
		for node := range nodeQueue {
			if node.Visited {
				wg.Done()
				continue
			}

			node.Visited = true
			stats.AddTraversed()

			if tree.IsGoalFound(node.Value, goal) {
				fmt.Println("Found!!")
				route := tree.GoalRoute(node)
				stats.AddRoute(route)
				results <- true
				close(done)
				return
			}

			links, _ := scraper.Scraper(scraper.StringToWikiUrl(node.Value))
			for _, link := range links {
				child := tree.NewNode(link.Name)
				node.AddChild(child)
				if !child.Visited {
					wg.Add(1)
					go func(ch *tree.Tree) {
						nodeQueue <- ch
						wg.Done()
					}(child)
				}
			}
			wg.Done()
		}
	}()

	go func() {
		wg.Wait()
		results <- false
	}()

	select {
	case found := <-results:
		return found
	case <-done:
		return true
	}
}

// func SearchForGoalBfsMT(root *tree.Tree, goal string, stats *golink.GoLinkStats) bool {
// 	start := time.Now() // delete
//     ctx, cancel := context.WithCancel(context.Background())
//     defer cancel() // Ensure cancellation signal is sent when function exits

//     queue := make(chan *tree.Tree, 100)
//     done := make(chan bool)
//     var wg sync.WaitGroup

//     queue <- root
//     wg.Add(1)

//     for i := 0; i < 20; i++ { // Start 20 workers
//         go func() {
//             defer wg.Done()
//             for {
//                 select {
//                 case n, ok := <-queue:
//                     if !ok {
//                         return // Channel is closed, exit goroutine
//                     }
//                     if n.Visited {
//                         continue
//                     }

//                     n.Visited = true
//                     stats.AddTraversed()

//                     if tree.IsGoalFound(n.Value, goal) {
// 						elapsed := time.Since(start) // delete
//                         fmt.Print("Found!!\n")
// 						fmt.Printf("Runtime: %v\n", elapsed) // delete
//                         route := tree.GoalRoute(n)
//                         stats.AddRoute(route)
//                         stats.AddChecked()
//                         stats.PrintStats()
//                         cancel() // Use cancel function to signal all goroutines to stop
//                         done <- true
//                         return
//                     }

//                     linkName := scraper.StringToWikiUrl(n.Value)
//                     links, err := scraper.Scraper2(linkName)
//                     if err != nil {
//                         fmt.Printf("Error scraping %s: %v\n", linkName, err)
//                         continue
//                     }
//                     n.NewNodeLink(links)

//                     for _, child := range n.Children {
//                         if !child.Visited {
//                             wg.Add(1)
//                             go func(ch *tree.Tree) {
//                                 select {
//                                 case queue <- ch:
//                                 case <-ctx.Done():
//                                     return // Stop sending if the context is cancelled
//                                 }
//                             }(child)
//                         }
//                     }
//                 case <-ctx.Done():
//                     return // Exit goroutine when context is cancelled
//                 }
//             }
//         }()
//     }

//     wg.Wait() // Wait for all goroutines to complete
//     close(queue) // Close the channel after all goroutines are done
//     select {
//     case <-done:
//         return true
//     default:
//         fmt.Println("Search cancelled or goal found.")
//         return false
//     }
// }


// func SearchForGoalBfsMT(root *tree.Tree, goal string, stats *golink.GoLinkStats) bool {
//     ctx, cancel := context.WithCancel(context.Background())
//     defer cancel() // Ensure cancellation signal is sent when function exits

//     queue := make(chan *tree.Tree, 100) // Use a buffered channel
//     done := make(chan bool)
//     var wg sync.WaitGroup

//     // Signal to close channels and cleanup
//     defer close(done)
//     defer func() {
//         wg.Wait() // Ensure all processing is done before closing the queue
//         close(queue)
//     }()

//     queue <- root
//     wg.Add(1)

//     for i := 0; i < 20; i++ { // Start 20 workers
//         go func() {
//             defer wg.Done()
//             for {
//                 select {
//                 case n := <-queue:
//                     if n == nil {
//                         return // Exit on nil signal, means channel is closing
//                     }

//                     if n.Visited {
//                         continue
//                     }

//                     n.Visited = true
//                     stats.AddTraversed()

//                     if tree.IsGoalFound(n.Value, goal) {
//                         fmt.Print("Found!!\n")
//                         route := tree.GoalRoute(n)
//                         stats.AddRoute(route)
// 						stats.AddChecked()
// 						stats.PrintStats()
// 						scraper.CancelScraping() // Signal scraper to stop
//                         cancel() // Use cancel function to signal all goroutines to stop
//                         done <- true
//                         return
//                     }

//                     linkName := scraper.StringToWikiUrl(n.Value)
//                     links, err := scraper.Scraper2(linkName)
//                     if err != nil {
//                         fmt.Printf("Error scraping %s: %v\n", linkName, err)
//                         continue
//                     }
//                     n.NewNodeLink(links)

//                     for _, child := range n.Children {
//                         if !child.Visited {
//                             wg.Add(1)
//                             go func(ch *tree.Tree) {
//                                 select {
//                                 case queue <- ch:
//                                 case <-ctx.Done():
//                                     return // Stop sending if the context is cancelled
//                                 }
//                             }(child)
//                         }
//                     }
//                 case <-ctx.Done():
//                     return // Exit goroutine when context is cancelled
//                 }
//             }
//         }()
//     }

//     // Wait for the "done" signal or all work to complete
//     select {
//     case <-done:
//         return true
//     case <-ctx.Done():
//         fmt.Println("Search cancelled or goal found.")
//         return false
//     }
// }



// func SearhForGoalBfsMT(n *tree.Tree, goal string, stats *golink.GoLinkStats) bool {
//     queue := make(chan *tree.Tree, 1)
//     queue <- n
//     found := make(chan bool)
//     sem := make(chan struct{}, 20) // Limit to 20 concurrent goroutines

//     go func() {
//         var wg sync.WaitGroup
//         for node := range queue {
//             wg.Add(1)
//             sem <- struct{}{} // Acquire a token
//             go func(n *tree.Tree) {
//                 defer wg.Done()
//                 defer func() { <-sem }() // Release the token

//                 if !n.Visited {
//                     n.Visited = true
//                     stats.AddTraversed()
//                 }

//                 fmt.Printf("%s \n", n.Value)

//                 if tree.IsGoalFound(n.Value, goal) {
//                     fmt.Print("Found!!\n")
//                     route := tree.GoalRoute(n)
//                     stats.AddRoute(route)
// 					stats.AddChecked()
//                     found <- true
//                     return
//                 }

//                 linkName := scraper.StringToWikiUrl(n.Value)
//                 links, _ := scraper.Scraper(linkName)
//                 n.NewNodeLink(links)

//                 for _, child := range n.Children {
//                     if !child.Visited {
//                         queue <- child
//                     }
//                 }
//             }(node)
//         }
//         wg.Wait()
//         close(found)
//     }()

//     return <-found
// }

// func searchForGoalBfsMT(root *tree.Tree, goal string, stats *golink.GoLinkStats) bool {
//     var wg sync.WaitGroup
//     currentLayer := []*tree.Tree{root}
//     nextLayer := []*tree.Tree{}
//     found := false

//     // Create a buffered channel as a semaphore
//     sem := make(chan struct{}, 20) // Limit to 20 concurrent goroutines

//     for len(currentLayer) > 0 && !found {
//         results := make(chan bool, len(currentLayer))

//         for _, node := range currentLayer {
//             wg.Add(1)
//             go func(n *tree.Tree) {
//                 sem <- struct{}{} // Acquire a token
//                 defer wg.Done()
//                 defer func() { <-sem }() // Release the token

//                 if tree.IsGoalFound(n.Value, goal) {
//                     results <- true
//                     return
//                 }

//                 linkName := scraper.StringToWikiUrl(n.Value)
//                 links, _ := scraper.Scraper(linkName)
//                 n.NewNodeLink(links)

//                 for _, link := range links {
//                     child := tree.NewNode(link.Name)
//                     n.AddChild(child)
//                     nextLayer = append(nextLayer, child)
//                 }
//                 results <- false
//             }(node)
//         }

//         // Wait for all goroutines in this layer to finish
//         wg.Wait()
//         close(results) // Closing results to stop range operation

//         // Check results for any found signal
//         for result := range results {
//             if result {
//                 found = true
//                 break
//             }
//         }

//         // Prepare the next layer to be the current layer in the next iteration
//         currentLayer = nextLayer
//         nextLayer = []*tree.Tree{}
//     }

//     return found
// }

// func searchForGoalBfsMT(root *tree.Tree, goal string, stats *golink.GoLinkStats) bool {
//     var wg sync.WaitGroup
//     currentLayer := []*tree.Tree{root}
//     nextLayer := []*tree.Tree{}
//     found := false

//     for len(currentLayer) > 0 && !found {
//         results := make(chan bool, len(currentLayer))

//         for _, node := range currentLayer {
//             wg.Add(1)
//             go func(n *tree.Tree) {
//                 defer wg.Done()
//                 if tree.IsGoalFound(n.Value, goal) {
//                     results <- true
//                     return
//                 }

//                 linkName := scraper.StringToWikiUrl(n.Value)
//                 links, _ := scraper.Scraper(linkName)
//                 n.NewNodeLink(links)

//                 for _, link := range links {
//                     child := tree.NewNode(link.Name)
//                     n.AddChild(child)
//                     nextLayer = append(nextLayer, child)
//                 }
//                 results <- false
//             }(node)
//         }

//         // Wait for all goroutines in this layer to finish
//         wg.Wait()
//         close(results) // Closing results to stop range operation

//         // Check results for any found signal
//         for result := range results {
//             if result {
//                 found = true
//                 break
//             }
//         }

//         // Prepare the next layer to be the current layer in the next iteration
//         currentLayer = nextLayer
//         nextLayer = []*tree.Tree{}
//     }

//     return found
// }

// func SearchForGoalBfsMT(root *tree.Tree, goal string, stats *golink.GoLinkStats) bool {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	queue := make(chan *tree.Tree, 100)
// 	var wg sync.WaitGroup
// 	found := false

// 	// Start worker goroutines
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for {
// 				select {
// 				case node, ok := <-queue:
// 					if !ok {
// 						return // Channel closed, exit the goroutine
// 					}
// 					if tree.IsGoalFound(node.Value, goal) {
// 						found = true
// 						cancel() // Cancel all goroutines
// 						return
// 					}

// 					links, err := scraper.Scraper(scraper.StringToWikiUrl(node.Value))
// 					if err != nil {
// 						fmt.Println("Error scraping:", err)
// 						continue
// 					}
// 					for _, link := range links {
// 						select {
// 						case queue <- tree.NewNode(link.Name):
// 						case <-ctx.Done():
// 							return // Exit on cancellation
// 						}
// 					}
// 				case <-ctx.Done():
// 					return // Exit on cancellation
// 				}
// 			}
// 		}()
// 	}

// 	// Send the root to the queue and start processing
// 	queue <- root

// 	wg.Wait() // Wait for all goroutines to complete
// 	close(queue) // Close the channel after all goroutines are done

// 	return found
// }