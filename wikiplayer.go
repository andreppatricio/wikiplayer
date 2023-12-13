package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	// var result []string
	// fmt.Println(result, result == nil, result != nil)
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("Usage: ./wikiplayer start end search [depth]")
		os.Exit(1)
	}

	// Access positional arguments
	start := args[0]
	end := args[1]
	type_of_search := args[2]

	if type_of_search == "dfs" && len(args) < 4 {
		fmt.Println("Search algorithm 'dfs' requires a maximum depth.")
		fmt.Println("Usage: ./wikiplayer start end search [depth]")
		os.Exit(1)
	}

	path := []string{}

	if type_of_search == "dfs" {
		max_depth, err := strconv.Atoi(args[3])
		if err != nil {
			panic(err)
		}
		path = dfs(start, end, max_depth)

	} else if type_of_search == "bfs" {
		path = bfs(start, end)

	} else if type_of_search == "bi_bfs" {
		path = bidirectional_bfs(start, end)
	} else if type_of_search == "cbfs" {
		path = cbfs(start, end, 10)
	} else if type_of_search == "cbi_bfs" {
		path = c_bidirectional_bfs(start, end, 10)
	} else {
		log.Fatalf("Type of search %s is not valid.", type_of_search)
	}

	if path == nil {
		fmt.Println("No path found")
	}
	for _, p := range path {
		fmt.Println(p)
	}

}
