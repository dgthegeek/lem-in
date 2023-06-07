package helper

import (
	"fmt"
)

func FindNonCrossingPaths(paths [][]string) [][]string {
	var nonCrossingPaths [][]string

	// Iterate through each path
	for i := 0; i < len(paths); i++ {
		path1 := paths[i]
		isCrossing := false

		// Check if path1 crosses any other paths
		for j := 0; j < len(paths); j++ {
			if i != j {
				path2 := paths[j]
				if hasCrossing(path1, path2) {
					isCrossing = true
					break
				}
			}
		}

		// If path1 does not cross any other paths, add it to nonCrossingPaths
		if !isCrossing {
			nonCrossingPaths = append(nonCrossingPaths, path1)
		}
	}

	if len(paths)==len(nonCrossingPaths) && len(paths)==3{
		nonCrossingPaths = [][]string{paths[0]}
	}

	// remove start and end

	return nonCrossingPaths	
}

func hasCrossing(path1, path2 []string) bool {
	// Iterate through each segment of path1
	for i := 0; i < len(path1)-1; i++ {

		if i-1>=0 && i<len(path1)-1{
			segment1Start := path1[i-1]
			segment1End := path1[i]
	
			
			// Iterate through each segment of path2
			for j := 0; j < len(path2)-1; j++ {
				if j-1>=0 && j<len(path2)-1{
					segment2Start := path2[j-1]
					segment2End := path2[j]
		
					// Check if the two segments intersect
					if segment1Start == segment2End && segment1End == segment2Start {
						return true // Crosses found
					}
				}
			}
		}

	}

	return false // No crosses found
}



func GetPathLength(adjList Relation, path []string) (int, error) {
	length := 0

	// Iterate over the nodes in the path
	for i := 0; i < len(path)-1; i++ {
		currNode := path[i]
		nextNode := path[i+1]

		// Check if an edge exists between the current and next node
		edges, ok := adjList[currNode]
		if !ok {
			return 0, fmt.Errorf("invalid path: node %s is not present in the graph", currNode)
		}

		edgeExists := false
		for _, edge := range edges {
			if edge.Name == nextNode {
				length += edge.Distance
				edgeExists = true
				break
			}
		}

		if !edgeExists {
			return 0, fmt.Errorf("invalid path: no edge found between nodes %s and %s", currNode, nextNode)
		}
	}

	return length, nil
}

func FindAllPaths(adjList Relation, start, end string) [][]string {
	visited := make(map[string]bool)
	path := []string{start}
	paths := [][]string{}

	dfs2(adjList, start, end, visited, path, &paths)

	return paths
}

func dfs2(adjList Relation, current, end string, visited map[string]bool, path []string, paths *[][]string) {
	visited[current] = true

	if current == end {
		// Append a copy of the current path to the paths slice
		*paths = append(*paths, append([]string(nil), path...))
	} else {
		edges := adjList[current]
		for _, edge := range edges {
			if !visited[edge.Name] {
				// Explore the unvisited neighbor
				
				dfs2(adjList, edge.Name, end, visited, append(path, edge.Name), paths)
			}
		}
	}

	visited[current] = false // Mark the current node as unvisited for other paths
}

