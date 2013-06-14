package main

import (
	"flag"
	"fmt"
	"github.com/mattomatic/go-graph/graph"
	"sort"
	"time"
)

func topN(items []int, n int) []int {
	result := make([]int, n)
	sort.Ints(items)

	for i := 0; i < len(items); i++ {
		result[i] = items[i]
	}

	return result
}

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		flag.PrintDefaults()
		return
	}

	filename := flag.Args()[0]

	graphTime := time.Now()
	g := graph.LoadGraph(filename)
	graphTimer := time.Now().Sub(graphTime)

	sccTime := time.Now()
	sccs := Kosaraju(g)

	lengths := make([]int, 0)
	for scc := range sccs {
		lengths = append(lengths, len(scc.Nodes))
	}
	sccTimer := time.Now().Sub(sccTime)

	answer := topN(lengths, len(lengths))
	fmt.Printf("graphTime: %v\n", graphTimer)
	fmt.Printf("sccTime: %v\n", sccTimer)
	fmt.Printf("answer: %v\n", answer)
}
