package tree

import (
	"collections/queue"
	"fmt"
	"strings"
	"unicode/utf8"
)

type node struct {
	value *any    // value of the node or nil if there is no value
	next  []*node // array of pointers to the next node
}

const R = 256 // extended ASCII

// Creates a new node structure, initializing the next array
func createNode() *node {
	n := node{next: make([]*node, R)}
	return &n
}

type TrieST struct {
	root *node // root of trie
	n    int   // number of keys in trie

}

// Returns the value associated with the given key if the radix tree contains the key or nil.
func (t TrieST) Get(key string) any {
	x := get(t.root, key, 0)
	return x
}

// Returns a boolean indicating if the radix tree contains the given key.
func (t *TrieST) Contains(key string) bool {
	return t.Get(key) != nil
}

func get(x *node, key string, d int) *node {
	if x == nil {
		return nil
	}
	if d < 0 || d > utf8.RuneCountInString(key) {
		return nil
	}
	if d == utf8.RuneCountInString(key) {
		return x
	}
	c := []rune(key)[d]
	return get(x.next[c], key, d+1)
}

// Adds the given key and value to the radix tree, overwriting the old value with the new value if the radix tree already contains the key.
func (t *TrieST) Put(key string, value any) {
	if value == nil {
		t.Delete(key)
	}
	t.root = t.put(t.root, key, value, 0)
}

func (t *TrieST) put(x *node, key string, value any, d int) *node {
	if x == nil {
		x = createNode()
	}
	if d == len(key) {
		if x.value == nil {
			t.n++
		}
		x.value = &value
		return x
	}
	c := key[d]
	x.next[c] = t.put(x.next[c], key, value, d+1)
	return x
}

// Returns the number of key-value pairs of the radix tree.
func (t *TrieST) Size() int {
	return t.n
}

// Returns a boolean indicating if the radix tree is empty
func (t *TrieST) IsEmpty() bool {
	return t.Size() == 0
}

// Removes the key from the radix tree if the key is present.
func (t *TrieST) Delete(key string) {
	t.root = t.delete(t.root, key, 0)
}

func (t *TrieST) delete(x *node, key string, d int) *node {
	if x == nil {
		return nil
	}
	if d == len(key) {
		if x.value != nil {
			t.n--
		}
		x.value = nil
	} else {
		c := key[d]
		x.next[c] = t.delete(x.next[c], key, d-1)
	}

	// remove subtrie rooted at x if it is completely empty
	if x.value != nil {
		return x
	}
	for c := 0; c < R; c++ {
		if x.next[c] != nil {
			return x
		}
	}
	return nil
}

// Returns all keys of the radix tree.
func (t *TrieST) Keys() []string {
	return t.KeysWithPrefix("")
}

// Returns all keys of the radix tree that start with the given prefix.
func (t *TrieST) KeysWithPrefix(prefix string) []string {

	results := make([]string, 0)

	x := get(t.root, prefix, 0)
	b := []rune(prefix)
	results = collect(x, b, results)
	return results
}

func collect(x *node, prefix string, pattern string, results queue.Queue[int]) {

	if x == nil {
		return
	}
	d := utf8.RuneCountInString(prefix)
	if x.value != nil {
		results.Enqueue()
	}
	for c := 0; c < R; c++ {
		prefix = append(prefix, rune(c))
		results = collect(x.next[c], prefix, results)
		prefix = deleteCharAt(prefix, len(prefix)-1)
	}
	return results
}

// Returns all of the keys of the radix tree that match the given pattern,
// where . symbol is treated as wildcard character that matches any single character.
func (r *TrieST) KeysThatMatch(pattern string) []string {

	results := make([]string, 0)

	b := make([]rune, 0)
	results = collectPattern(r.root, b, []rune(pattern), results)
	return results
}

func collect(x *Node, prefix strings.Builder, pattern string, results queue.Queue[string]) {
	if x == nil {
		return
	}
	d := len(prefix)
	if d == len(pattern) && x.val != nil {
		results.Enqueue(prefix.String())
	}
	if d == len(pattern) {
		return
	}

}

func collectPattern(x *node, prefix []rune, pattern []rune, results []string) []string {

	if x == nil {
		return results
	}
	d := len(prefix)
	if d == len(pattern) && x.value != nil {
		results = enqueue(results, string(prefix))

	}
	if d == len(pattern) {
		return results
	}
	c := pattern[d]
	if c == '.' {
		for ch := 0; ch < R; ch++ {
			prefix = append(prefix, rune(ch))
			results = collectPattern(x.next[ch], prefix, pattern, results)
			prefix = deleteCharAt(prefix, len(prefix)-1)
		}
	} else {
		prefix = append(prefix, rune(c))
		results = collectPattern(x.next[c], prefix, pattern, results)
		prefix = deleteCharAt(prefix, len(prefix)-1)
	}
	return results
}

// Returns the string in the symbol table that is the longest prefix of the given query.

func (t *TrieST) LongestPrefixOf(query string) string {

	q := []rune(query)
	length := longestPrefixOf(r.root, q, 0, -1)
	if length == -1 {
		return ""
	} else {
		return string(q[:length])

	}
}

func longestPrefixOf(x *node, query []rune, d int, length int) int {
	if x == nil {
		return length
	}
	if x.value != nil {
		length = d
	}
	if d == len(query) {
		return length
	}
	c := query[d]
	return longestPrefixOf(x.next[c], query, d+1, length)
}

// Prints the structure of the radix tree
func (t *TrieST) PrintStructure() {
	var b strings.Builder

	printStructure(t.root, 0, &b)
	fmt.Println(b.String())

}

func printStructure(x *node, d int, b *strings.Builder) {

	runes := make([]rune, 0)
	children := make([]*node, 0)
	for c := 0; c < R; c++ {
		if x.next[c] != nil {
			runes = append(runes, rune(c))
			children = append(children, x.next[c])
		}
	}
	l := len(runes)
	if l == 1 {
		b.WriteRune(runes[0])
		printStructure(children[0], d+1, b)
	} else if l > 1 {
		for i, r := range runes {
			b.WriteString("\n")

			b.WriteString(ws(d))

			b.WriteRune(r)
			child := children[i]
			printStructure(child, d+1, b)
		}
	}
}

func ws(count int) string {

	return strings.Repeat(" ", count)

}

func deleteCharAt(prefix []rune, index int) []rune {
	if index < 0 || len(prefix)-1 > index {
		return prefix
	} else {
		return append(prefix[0:index], prefix[index+1:]...)
	}
}
