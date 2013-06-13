package graph

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