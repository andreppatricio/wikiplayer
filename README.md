# The Wikipedia Game Player

Welcome to the **The Wikipedia Game Player** project!
This Go(lang) project gives solutions for [The Wikipedia Game](https://en.wikipedia.org/wiki/Wikipedia:Wiki_Game) through web scrapping, various search algorithms, and parallel computing (hooray for Go routines).

The Wikipedia Game is a game where a player receives two distinct Wikipedia pages, for example, [Eiffel Tower](https://en.wikipedia.org/wiki/Eiffel_Tower) and [Soybean](https://en.wikipedia.org/wiki/Soybean).
The goal is to find a path between the two pages using only Wikipedia links within each page.
A possible solution would be the path: [Eiffel Tower](https://en.wikipedia.org/wiki/Eiffel_Tower) &#8594; [List of tallest towers](https://en.wikipedia.org/wiki/List_of_tallest_towers) &#8594; [Brazil](https://en.wikipedia.org/wiki/Brazil) &#8594; [Soybean](https://en.wikipedia.org/wiki/Soybean)

## Search Algorithms
These are the search algorithms that are currently implemented

### Depth-First Search (DFS)
Depth-First Search (DFS) is a search algorithm that explores as far as possible along each branch before backtracking.
Since Wikipedia Pages form a [Cyclic Graph](https://en.wikipedia.org/wiki/Cyclic_graph), this is in fact a **Depth-Limited Search** to avoid an infinite search.
The maximum depth is given as input.

### Breadth-First Search (BFS)
Breadth-First Search (BFS) is a search algorithm that explores level by level, visiting all nodes at a certain depth level before moving on to the next.

### Bidirectional Breadth-First Search (BDBFS)
Bidirectional Breadth-First Search (BDBFS) is a search algorithm that simultaneously explores from the start and goal nodes, meeting somewhere in the middle.
This algorithm is used with Undirected Graphs since there is the need to find a path starting from the end that can then be inverted.
Wikipedia is instead a Directed Graph, meaning that A can have a link to B but not necessarily the other way around.
To adapt for this, the algorithm is altered so that every time a path from the end to a common middle point is found, it verifies if said path exists in the reverse order.
This adaption works for this specific problem since there is a moderately high level of probability (although not too high) that if A has a link to B, B will also have a link to A.

### Differences Between Search Algorithms
The **Bidirectional Breadth-First Search** is, in most cases, the fastest algorithm. Its drawbacks are that it doesn't guarantee the shortest path and only finds bidirectional paths.
The **Breadth-First Search** is the only one that guarantees finding the shortest path. It can beat **BDBFS** in some cases but it's not as consistent.
The **Depth-First Search** can be quite effective given two conditions: i) the order in which it searches the links is favorable; and ii) we already know the length of the shortest path.

The following table shows the results and times of the various algorithms across different games with single- and multi-threaded variants (the defined maximum depth for DFS corresponds to the depth of the shortest path).

|                     **Game**                     |                                                       **DFS** (Single-threaded)                                                      |                                                       **BFS** (Single-threaded)                                                      |                                                       **BFS** (Multi-threaded)                                                       |                                                                  **BDBFS** (Single-threaded)                                                                 |                                                                  **BDBFS** (Multi-threaded)                                                                 |
|:------------------------------------------------:|:------------------------------------------------------------------------------------------------------------------------------------:|:------------------------------------------------------------------------------------------------------------------------------------:|:------------------------------------------------------------------------------------------------------------------------------------:|:------------------------------------------------------------------------------------------------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------------------------------------------------------:|
|           Jesus <br/> &#8595; <br/>Iron          |                                     Jesus <br/> &#8595; <br/> Sanhedrin <br/> &#8595; <br/> Iron                                     |                                     Jesus <br/> &#8595; <br/> Sanhedrin <br/> &#8595; <br/> Iron                                     |                                     Jesus <br/> &#8595; <br/> Sanhedrin <br/> &#8595; <br/> Iron                                     |                                  Jesus <br/> &#8595; <br/> Jerusalem <br/> &#8595; <br/> Bronze_Age <br/> &#8595; <br/> Iron                                 |                           Jesus <br/> &#8595; <br/> Judaea_(Roman_province) <br/> &#8595; <br/> Iron_Age <br/> &#8595; <br/> Iron                           |
|                **Time (seconds)**                |                                                                  7.8                                                                 |                                                                  7.2                                                                 |                                                                  2.2                                                                 |                                                                              6.7                                                                             |                                                                             3.6                                                                             |
| Incandescent light bulb <br/> &#8595; <br/>Logic | Incandescent_light_bulb <br/> &#8595; <br/> Gender_of_connectors_and_fasteners <br/> &#8595; <br/> Machine <br/> &#8595; <br/> Logic | Incandescent_light_bulb <br/> &#8595; <br/> Gender_of_connectors_and_fasteners <br/> &#8595; <br/> Machine <br/> &#8595; <br/> Logic | Incandescent_light_bulb <br/> &#8595; <br/> Gender_of_connectors_and_fasteners <br/> &#8595; <br/> Machine <br/> &#8595; <br/> Logic | Incandescent_light_bulb <br/> &#8595; <br/> Electric_light <br/> &#8595; <br/> Age_of_Enlightenment <br/> &#8595; <br/> Philosophy <br/> &#8595; <br/> Logic | Incandescent_light_bulb <br/> &#8595; <br/> Vacuum <br/> &#8595; <br/> Greek_philosophy <br/> &#8595; <br/> Outline_of_philosophy <br/> &#8595; <br/> Logic |
|                **Time (seconds)**                |                                                                  5.3                                                                 |                                                                 17.9                                                                 |                                                                  3.1                                                                 |                                                                             1.53                                                                             |                                                                             3.0                                                                             |
| DNA sequencing <br/> &#8595; <br/>  Steam engine |          DNA_sequencing <br/> &#8595; <br/> Genetics <br/> &#8595; <br/> Negative_feedback <br/> &#8595; <br/> Steam_engine          |          DNA_sequencing <br/> &#8595; <br/> Genetics <br/> &#8595; <br/> Negative_feedback <br/> &#8595; <br/> Steam_engine          |       DNA_sequencing <br/> &#8595; <br/> Genetic_variation <br/> &#8595; <br/> Erasmus_Darwin <br/> &#8595; <br/> Steam_engine       |     DNA_sequencing <br/> &#8595; <br/> Genome <br/> &#8595; <br/> Homo_sapiens <br/> &#8595; <br/> History_of_technology <br/> &#8595; <br/> Steam_engine    | DNA_sequencing <br/> &#8595; <br/> Genome <br/> &#8595; <br/> Economies_of_scale <br/> &#8595; <br/> Industrial_Revolution <br/> &#8595; <br/> Steam_engine |
|                **Time (seconds)**                |                                                                 16.0                                                                 |                                                                 38.6                                                                 |                                                                 18.4                                                                 |                                                                              2.9                                                                             |                                                                             3.7                                                                             |
| Medication <br/> &#8595; <br/> Maya civilization |                               Medication <br/> &#8595; <br/> Honey <br/> &#8595; <br/ Maya_civilization                              |                               Medication <br/> &#8595; <br/> Honey <br/> &#8595; <br/ Maya_civilization                              |                               Medication <br/> &#8595; <br/> Honey <br/> &#8595; <br/ Maya_civilization                              |                      Medication <br/> &#8595; <br/> Pharmacy <br/> &#8595; <br/> Mortar_and_pestle <br/> &#8595; <br/> Maya_civilization                     |                     Medication <br/> &#8595; <br/> Pharmacy <br/> &#8595; <br/> Mortar_and_pestle <br/> &#8595; <br/> Maya_civilization                     |
|                **Time (seconds)**                |                                                                 24.6                                                                 |                                                                 22.3                                                                 |                                                                  7.3                                                                 |                                                                              3.1                                                                             |                                                                             2.4                                                                             |

## Usage

The search can be run using the command-line as follows:
```bash
./wikiplayer start_page goal_page search_type [depth] [workers]
```
Where:
- `start_page` and `goal_page` define the starting and ending pages of the desired path, which can be full Wikipedia URLs (https://en.wikipedia.org/wiki/Software_engineering) or just the final sections of the URLs representing the topic (Software_engineering).
- `search_type` defines which search algorithm to use, can be `dfs` for DFS, `bfs` for BFS, and `bdbfs` for BDBFS.
- `depth` is only needed when using DFS and is an integer defining the maximum depth of the search.
- `workers` is only needed when using BFS or BDBFS and is an integer defining the maximum number of Go routines to create during the algorithm. The default value is `1`, resulting in single-threaded execution. The actual number of cores used for parallel computation will correspond to the available number of logical CPUs, as returned by `runtime.NumCPU()`.


**Examples**

```bash
./wikiplayer Jonas_Brothers Kofi_Annan bfs 5
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
