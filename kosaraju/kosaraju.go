package kosaraju

import (
	"github.com/mattomatic/go-graph/graph"
)

func Kosaraju(g *graph.Graph) chan *graph.Graph {
	ch := make(chan *graph.Graph)
	go kosaraju(g, ch)
	return ch
}

func kosaraju(g *graph.Graph, ch chan *graph.Graph) {
	defer close(ch)

	// Get the finishing ordering of the nodes
	ordering := computeOrdering(g)

	// Run modified DFS to get the SCCs
	for _, node := range ordering {
		g := graph.NewGraph()

		f := func(node *graph.Node) {
			g.AddNodes(node)
		}

		graph.DepthFirstSearchNode(node, f)

		if len(g.Nodes) > 0 {
			ch <- g
		}
	}
}

func computeOrdering(g *graph.Graph) []*graph.Node {
	ordering := make([]*graph.Node, len(g.Nodes))
	index := len(g.Nodes) - 1
	indexRef := &index

	f := func(node *graph.Node) {
		ordering[*indexRef] = node
		*indexRef--
	}

	g.Reverse()
	g.DepthFirstSearch(f)
	g.Reverse()
	g.Reset()

	return ordering
}
