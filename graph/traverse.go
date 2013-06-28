package graph

func (g *Graph) DepthFirstSearch(f func(*Node)) {
	g.Reset()

	for _, node := range g.Nodes {
		DepthFirstSearchNode(node, f)
	}
}

func DepthFirstSearchNode(node *Node, f func(*Node)) {
	if node.Visited {
		return
	}

	node.Visited = true

	for _, next := range node.Nodes {
		DepthFirstSearchNode(next, f)
	}

	f(node)
}
