package tree

import (
	"collections/queue"
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

/**
 *  The {@code TrieST} class represents a symbol table of key-value
 *  pairs, with string keys and generic values.
 *  It supports the usual <em>put</em>, <em>get</em>, <em>contains</em>,
 *  <em>delete</em>, <em>size</em>, and <em>is-empty</em> methods.
 *  It also provides character-based methods for finding the string
 *  in the symbol table that is the <em>longest prefix</em> of a given prefix,
 *  finding all strings in the symbol table that <em>start with</em> a given prefix,
 *  and finding all strings in the symbol table that <em>match</em> a given pattern.
 *  A symbol table implements the <em>associative array</em> abstraction:
 *  when associating a value with a key that is already in the symbol table,
 *  the convention is to replace the old value with the new value.
 *  Unlike {@link java.util.Map}, this class uses the convention that
 *  values cannot be {@code null}—setting the
 *  value associated with a key to {@code null} is equivalent to deleting the key
 *  from the symbol table.
 *  <p>
 *  This implementation uses a 256-way trie.
 *  The <em>put</em>, <em>contains</em>, <em>delete</em>, and
 *  <em>longest prefix</em> operations take time proportional to the length
 *  of the key (in the worst case). Construction takes constant time.
 *  The <em>size</em>, and <em>is-empty</em> operations take constant time.
 *  Construction takes constant time.
 *  <p>
 *  For additional documentation, see <a href="https://algs4.cs.princeton.edu/52trie">Section 5.2</a> of
 *  <i>Algorithms, 4th Edition</i> by Robert Sedgewick and Kevin Wayne.
 *
 * Ported from Robert Sedgewicks Java code
 */
type TrieST struct {
	root *node // root of trie
	n    int   // number of keys in trie

}

/**
 * Does this symbol table contain the given key?
 * @param key the key
 * @return {@code true} if this symbol table contains {@code key} and
 *     {@code false} otherwise
 * @throws IllegalArgumentException if {@code key} is {@code null}
 */
func (t *TrieST) Contains(key string) bool {
	return get(t.root, key, 0) != nil
}

// Returns the value associated with the given key if the radix tree contains the key or nil.
func (t TrieST) Get(key string) any {
	x := get(t.root, key, 0)
	return x
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

/**
 * Inserts the key-value pair into the symbol table, overwriting the old value
 * with the new value if the key is already in the symbol table.
 * If the value is {@code null}, this effectively deletes the key from the symbol table.
 * @param key the key
 * @param val the value
 * @throws IllegalArgumentException if {@code key} is {@code null}
 */ /**
 * Inserts the key-value pair into the symbol table, overwriting the old value
 * with the new value if the key is already in the symbol table.
 * If the value is {@code null}, this effectively deletes the key from the symbol table.
 * @param key the key
 * @param val the value
 * @throws IllegalArgumentException if {@code key} is {@code null}
 */
func (t *TrieST) Put(key string, value any) {
	t.root = t.put(t.root, key, value, 0)
}

func (t *TrieST) put(x *node, key string, value any, d int) *node {
	if x == nil {
		x = createNode()
	}
	if d == utf8.RuneCountInString(key) {
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

/**
 * Returns the number of key-value pairs in this symbol table.
 * @return the number of key-value pairs in this symbol table
 */
func (t *TrieST) Size() int {
	return t.n
}

/**
 * Is this symbol table empty?
 * @return {@code true} if this symbol table is empty and {@code false} otherwise
 */
func (t *TrieST) IsEmpty() bool {
	return t.Size() == 0
}

/**
 * Returns all keys in the symbol table as an {@code Iterable}.
 * To iterate over all of the keys in the symbol table named {@code st},
 * use the foreach notation: {@code for (Key key : st.keys())}.
 * @return all keys in the symbol table as an {@code Iterable}
 */
func (t *TrieST) Keys() []string {
	return t.KeysWithPrefix("")
}

/**
 * Returns all keys in the symbol table as an {@code Iterable}.
 * To iterate over all of the keys in the symbol table named {@code st},
 * use the foreach notation: {@code for (Key key : st.keys())}.
 * @return all keys in the symbol table as an {@code Iterable}
 */
func (t *TrieST) KeysWithPrefix(prefix string) []string {

	results := queue.New()
	x := get(t.root, prefix, 0)
	var sb strings.Builder
	sb.WriteString(prefix)
	collect(x, sb, results)
	return results
}

func collect(x *node, prefix strings.Builder, results queue.Queue[string]) {
	if x == nil {
		return
	}
	if (*x).value != nil {
		results.Enqueue(prefix.String())
	}

	for c := 0; c < R; c++ {
		prefix.WriteRune(rune(c))
		collect(x.next[c], prefix, results)
		r := []rune(prefix.String())
		r = deleteCharAt(r, len(r)-1)
		prefix.Reset()
		prefix.WriteString(string(r))

	}
}

/**
 * Returns all of the keys in the symbol table that match {@code pattern},
 * where the character '.' is interpreted as a wildcard character.
 * @param pattern the pattern
 * @return all of the keys in the symbol table that match {@code pattern},
 *     as an iterable, where . is treated as a wildcard character.
 */
func (t *TrieST) KeysThatMatch(pattern string) []string {

	results := make([]string, 0)

	b := make([]rune, 0)
	results = collectPattern(t.root, b, []rune(pattern), results)
	return results
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

/**
 * Returns the string in the symbol table that is the longest prefix of {@code query},
 * or {@code null}, if no such string.
 * @param query the query string
 * @return the string in the symbol table that is the longest prefix of {@code query},
 *     or {@code null} if no such string
 * @throws IllegalArgumentException if {@code query} is {@code null}
 */
func (t *TrieST) LongestPrefixOf(query string) string {

	q := []rune(query)
	length := longestPrefixOf(r.root, q, 0, -1)
	if length == -1 {
		return ""
	} else {
		return string(q[:length])

	}
}

// returns the length of the longest string key in the subtrie
// rooted at x that is a prefix of the query string,
// assuming the first d character match and we have already
// found a prefix match of given length (-1 if no such match)
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

/**
 * Removes the key from the set if the key is present.
 * @param key the key
 * @throws IllegalArgumentException if {@code key} is {@code null}
 */
func (t *TrieST) Delete(key string) {
	t.root = t.delete(t.root, key, 0)
}

func (t *TrieST) delete(x *node, key string, d int) *node {
	if x == nil {
		return nil
	}
	if d == utf8.RuneCountInString(key) {
		if x.value != nil {
			t.n--
		}
		x.value = nil
	} else {
		c := []rune(key)[d]
		x.next[c] = t.delete(x.next[c], key, d+1)
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

func deleteCharAt(prefix []rune, index int) []rune {
	if index < 0 || len(prefix)-1 > index {
		return prefix
	} else {
		return append(prefix[0:index], prefix[index+1:]...)
	}
}
