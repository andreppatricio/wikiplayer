package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]
	start := args[0]
	end := args[1]
	type_of_search := args[2]

	path := []string{}

	if type_of_search == "dfs" {
		max_depth, _ := strconv.Atoi(args[3])
		path = dfs(start, end, max_depth)

	} else if type_of_search == "bfs" {
		path = bfs(start, end)

	} else if type_of_search == "bi_bfs" {
		// path = getWikiLinks(start)
		path = bidirectional_bfs(start, end)
	} else {
		log.Fatalf("Type of search %s is not valid.", type_of_search)
	}

	for _, p := range path {
		fmt.Println(p)
	}

}
