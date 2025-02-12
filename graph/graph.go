package graph

import "collections/bag"

type Graph struct {
	V   int
	E   int
	adj []bag.Bag[int]
}

func New(v int) *Graph {
	return &Graph{
		V:   v,
		E:   0,
		adj: make([]bag.Bag[int], v),
	}
}

func (g Graph) AddEdge(v, w int) {
	g.adj[v].Add(w) // Add w to v’s list.
	g.adj[w].Add(v) // Add v to w’s list.
}
