package main

import (
	"fmt"
	"math"
)

func (g *Graph) runDijkstra(startingVertex string) []selectedVertex {
	// Stores the selected vertexes as we iterate through the graph
	var selectedVertexes []selectedVertex

	// create slice of nodes and remove starting node from nodes to check
	nodes := g.nodes
	deleteElementFromSlice(&nodes, startingVertex)

	// Create every node as infinity to begin with.
	nodesToWeight := make(nodeEdge)
	for _, xy := range nodes {
		nodesToWeight[xy] = math.Inf(1)
	}

	return dijkstraAlgorithm(startingVertex, &nodes, g, &selectedVertexes, &nodesToWeight)
}

// runs the dijkstra algorithm against the graph
func dijkstraAlgorithm(currentVertex string, nodes *[]string, graph *Graph, selectedVertexes *[]selectedVertex, nodesToWeightMap *nodeEdge) []selectedVertex {
	if len(*nodes) < 1 {
		return *selectedVertexes
	}

	for _, val := range *nodes {
		value, ok := graph.structure[currentVertex][val]
		if ok {
			potentialWeight := calculateNewPathWeight(selectedVertexes, value)
			if potentialWeight < (*nodesToWeightMap)[val] {
				(*nodesToWeightMap)[val] = potentialWeight
			}
		}
	}

	toInject := selectedVertex{weight: math.Inf(1)}
	for key, value := range *nodesToWeightMap {
		if value < toInject.weight {
			toInject = selectedVertex{node: key, weight: float64(value)}
		}
	}

	// If we can't pick one, then we need to return early - the values will be the last value they were written
	// this will likely be infinity if it cannot be reached
	if toInject.node == "" {
		for nsKey, nsVal := range *nodesToWeightMap {
			*selectedVertexes = append(*selectedVertexes, selectedVertex{node: nsKey, weight: nsVal})
		}
		return *selectedVertexes
	}

	*selectedVertexes = append(*selectedVertexes, toInject)

	deleteElementFromSlice(nodes, toInject.node)
	delete(*nodesToWeightMap, toInject.node)

	fmt.Println(*nodes)
	fmt.Println(*selectedVertexes)
	fmt.Println(*nodesToWeightMap)
	return dijkstraAlgorithm(toInject.node, nodes, graph, selectedVertexes, nodesToWeightMap)
}
