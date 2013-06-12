package graph

import (
	"testing"
)

func TestGraphSimple(t *testing.T) {
	a := NewNode(1)
	b := NewNode(2)
	g := NewGraph(a, b)

	if len(g.Nodes) != 2 {
		t.Error()
	}
}

func TestDFS(t *testing.T) {
	a := NewNode(1)
	b := NewNode(2)
	c := NewNode(3)
	d := NewNode(4)
	e := NewNode(5)

	a.AddEdges(b, e)
	b.AddEdges(c, d)
	c.AddEdges(a) // add a cycle, just because

	g := NewGraph(a, b, c, d, e)

	ch := make(chan *Node)

	go g.DFS(ch)

	if !checkOrdering(ch, a, b, c, d, e) {
		t.Error()
	}

	if !(len(ch) == 0) {
		t.Error()
	}
}

func TestReset(t *testing.T) {
	a := NewNode(1)
	b := NewNode(2)
	c := NewNode(3)

	a.AddEdges(b)
	b.AddEdges(c)
	c.AddEdges(a)

	g := NewGraph(a, b, c)

	ch := make(chan *Node)

	go g.DFS(ch)
	if !checkOrdering(ch, a, b, c) {
		t.Error()
	}

	ch = make(chan *Node)
	go g.DFS(ch)
	if !checkOrdering(ch, a, b, c) {
		t.Error()
	}
}

func TestReverse(t *testing.T) {
	a := NewNode(1)
	b := NewNode(2)
	c := NewNode(3)

	a.AddEdges(b)
	b.AddEdges(c)
	c.AddEdges(a)

    if !testReverse(a, b, c) {
        t.Error()
    }
}

func TestReverseTwo(t *testing.T) {
	a := NewNode(1)
	b := NewNode(2)
    c := NewNode(3)
    
	a.AddEdges(b)
	b.AddEdges(a)

    if !testReverse(a, b, c) {
        t.Error()
    }
}

func TestReverseThree(t *testing.T) {
	a := NewNode(1)
	b := NewNode(2)
	c := NewNode(3)
	d := NewNode(4)
	e := NewNode(5)
	f := NewNode(6)

	a.AddEdges(b)
	b.AddEdges(a)
	b.AddEdges(c)
	c.AddEdges(a)
	d.AddEdges(e)
	e.AddEdges(f)
	f.AddEdges(d)
	a.AddEdges(d)

    if !testReverse(a, b, c, d, e, f) {
        t.Error()
    }
}

func TestReverseFour(t *testing.T) {
	a := NewNode(1)
	b := NewNode(2)
	c := NewNode(3)
	d := NewNode(4)
	
	nodes := []*Node{a, b, c, d}
	
	for _, n1 := range nodes {
	    for _, n2 := range nodes {
	        n1.AddEdges(n2)
	    }
	}
	
	if !testReverse(nodes...) {
	    t.Error()
	}
}

func testReverse(nodes ...*Node) (bool) {
    forward := NewGraph(nodes...)
    reverse := forward.Copy()
    reverse.Reverse()
    
    if !isReverseOf(forward, reverse) {
        return false
    }
    
    reverse.Reverse()
    
    if !isEqualTo(forward, reverse) {
        return false
    }
    
    return true
}

func isReverseOf(lhs *Graph, rhs *Graph) bool {
	for _, lhsNode := range lhs.Nodes {
		for _, lhsNext := range lhsNode.Nodes {
			rhsNode, _ := rhs.Nodes[lhsNode.Id]
			rhsNext, _ := rhs.Nodes[lhsNext.Id]

			if !rhsNext.Adjacent(rhsNode) {
				return false
			}
		}
	}

	return true
}

func isEqualTo(lhs *Graph, rhs *Graph) bool {
	for _, lhsNode := range lhs.Nodes {
		for _, lhsNext := range lhsNode.Nodes {
			rhsNode, _ := rhs.Nodes[lhsNode.Id]
			rhsNext, _ := rhs.Nodes[lhsNext.Id]

			if !rhsNode.Adjacent(rhsNext) {
				return false
			}
		}
	}

	return true
}

func checkOrdering(ch chan *Node, nodes ...*Node) bool {
	for _, node := range nodes {
		res := <-ch
		if !(res == node) {
			return false
		}
	}

	return true
}
