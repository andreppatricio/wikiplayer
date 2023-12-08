# The Wikipedia Game Player

Welcome to the **The Wikipedia Game Player** project!
This Go(lang) project gives solutions for [The Wikipedia Game](https://en.wikipedia.org/wiki/Wikipedia:Wiki_Game) through web scrapping and various search algorithms.

The Wikipedia Game is a game where a player receives two distinct Wikipedia pages, for example, [Eiffel Tower](https://en.wikipedia.org/wiki/Eiffel_Tower) and [Soybean](https://en.wikipedia.org/wiki/Soybean).
The goal is to find a path between the two pages using only Wikipedia links within each page.
A possible solution would be the path: [Eiffel Tower](https://en.wikipedia.org/wiki/Eiffel_Tower) &#8594; [List of tallest towers](https://en.wikipedia.org/wiki/List_of_tallest_towers) &#8594; [Brazil](https://en.wikipedia.org/wiki/Brazil) &#8594; [Soybean](https://en.wikipedia.org/wiki/Soybean)

## Search Algorithms
These are the search algorithms that are currently implemented (more to come!)

### Depth-First Search (DFS)
Depth-First Search (DFS) is a search algorithm that explores as far as possible along each branch before backtracking.
Since Wikipedia Pages form a [Cyclic Graph](https://en.wikipedia.org/wiki/Cyclic_graph), this is in fact a **Depth-Limited Search** to avoid an infinite search.
The maximum depth is given as input.

### Breadth-First Search (BFS)
Breadth-First Search (BFS) is a search algorithm that explores level by level, visiting all nodes at a certain depth level before moving on to the next.

### Bidirectional Breadth-First Search (BDBFS)
Bidirectional Breadth-First Search (BDBFS) is a search algorithm that simultaneously explores from the start and goal nodes, meeting in the middle to find the path.
This algorithm is used with Undirected Graphs since there is the need to find a path starting from the end that can then be inverted.
Wikipedia is instead a Directed Graph, meaning that A can have a link to B but not necessarily the other way around.
To adapt for this, the algorithm is altered so that every time a path from the end to a common middle point is found, it verifies if said path exists in the reverse order.
This adaption works for this specific problem since there is a moderately high level of probability (although not too high) that if A has a link to B, B will also have a link to A.

### Differences Between Search Algorithms
The **Bidirectional Breadth-First Search** will be, in most cases, the fastest algorithm. Its drawbacks are that it doesn't guarantee the shortest path and finds bidirectional paths.
The **Breadth-First Search** is the only one that guarantees finding the shortest path. It works well on paths with a length of 4 or less, but after that, the search tree gets too big.
The **Depth-First Search** can be quite effective given two conditions: i) the order in which it searches the links is favorable; and ii) we already know the length of the shortest path.

The following table shows the results and times of the various algorithms across different games (DFS maximum depth corresponds to the depth of the shortest path).

|                       Game                       |                                                                  DFS                                                                 | Time |                                                                  BFS                                                                 | Time |                                                                                    BDBFS                                                                                   | Time |
|:------------------------------------------------:|:------------------------------------------------------------------------------------------------------------------------------------:|:----:|:------------------------------------------------------------------------------------------------------------------------------------:|:----:|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|:----:|
|   Jonas Brothers <br/> &#8595; <br/> Kofi Annan  |                               Jonas_Brothers <br/> &#8595; <br/> We_Day <br/> &#8595; <br/> Kofi_Annan                               |  15s |                               Jonas_Brothers <br/> &#8595; <br/> We_Day <br/> &#8595; <br/> Kofi_Annan                               |  20s | Jonas_Brothers <br/> &#8595; <br/> Polydor_Records <br/> &#8595; <br/> Portugal <br/> &#8595; <br/> Secretary-General_of_the_United_Nations <br/> &#8595; <br/> Kofi_Annan |  7s  |
| Incandescent light bulb <br/> &#8595; <br/>Logic | Incandescent_light_bulb <br/> &#8595; <br/> Gender_of_connectors_and_fasteners <br/> &#8595; <br/> Machine <br/> &#8595; <br/> Logic |  7s  | Incandescent_light_bulb <br/> &#8595; <br/> Gender_of_connectors_and_fasteners <br/> &#8595; <br/> Machine <br/> &#8595; <br/> Logic |  57s |        Incandescent_light_bulb <br/> &#8595; <br/> Electric_light <br/> &#8595; <br/> Age_of_Enlightenment <br/> &#8595; <br/> Philosophy <br/> &#8595; <br/> Logic        |  3s  |
|           Jesus <br/> &#8595; <br/>Iron          |                                     Jesus <br/> &#8595; <br/> Sanhedrin <br/> &#8595; <br/> Iron                                     |  8s  |                                     Jesus <br/> &#8595; <br/> Sanhedrin <br/> &#8595; <br/> Iron                                     |  14s |                                         Jesus <br/> &#8595; <br/> Jerusalem <br/> &#8595; <br/> Bronze_Age <br/> &#8595; <br/> Iron                                        |  6s  |



## Usage

The search can be run using the command-line as follows:
```bash
./wikiplayer start_page goal_page search_type [depth]
```
Where:
- `start_page` and `goal_page` define the starting and ending pages of the desired path, which can be a full Wikipedia URL (https://en.wikipedia.org/wiki/Software_engineering) or just the final section of the URL representing the topic (Software_engineering).
- `search_type` defines which search algorithm to use, can be `dfs` for DFS, `bfs` for BFS, and `bi_bfs` for BDBFS.
- `depth` is only needed when using DFS and is an integer defining the maximum depth of the search.


**Examples**

```bash
./wikiplayer Jonas_Brothers Kofi_Annan bfs
```
```bash
Jonas_Brothers
We_Day
Kofi_Annan
```

```bash
./wikiplayer https://en.wikipedia.org/wiki/Incandescent_light_bulb https://en.wikipedia.org/wiki/Logic dfs 4
```
```bash
Incandescent_light_bulb
Gender_of_connectors_and_fasteners
Machine
Logic
```
