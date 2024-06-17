package btree

import (
	"bytes"
)

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
	items    [MAX_ITEMS]*item
	children [MAX_CHILDREN]*node

	numItems    int
	numChildren int
}

func (n *node) isLeaf() bool {
	return n.numChildren == 0
}

// search: search index of key
func (n *node) search(key []byte) (int, bool) {
	low, high := 0, n.numItems

	for low < high {
		midIndex := (low + high) / 2
		midItem := n.items[midIndex]

		cmp := bytes.Compare(key, midItem.key)

		switch {
		case cmp > 0:
			low = midIndex + 1
		case cmp < 0:
			high = midIndex
		case cmp == 0:
			return midIndex, true
		}
	}

	return low, false
}

// insertItemAt
func (n *node) insertItemAt(pos int, i *item) {
	if pos < n.numItems {
		copy(n.items[pos+1:n.numItems+1], n.items[pos:n.numItems])
	}

	n.items[pos] = i
	n.numItems++
}

// insertChildAt
func (n *node) insertChildAt(pos int, child *node) {
	if pos < n.numChildren {
		copy(n.children[pos+1:n.numChildren+1], n.children[pos:n.numChildren])
	}

	n.children[pos] = child
	n.numChildren++
}

// split: split the node, return new item and new node
func (n *node) split() (*item, *node) {
	midIdx := MIN_ITEMS
	midItem := n.items[midIdx]

	// Create a new node
	// copy half of the items from the current node to the new node.
	newNode := &node{}
	copy(newNode.items[:], n.items[midIdx+1:])
	newNode.numItems = MIN_ITEMS

	// copy child
	if !n.isLeaf() {
		copy(newNode.children[:], n.children[midIdx+1:])
		newNode.numChildren = MIN_ITEMS + 1
	}

	// Remove data items and child pointers from the current node that were moved to the new node.
	for i, l := midIdx, n.numItems; i < l; i++ {
		n.items[i] = nil
		n.numItems--

		if !n.isLeaf() {
			n.children[i+1] = nil
			n.numChildren--
		}
	}

	return midItem, newNode
}

// insert
func (n *node) insert(item *item) bool {
	// 1. search
	pos, found := n.search(item.key)

	// 2. found -> update value of item
	if found {
		n.items[pos] = item
		return false
	}

	// 3. not found -> try to insert
	// isLeaf -> insert to items
    if n.isLeaf() {
        n.insertItemAt(pos, item)
        return true
    }
    
	// need split -> split
    if n.children[pos].numItems >= MAX_ITEMS {
        midItem, newNode := 
    }

	return n.children[pos].insert(item)
}
