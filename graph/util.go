package graph

func expandNode(node *Node) []*Node {
	nodes := make([]*Node, len(node.Nodes))
	id := 0

	for _, node := range node.Nodes {
		nodes[id] = node
		id++
	}

	return nodes
}
