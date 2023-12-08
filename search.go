package main

import (
	"fmt"
	"time"
)

//

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func printWiki(wiki []string) {
	s := make([]string, len(wiki))
	copy(s, wiki)
	for i := range s {
		s[i] = s[i][30:]
	}
	fmt.Println(s)
}

func dfs(start string, end string, max_depth int) []string {

	path := []string{start}
	pages_visited := make(map[string]int)

	var recursiveHelper func(p []string)
	recursiveHelper = func(p []string) {

		last := p[len(p)-1]

		pages_visited[last] = len(p)

		if last == end {
			fmt.Println("FOUND PATH")
			printWiki(p)
		}

		if len(p) >= max_depth {
			return
		}

		children := getWikiLinks(last)
		for _, c := range children {
			value, exists := pages_visited[c]
			if !exists || len(p)+1 < value {
				p = append(p, c)
				recursiveHelper(p)
				p = p[:len(p)-1]
			}
		}
	}
	recursiveHelper(path)
	return path
}

func bfs(start string, end string) []string {
	// durations := []float64{}
	queue := Queue{}
	pages_visited := make(map[string]bool)
	parents := make(map[string]string)

	queue.Add(start)
	pages_visited[start] = true

	counter := 0
	start_time := time.Now()
	for !queue.IsEmpty() {
		counter += 1
		if counter%100 == 0 {
			duration := time.Since(start_time).Seconds() / float64(counter)
			fmt.Println("Page ", counter, " - Avg time per iteration: ", duration)
		}
		current := queue.Pop()
		children := getWikiLinks(current)
		// startt := time.Now()
		queue.Add(children...)
		// durations = append(durations, float64(time.Since(startt).Seconds()))

		for _, c := range children {
			if !pages_visited[c] {
				parents[c] = current
				pages_visited[c] = true

				if c == end {
					return getPath(start, end, parents)
				}
			}
		}
	}
	return nil

}

func bidirectional_bfs(start string, end string) []string {
	// durations := []float64{}
	start_queue := Queue{}
	end_queue := Queue{}
	start_visited := make(map[string]bool)
	end_visited := make(map[string]bool)
	start_parents := make(map[string]string)
	end_parents := make(map[string]string)

	start_queue.Add(start)
	end_queue.Add(end)

	start_visited[start] = true
	end_visited[end] = true

	for !start_queue.IsEmpty() && !end_queue.IsEmpty() {

		start_current := start_queue.Pop()
		end_current := end_queue.Pop()

		start_children := getWikiLinks(start_current)
		start_queue.Add(start_children...)
		end_children := getWikiLinks(end_current)
		end_queue.Add(end_children...)

		for _, c := range start_children {

			if !start_visited[c] {
				start_parents[c] = start_current
				start_visited[c] = true

				if end_visited[c] {
					possible_path := getBidirectionalPath(start, end, c, start_parents, end_parents)
					// fmt.Println("Found possible connection (start path):", c)
					// printWiki(possible_path)
					if testBidirectionalPath(possible_path, c) {
						fmt.Println("It is a real path!")
						return possible_path
					}
					// else {
					// 	fmt.Println("It is not a real path")
					// 	// fmt.Scanln()
					// }
				}
			}
		}
		for _, c := range end_children {

			if !end_visited[c] {
				end_parents[c] = end_current
				end_visited[c] = true

				if start_visited[c] {
					possible_path := getBidirectionalPath(start, end, c, start_parents, end_parents)
					// fmt.Println("Found possible connection (end path):", c)
					// printWiki(possible_path)
					if testBidirectionalPath(possible_path, c) {
						// fmt.Println("It is a real path!")
						return possible_path
					}
					// else {
					// 	fmt.Println("It is not a real path")
					// 	// fmt.Scanln()
					// }
				}
			}
		}
	}
	return nil
}

func testBidirectionalPath(path []string, connection string) bool {

	var connection_id int = -1
	for ind, node := range path {
		if node == connection {
			connection_id = ind
		}
	}

	test_path := path[connection_id:]
	current_ind := 0
	end := test_path[len(test_path)-1]
	for test_path[current_ind] != end {
		current := test_path[current_ind]
		next := test_path[current_ind+1]
		children := getWikiLinks(current)
		if !contains(children, next) {
			return false
		}
		current_ind += 1
	}
	return true
}

func getBidirectionalPath(start string, end string, connection string, start_parents map[string]string, end_parents map[string]string) []string {

	start_path := []string{}
	end_path := []string{}

	current := connection
	for current != start {
		start_path = append([]string{current}, start_path...)
		current = start_parents[current]
	}
	start_path = append([]string{start}, start_path...)

	current = connection
	for current != end {
		end_path = append(end_path, current)
		current = end_parents[current]
	}
	end_path = append(end_path, end)
	// printWiki(start_path)
	// printWiki(end_path)
	path := append(start_path, end_path[1:]...)
	return path

}

func getPath(start string, end string, parents map[string]string) []string {
	path := []string{}
	current := end

	for current != start {
		path = append([]string{current}, path...)
		current = parents[current]
	}
	path = append([]string{start}, path...)
	return path
}

func getParentDepth(current string, parents map[string]string) int {
	counter := 0
	exists := true
	for exists {
		counter = counter + 1
		current, exists = parents[current]
	}
	return counter
}

// for !stack.IsEmpty() {
// 	// fmt.Scanln()
// 	visit, _ := stack.Pop()
// 	// fmt.Println("Visit: ", visit)
// 	// fmt.Println("Path: ", path)
// 	// printWiki(path)
// 	// fmt.Println("Stack: ")
// 	// printWiki(stack.items)
// 	if visit == "https://en.wikipedia.org/wiki/University_of_Toronto" {
// 		fmt.Println("found Toronto")
// 		fmt.Println(path)
// 		fmt.Scanln()
// 	}
// 	if contains(path, visit) {
// 		// fmt.Println("Found loop")
// 		// fmt.Println(path)
// 		// fmt.Println(visit)
// 		continue
// 	}
// 	path = append(path, visit)

// 	// if len(path) >= 2 && path[1] == "https://en.wikipedia.org/wiki/University_of_Toronto" {
// 	// 	fmt.Println("found Toronto")
// 	// 	fmt.Println(path)
// 	// 	fmt.Scanln()
// 	// }

// 	if visit == end {
// 		break
// 	}
// 	if len(path) <= max_depth {
// 		// fmt.Println("Getting new links from ", visit)
// 		links := getWikiLinks(visit, section)
// 		stack.Push(links...)
// 	} else {
// 		// fmt.Println("Path at max size")
// 		index := len(path) - 1
// 		path = path[:index]
// 	}
// 	// time.Sleep(2 * time.Second)
// }
// if end == path[len(path)-1] {
// 	return path
// } else {
// 	fmt.Println("No path found with maximum depth of ", max_depth)
// 	return []string{}
// }
// }
