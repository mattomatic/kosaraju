package graph

type Node struct {
	Id      int
	Visited bool
	Nodes   map[int]*Node
}

func NewNode(id int, edges ...*Node) *Node {
	node := &Node{Id: id, Nodes: make(map[int]*Node)}
	node.AddEdges(edges...)
	return node
}

func (n *Node) AddEdges(edges ...*Node) {
	for _, edge := range edges {
		n.Nodes[edge.Id] = edge
	}
}

func (n *Node) RemoveEdges(edges ...*Node) {
	for _, edge := range edges {
		delete(n.Nodes, edge.Id)
	}
}

func (n *Node) Adjacent(edges ...*Node) bool {
	for _, edge := range edges {
		_, ok := n.Nodes[edge.Id]
		if !ok {
			return false
		}
	}

	return true
}
