package main

import (
	"fmt"
	"runtime"
	"sync"
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

func bfs(start string, end string) []string {
	queue := Queue{}
	pages_visited := make(map[string]bool)
	parents := make(map[string]string)

	queue.Add(start)
	pages_visited[start] = true

	for !queue.IsEmpty() {
		// check channel antes de lançar go routine
		// add
		//comeca go routine
		current := queue.Pop()
		children := getWikiLinks(current)
		// TODO: only add after verification that was not visited already
		queue.Add(children...)

		for _, c := range children {
			if !pages_visited[c] {
				parents[c] = current
				pages_visited[c] = true

				if c == end {
					// mensagem channel a dizer para parar tudo
					return getPath(start, end, parents)
				}
			}
		}
	}
	// wait
	return nil
}

type Node struct {
	Value  string
	Parent *Node
}

func cbfs(start, end string, workers int) []string {

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
							result = GetParentsPath(childrenNode)
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

func GetParentsPath(node Node) []string {
	result := []string{node.Value}
	for node.Parent != nil {
		result = append([]string{node.Parent.Value}, result...)
		node = *node.Parent
	}
	// result = append(result, node.Parent.Value)
	return result
}

func bidirectional_bfs(start string, end string) []string {
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
					if testBidirectionalPath(possible_path, c) {
						return possible_path
					}
				}
			}
		}
		for _, c := range end_children {

			if !end_visited[c] {
				end_parents[c] = end_current
				end_visited[c] = true

				if start_visited[c] {
					possible_path := getBidirectionalPath(start, end, c, start_parents, end_parents)
					if testBidirectionalPath(possible_path, c) {
						return possible_path
					}
				}
			}
		}
	}
	return nil
}

// func exploreChidrenBidirectional(current Node, wg *sync.WaitGroup, myVisited, otherVisited *sync.Map, myChildren *CQueue) {

// 	defer func() {
// 		wg.Done()
// 	}()

// 	children := getWikiLinks(current.Value)
// 	for _, c := range children {
// 		_, exists := myVisited.Load(c)
// 		if !exists {
// 			childrenNode := Node{Value: c, Parent: &current}
// 			myChildren.Add(childrenNode)
// 			myVisited.Store(c, true)

// 			_, otherExists := otherVisited.Load(c)
// 			if otherExists {
// 				// TODO: add logic to get and test bi-path
// 				fmt.Println("EXISTS")
// 				possiblePath := getBidirectionalPath()
// 			}
// 		}
// 	}
// }

func c_bidirectional_bfs(start string, end string, workers int) []string {
	runtime.GOMAXPROCS(runtime.NumCPU())
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

func cGetPath(start string, end string, parents sync.Map) []string {
	path := []string{}
	current := end

	for current != start {
		path = append([]string{current}, path...)
		res, _ := parents.Load(current)
		current, _ = res.(string)
	}
	path = append([]string{start}, path...)
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

// Concurrency Queue Implementation
type CQueue struct {
	items []Node
	lock  sync.Mutex
}

func (q *CQueue) Add(item Node) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = append(q.items, item)
}

func (q *CQueue) AddAll(items ...Node) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = append(q.items, items...)
}

func (q *CQueue) Pop() Node {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.items) == 0 {
		return Node{}
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *CQueue) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.items) == 0
}

// Queue Nodes

type NQueue struct {
	items []Node
}

func (q *NQueue) Add(item Node) {
	q.items = append(q.items, item)
}

func (q *NQueue) AddAll(items ...Node) {
	q.items = append(q.items, items...)
}

func (q *NQueue) Pop() Node {
	if len(q.items) == 0 {
		return Node{}
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *NQueue) IsEmpty() bool {
	return len(q.items) == 0
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
