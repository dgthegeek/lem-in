package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	nodes map[string]*Node
}

type Node struct {
	name      string
	neighbors []*Node
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]*Node),
	}
}

func (g *Graph) AddNode(name string) {
	if _, ok := g.nodes[name]; !ok {
		g.nodes[name] = &Node{
			name: name,
		}
	}
}

func (g *Graph) AddEdge(src, dst string) {
	srcNode, ok := g.nodes[src]
	if !ok {
		log.Fatalf("Node '%s' not found in the graph", src)
	}
	dstNode, ok := g.nodes[dst]
	if !ok {
		log.Fatalf("Node '%s' not found in the graph", dst)
	}

	srcNode.neighbors = append(srcNode.neighbors, dstNode)
	dstNode.neighbors = append(dstNode.neighbors, srcNode)
}

func ParseInputFile(filename string) (*Graph, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	graph := NewGraph()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Skip comments or empty lines
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}

		fields := strings.Fields(line)

		switch fields[0] {
		case "##start", "##end":
			// Skip start and end room directives for now
		case "L":
			// Skip ant movement lines for now
		default:
			if len(fields) == 3 {
				roomName := fields[0]
				graph.AddNode(roomName)
			} else if len(fields) == 1 {
				// End of room definitions, start reading links
				break
			} else {
				return nil, fmt.Errorf("Invalid input format")
			}
		}
	}

	// Parse the links between rooms
	for scanner.Scan() {
		line := scanner.Text()

		// Skip comments or empty lines
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}

		fields := strings.Split(line, "-")
		if len(fields) != 2 {
			return nil, fmt.Errorf("Invalid input format")
		}

		src := fields[0]
		dst := fields[1]

		graph.AddEdge(src, dst)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return graph, nil
}

func main() {
	graph, err := ParseInputFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Use the graph representation for further processing
	fmt.Println(graph)
}
