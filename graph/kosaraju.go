package graph

func (g *Graph) Kosaraju(ch chan *Graph) {
	defer close(ch)

	// Get the finishing ordering of the nodes
	ordering := computeOrdering(g)

	// Run modified DFS to get the SCCs
	for _, node := range ordering {
		g := NewGraph()
		addFromDFS(g, node)

		if len(g.Nodes) > 0 {
			ch <- g
		}
	}
}

func computeOrdering(g *Graph) []*Node {
	ordering := make([]*Node, len(g.Nodes))
	index := len(g.Nodes) - 1

	// Reverse all edges in the graph
	g.Reverse()
	g.Reset()

	for _, node := range g.Nodes {
		compute(ordering, &index, node)
	}

	// Get the edges back to their original directions
	g.Reverse()
	g.Reset()

	return ordering
}

func compute(ordering []*Node, index *int, node *Node) {
	if node.Visited {
		return
	}

	node.Visited = true

	for _, next := range node.Nodes {
		compute(ordering, index, next)
	}

	ordering[*index] = node
	*index--
}

func addFromDFS(g *Graph, node *Node) {
	if node.Visited {
		return
	}

	node.Visited = true
	g.AddNodes(node)

	for _, next := range node.Nodes {
		addFromDFS(g, next)
	}
}
