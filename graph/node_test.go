package graph

import (
	"testing"
)

func TestMakeNodes(t *testing.T) {
	a := NewNode(1)

	if a.Id != 1 {
		t.Error()
	}

}

func TestMakeEdges(t *testing.T) {
	a := NewNode(1)
	b := NewNode(2)
	a.AddEdges(b)

	if len(a.Nodes) != 1 {
		t.Error()
	}

	if a.Nodes[b.Id] != b {
		t.Error()
	}

	if !(a.Adjacent(b)) {
		t.Error()
	}

	if b.Adjacent(a) {
		t.Error()
	}
}
