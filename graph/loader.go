package graph

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func LoadGraph(filename string) (*Graph) {
	g := NewGraph()
	ch := make(chan string)

	go loadLines(filename, ch)

	for line := range ch {
		loadLine(g, line)
	}

    return g
}

func loadLines(filename string, ch chan string) {
    defer close(ch)
    
	f, _ := os.Open(filename)
	r := bufio.NewReader(f)

	for line, _, err := r.ReadLine(); err != io.EOF; line, _, err = r.ReadLine() {
	    s := string(line)
	    
	    if !strings.HasPrefix(s, "#") {
		    ch <- s
		}
	}
}

func loadLine(g *Graph, line string) {
    fields := strings.Split(line, " ")
	srcId, _ := strconv.Atoi(fields[0])
	dstId, _ := strconv.Atoi(fields[1])
	srcNode := getOrCreate(g, srcId)
	dstNode := getOrCreate(g, dstId)
	srcNode.AddEdges(dstNode)
}

func getOrCreate(g *Graph, id int) (*Node) {
    if !g.ContainsIds(id) {
        g.AddNodes(NewNode(id))
    }
    
    return g.Nodes[id]
}
