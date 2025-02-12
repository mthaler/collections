package graph

import "collections/bag"

type Graph struct {
	v   int
	E   int
	adj bag.Bag[int]
}

func NewWithVertices() *Graph {

}
