package main

import (
	"sync"
)

// Structs

// Helper Functions

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getParentsPath(node Node) []string {
	result := []string{node.Value}
	for node.Parent != nil {
		result = append([]string{node.Parent.Value}, result...)
		node = *node.Parent
	}
	return result
}

// Search Algorithms

func dfs(start string, end string, max_depth int) []string {

	path := []string{start}
	pages_visited := make(map[string]int)

	var recursiveHelper func(p []string) ([]string, bool)
	recursiveHelper = func(p []string) ([]string, bool) {
		last := p[len(p)-1]

		pages_visited[last] = len(p)

		if last == end {
			return p, true
		}

		if len(p) >= max_depth {
			return nil, false
		}

		children := getWikiLinks(last)
		for _, c := range children {

			value, exists := pages_visited[c]
			if !exists || len(p)+1 < value {
				p = append(p, c)
				res, found := recursiveHelper(p)
				if found {
					return res, found
				}
				p = p[:len(p)-1]
			}
		}
		return nil, false
	}

	result, found := recursiveHelper(path)
	if found {
		return result
	} else {
		return nil
	}
}

func bfs(start, end string, workers int) []string {

	// runtime.GOMAXPROCS(runtime.NumCPU())
	// runtime.GOMAXPROCS(1)

	queue := CQueue{}
	childrenQueue := CQueue{}

	var pagesVisited sync.Map
	var wg sync.WaitGroup
	var result []string

	ch := make(chan struct{}, workers)

	startNode := Node{Value: start, Parent: nil}
	queue.Add(startNode)
	pagesVisited.Store(start, true)

	for {
		for !queue.IsEmpty() && result == nil {
			current := queue.Pop()

			ch <- struct{}{}
			wg.Add(1)
			go func(current Node) {
				defer func() {
					<-ch
					wg.Done()
				}()

				children := getWikiLinks(current.Value)

				for _, c := range children {
					_, exists := pagesVisited.Load(c)
					if !exists {
						childrenNode := Node{Value: c, Parent: &current}
						childrenQueue.Add(childrenNode)
						pagesVisited.Store(c, true)

						if c == end && result == nil {
							result = getParentsPath(childrenNode)
						}
					}
				}
			}(current)

		}
		wg.Wait()

		if result == nil {
			queue = CQueue{items: childrenQueue.items}
			childrenQueue = CQueue{}
		} else {
			break
		}

	}
	return result
}

func bidirectional_bfs(start string, end string, workers int) []string {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	// runtime.GOMAXPROCS(1)

	startQueue := CQueue{}
	endQueue := CQueue{}
	startChildren := CQueue{}
	endChildren := CQueue{}
	var wg sync.WaitGroup
	var startVisited sync.Map
	var endVisited sync.Map
	var result []string

	ch := make(chan struct{}, workers)

	// start_parents := make(map[string]string)
	// end_parents := make(map[string]string)
	startNode := Node{Value: start, Parent: nil}
	endNode := Node{Value: end, Parent: nil}
	startQueue.Add(startNode)
	endQueue.Add(endNode)

	startVisited.Store(start, startNode)
	endVisited.Store(end, endNode)

	for {
		for !startQueue.IsEmpty() && !endQueue.IsEmpty() && result == nil {

			if !startQueue.IsEmpty() {
				startCurrent := startQueue.Pop()

				ch <- struct{}{}
				wg.Add(1)
				go func(current Node) {
					defer func() {
						<-ch
						wg.Done()
					}()

					children := getWikiLinks(current.Value)
					for _, c := range children {
						_, exists := startVisited.Load(c)
						if !exists {
							childrenNode := Node{Value: c, Parent: &current}
							startChildren.Add(childrenNode)
							startVisited.Store(c, childrenNode)

							endNode, endExists := endVisited.Load(c)
							if endExists {
								// fmt.Println("EXISTS start", c)
								endNode := endNode.(Node)
								// fmt.Println(childrenNode, endNode)
								possible_path := getNodeBidirectionalPath(childrenNode, endNode)
								// fmt.Println(possible_path)
								// fmt.Scanln()
								if testBidirectionalPath(possible_path, c) && result == nil {
									result = possible_path
								}

							}
						}
					}
				}(startCurrent)
			}

			if !endQueue.IsEmpty() {
				endCurrent := endQueue.Pop()

				ch <- struct{}{}
				wg.Add(1)
				go func(current Node) {
					defer func() {
						<-ch
						wg.Done()
					}()

					children := getWikiLinks(current.Value)
					for _, c := range children {
						_, exists := endVisited.Load(c)
						if !exists {
							childrenNode := Node{Value: c, Parent: &current}
							endChildren.Add(childrenNode)
							endVisited.Store(c, childrenNode)

							startNode, startExists := startVisited.Load(c)
							if startExists {
								// fmt.Println("EXISTS end", c)
								startNode := startNode.(Node)
								possible_path := getNodeBidirectionalPath(startNode, childrenNode)
								// fmt.Println(possible_path)
								// fmt.Scanln()
								if testBidirectionalPath(possible_path, c) && result == nil {
									result = possible_path
								}
							}
						}
					}
				}(endCurrent)
			}

		}
		wg.Wait()

		if result == nil {
			startQueue = CQueue{items: startChildren.items}
			startChildren = CQueue{}
			endQueue = CQueue{items: endChildren.items}
			endChildren = CQueue{}
		} else {
			break
		}
	}
	return result
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

func getNodeBidirectionalPath(startNode, endNode Node) []string {
	startPath := []string{startNode.Value}
	endPath := []string{}

	for startNode.Parent != nil {
		startPath = append([]string{startNode.Parent.Value}, startPath...)
		startNode = *startNode.Parent
	}
	for endNode.Parent != nil {
		endPath = append(endPath, endNode.Parent.Value)
		endNode = *endNode.Parent
	}
	path := append(startPath, endPath...)
	return path
}
