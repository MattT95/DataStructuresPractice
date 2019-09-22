package main

import (
	"container/list"
	"fmt"
)

// Route Between Nodes
//  Given a directed graph, design an algorithm to find out whether there is a route between two nodes.
func main() {

	graph := Graph{make(map[int]*list.List)}

	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)
	graph.AddVertex(5)

	graph.AddEdge(1, 3)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)
	graph.AddEdge(1, 3)

	fmt.Println(IsRouteBetweenNodes(graph, 1, 2))
}

func IsRouteBetweenNodes(graph Graph, a int, b int) bool {

	if a == b {
		return true
	}

	adjacentNodes := *graph.Adjacent(a)

	for i := adjacentNodes.Front(); i != nil; i = i.Next() {

		if IsRouteBetweenNodes(graph, i.Value.(int), b) {
			return true
		}
	}

	return false
}

type Graph struct {
	AdjacencyLists map[int]*list.List
}

func (graph *Graph) AddVertex(a int) {
	graph.AdjacencyLists[a] = list.New()
}

func (graph *Graph) AddEdge(a int, b int) {
	adjacencyList := graph.AdjacencyLists[a]

	(*adjacencyList).PushBack(b)
}

func (graph *Graph) Adjacent(a int) *list.List {
	return graph.AdjacencyLists[a]
}
