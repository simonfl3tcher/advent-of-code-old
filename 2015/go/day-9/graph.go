package main

type nodeEdge map[string]float64

// Graph struct for building graph representation
type Graph struct {
	structure map[string]nodeEdge
	nodes     []string
	directed  bool
}

type selectedVertex struct {
	node   string
	weight float64
}

func (g *Graph) addEdge(source, target string, weight int) {
	if _, ok := g.structure[source]; !ok {
		m := make(nodeEdge)
		m[target] = float64(weight)
		g.structure[source] = m
	} else {
		g.structure[source][target] = float64(weight)
	}

	// Non directed graph
	if !g.directed {
		if _, ok := g.structure[target]; !ok {
			m := make(nodeEdge)
			m[source] = float64(weight)
			g.structure[target] = m
		} else {
			g.structure[target][source] = float64(weight)
		}
	}
	// End code for non directed graph

	// Adds source + target to nodes
	if ok := contains(g.nodes, source); !ok {
		g.nodes = append(g.nodes, source)
	}
	if ok := contains(g.nodes, target); !ok {
		g.nodes = append(g.nodes, target)
	}
}

func calculateNewPathWeight(selected *[]selectedVertex, val float64) float64 {
	var weight float64
	if len(*selected) > 0 {
		weight = (*selected)[len(*selected)-1].weight
	}
	return weight + val
}

func deleteElementFromSlice(nodes *[]string, vertex string) {
	for i, v := range *nodes {
		if v == vertex {
			*nodes = append((*nodes)[:i], (*nodes)[i+1:]...)
			return
		}
	}
}

// Contains tells whether a contains x.
func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
