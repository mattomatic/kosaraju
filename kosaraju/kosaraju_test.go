package kosaraju

import (
	"testing"
	"github.com/mattomatic/go-graph/graph"
)

func TestKosaraju(t *testing.T) {
	a := graph.NewNode(1)
	b := graph.NewNode(2)
	c := graph.NewNode(3)

	d := graph.NewNode(4)
	e := graph.NewNode(5)
	f := graph.NewNode(6)

	a.AddEdges(b)
	b.AddEdges(c)
	c.AddEdges(a)

	d.AddEdges(e)
	e.AddEdges(f)
	f.AddEdges(d)

	g := graph.NewGraph(a, b, c, d, e, f)

	sccs := Kosaraju(g)
	scc1 := <-sccs
	scc2 := <-sccs

	if !isComponent(scc1, d, e, f) {
		t.Error()
	}

	if !isComponent(scc2, a, b, c) {
		t.Error()
	}
}

func TestKosarajuTwo(t *testing.T) {
	a := graph.NewNode(1)
	b := graph.NewNode(2)
	c := graph.NewNode(3)
	d := graph.NewNode(4)
	e := graph.NewNode(5)
	f := graph.NewNode(6)

	a.AddEdges(b)
	b.AddEdges(c)
	c.AddEdges(a)

	d.AddEdges(e)
	e.AddEdges(f)
	f.AddEdges(d)

	a.AddEdges(d) // Extra edges from c2 to c1
	a.AddEdges(e)
	a.AddEdges(f)

	g := graph.NewGraph(a, b, c, d, e, f)
	sccs := Kosaraju(g)
	scc1 := <-sccs
	scc2 := <-sccs

	if !isComponent(scc1, d, e, f) {
		t.Error()
	}

	if !isComponent(scc2, a, b, c) {
		t.Error()
	}
}

func TestSimple(t *testing.T) {
	a := graph.NewNode(1)
	b := graph.NewNode(2)
	c := graph.NewNode(3)
	d := graph.NewNode(4)
	e := graph.NewNode(5)
	f := graph.NewNode(6)

	a.AddEdges(b)
	b.AddEdges(c)
	c.AddEdges(a)

	d.AddEdges(e)
	e.AddEdges(f)
	f.AddEdges(d)

	a.AddEdges(d)

	g := graph.NewGraph(a, b, c, d, e, f)

	expected := [...]*graph.Node{d, f, e, a, c, b}
	ordering := computeOrdering(g)

	for i := 0; i < len(expected); i++ {
		if expected[i] != ordering[i] {
			t.Error()
		}
	}
}

func isComponent(g *graph.Graph, nodes ...*graph.Node) bool {
	if len(g.Nodes) != len(nodes) {
		return false
	}

	for _, node := range nodes {
		if !g.Contains(node) {
			return false
		}
	}

	return true
}
