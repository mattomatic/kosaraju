package graph

func (g *Graph) Kosaraju() (chan *Graph) {
    ch := make(chan *Graph)
    go kosaraju(g, ch)
    return ch
}

func kosaraju(g *Graph, ch chan *Graph) {
	defer close(ch)

	// Get the finishing ordering of the nodes
	ordering := computeOrdering(g)
	
	// Run modified DFS to get the SCCs
	for _, node := range ordering {
		g := NewGraph()
		
		f := func (node *Node) {
		    g.AddNodes(node)
		}
		
		traverseDFS(node, f)
		
		if len(g.Nodes) > 0 {
            ch <- g
        }
	}
}

func computeOrdering(g *Graph) []*Node {
	ordering := make([]*Node, len(g.Nodes))
	index := len(g.Nodes) - 1
	indexRef := &index
	
    f := func(node *Node) {
        ordering[*indexRef] = node
        *indexRef--
    }
    
    g.Reverse()
    g.TraverseDFS(f)
    g.Reverse()
    g.Reset()
    
    return ordering
}