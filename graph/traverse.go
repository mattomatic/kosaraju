package graph

// Perform a depth-first search of the graph
func (g *Graph) DFS(ch chan *Node) {
	defer close(ch)
	g.Reset()

	for _, node := range g.Nodes {
		dfs(node, ch)
	}
}

func dfs(n *Node, ch chan *Node) {
	if n.Visited {
		return
	}

	n.Visited = true
	ch <- n

	for _, next := range n.Nodes {
		dfs(next, ch)
	}
}

func (g *Graph) TraverseDFS(f func (*Node)) {
    g.Reset()
    
    for _, node := range g.Nodes {
        traverseDFS(node, f)
    }
}

func traverseDFS(node *Node, f func (*Node)) {
    if node.Visited {
        return
    }
    
    node.Visited = true
    
    for _, next := range node.Nodes {
        traverseDFS(next, f)
    }
    
    f(node)
}