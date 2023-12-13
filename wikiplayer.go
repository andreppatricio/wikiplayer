package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	workers := 1
	// fmt.Println(result, result == nil, result != nil)
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("Usage: ./wikiplayer start end search [depth] [workers]")
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

	if (type_of_search == "bfs" || type_of_search == "bdbfs") && len(args) >= 4 {
		var err error
		workers, err = strconv.Atoi(args[3])
		if err != nil {
			panic(err)
		}
	}

	path := []string{}

	if type_of_search == "dfs" {
		max_depth, err := strconv.Atoi(args[3])
		if err != nil {
			panic(err)
		}
		path = dfs(start, end, max_depth)

	} else if type_of_search == "bfs" {
		path = bfs(start, end, workers)

	} else if type_of_search == "bdbfs" {
		path = bidirectional_bfs(start, end, workers)

	} else {
		log.Fatalf("Type of search %s is not valid.", type_of_search)
	}

	// Print path
	if path == nil {
		fmt.Println("No path found")
	}
	for _, p := range path {
		fmt.Println(p)
	}

}
