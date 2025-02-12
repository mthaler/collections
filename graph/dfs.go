package graph

import "collections/util"

type DFS struct {
	marked []bool
	count  int
}

func NewDFS(g Graph) *DFS {
	return &DFS{marked: make([]bool, g.V)}
}

func (s *DFS) dfs(g Graph, v int) {
	s.marked[v] = true
	s.count++
	for _, w := range util.ToSclice[int](g.adj[v]) {
		if !s.marked[w] {
			s.dfs(g, w)
		}
	}
}
