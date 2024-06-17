package btree

const (
	DEGREE       = 5
	MAX_CHILDREN = 2 * DEGREE       // 10
	MAX_ITEMS    = MAX_CHILDREN - 1 // 9
	MIN_ITEMS    = DEGREE - 1       // 4
)

type item struct {
	key []byte
	val []byte
}

type node struct {
	items [MAX_ITEMS]*item
	child [MAX_CHILDREN]*item

	numItems    int
	numChildren int
}

// isLeaf
func (n *node) isLeaf() bool {
	return n.numChildren == 0
}

// search: search by key, return index

// insertItemAt

// insertChildAt

// split

// insert
