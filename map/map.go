package mymap

type HashST struct {
	n  int    // number of key-value pairs
	m  int    // number of chains
	st []Node // array of linked-list symbol tables
}

func (st *HashST) Size() int {
	return st.n
}
