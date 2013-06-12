package graph

type Graph struct {
	Nodes map[int]*Node
}

type Edge struct {
    src *Node
    dst *Node
}

func NewGraph(nodes ...*Node) *Graph {
	g := &Graph{make(map[int]*Node)}
	g.AddNodes(nodes...)
	return g
}

func (g *Graph) AddNodes(nodes ...*Node) {
	for _, node := range nodes {
		g.Nodes[node.Id] = node
	}
}

func (g *Graph) Reset() {
	for _, node := range g.Nodes {
		node.Visited = false
	}
}       

func (g *Graph) ContainsIds(ids ...int) bool {
	for _, id := range ids {
		_, ok := g.Nodes[id]

		if !ok {
			return false
		}
	}

	return true
}

func (g *Graph) Contains(nodes ...*Node) bool {
	for _, node := range nodes {
		_, ok := g.Nodes[node.Id]
		if !ok {
			return false
		}
	}

	return true
}

// Perform a depth-first search of the graph
func (g *Graph) DFS(ch chan *Node) {
	defer close(ch)
	g.Reset()

	for _, node := range g.Nodes {
		dfs(node, ch)
	}
}

func (g *Graph) Copy() (*Graph) {
    m := NewGraph()
    edges := g.GetEdges()
    
    for edge := range edges {
        if !m.ContainsIds(edge.src.Id) {
            m.AddNodes(NewNode(edge.src.Id))
        }
        
        if !m.ContainsIds(edge.dst.Id) {
            m.AddNodes(NewNode(edge.dst.Id))
        }
        
        src, _ := m.Nodes[edge.src.Id]
        dst, _ := m.Nodes[edge.dst.Id]
        
        src.AddEdges(dst)
    }
    
    return m
}

func (g *Graph) GetEdges() (chan *Edge) {
    edges := make(chan *Edge)
    
    go func() {
        defer close(edges)
        
        for _, node := range g.Nodes {
            for _, next := range node.Nodes {
                edges <- &Edge{node, next}
            }
        }
    }()        
    
    return edges
}

// Reverse all the edges in a graph
func (g *Graph) Reverse() {
	g.Reset()

	for _, node := range g.Nodes {
		reverse(node)
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

func reverse(node *Node) {
	if node.Visited {
		return
	}

	node.Visited = true

	for _, next := range expandNode(node) {
		reverse(next)
		
		// Only reverse directed edges
		if !next.Adjacent(node) {
		    next.AddEdges(node)
		    node.RemoveEdges(next)
		}
	}
}
