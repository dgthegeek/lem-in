package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Name      string
	Neighbors []*Node
}

type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}

func (g *Graph) AddRoom(name string) {
	node := &Node{Name: name}
	g.Nodes[name] = node
}

func (g *Graph) AddArc(from, to string) {
	fromNode := g.Nodes[from]
	toNode := g.Nodes[to]

	fromNode.Neighbors = append(fromNode.Neighbors, toNode)
}

func ParseInputFile(filename string) (*Graph, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	graph := NewGraph()

	var numAnts int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "#") {
			// Ignore comments
			continue
		}

		if numAnts == 0 {
			// Parse the number of ants
			numAnts, err = strconv.Atoi(line)
			if err != nil {
				return nil, err
			}
		} else if line == "##start" || line == "##end" {
			// Ignore start and end directives for now
			continue
		} else if strings.Contains(line, "-") {
			// Parse room connections
			parts := strings.Split(line, "-")
			from := parts[0]
			to := parts[1]
			graph.AddArc(from, to)
			graph.AddArc(to, from)
		} else {
			// Parse room definition
			parts := strings.Split(line, " ")
			name := parts[0]
			graph.AddRoom(name)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return graph, nil
}

func main() {
	// Replace "input.txt" with the actual input file path
	graph, err := ParseInputFile("input.txt")
	if err != nil {
		fmt.Println("Error parsing input file:", err)
		return
	}


	for nodeName, node := range graph.Nodes {
		fmt.Printf("Node: %s, Neighbors: ", nodeName)
		for _, neighbor := range node.Neighbors {
			fmt.Printf("%s ", neighbor.Name)
		}
		fmt.Println()
	}

}
