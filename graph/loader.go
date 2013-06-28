package graph

import (
	"bufio"
	"fmt"
	"os"
)

func LoadGraph(filename string) *Graph {
	fp, _ := os.Open(filename)
	reader := bufio.NewReader(fp)
	g := NewGraph()
	readHeader(reader)
	var src, dst int

	for scanLine(reader, &src, &dst) {
		srcNode := getOrCreate(g, src)
		dstNode := getOrCreate(g, dst)
		srcNode.AddEdges(dstNode)
	}

	return g
}

func scanLine(reader *bufio.Reader, src *int, dst *int) bool {
	n, err := fmt.Fscanf(reader, "%d %d\n", src, dst)
	return n == 2 && err == nil
}

func readHeader(reader *bufio.Reader) {
	b, err := reader.Peek(1)
	
	if err != nil {
		panic("could not peek file!")
	}
	
	if b[0] == '#' {
		reader.ReadLine()
	}
}

func getOrCreate(g *Graph, id int) *Node {
	if !g.ContainsIds(id) {
		g.AddNodes(NewNode(id))
	}

	return g.Nodes[id]
}
