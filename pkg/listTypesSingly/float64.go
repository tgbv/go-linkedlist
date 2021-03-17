package listTypesSingly

import (
	"github.com/tgbv/go-linkedlist/internal/listLengths"
	"github.com/tgbv/go-linkedlist/internal/nodeTypesSingly"
)

// type alias
type NodeFloat64 = nodeTypesSingly.NodeFloat64

// actual struct
type Float64Length32 struct {
	listLengths.List32Length
	head *NodeFloat64
	tail *NodeFloat64
}

/*
*	adds a value to the list
 */
func (l *Float64Length32) Add(v float64) *Float64Length32 {
	// first check list length
	if l.Length == 0 {
		l.head = &NodeFloat64{Value: v}
		l.tail = l.head
	} else

	// otherwise assign node to tail.Next
	{
		l.tail.Next = &NodeFloat64{Value: v}
		l.tail = l.tail.Next
	}

	l.Length++

	return l
}

/*
*	retrieve node at index
 */
func (l *Float64Length32) Get(i uint32) *NodeFloat64 {
	// check if index is out of boundaries
	if i < 0 || i > l.Length-1 {
		panic("Index out of boundaries")
	}

	// iterate list until we find the node
	node := l.head
	j := uint32(1) // because there has already been an iteration (node := l.head)

	for j <= i {
		node = node.Next
		j++
	}

	return node
}

/*
*	special function for deleting the head
 */
func (l *Float64Length32) DeleteHead() *Float64Length32 {
	// check list length
	if l.Length == 0 {
		return l
	}

	l.head = l.head.Next

	l.Length--

	return l
}

/*
*	delete node at index
 */
func (l *Float64Length32) DeleteAtIndex(i uint32) *Float64Length32 {
	// check if index is out of boundaries
	if i < 0 || i > l.Length-1 {
		panic("Index out of boundaries")
	}

	// check if it's zero
	// chop head if so
	if i == 0 {
		return l.DeleteHead()
	} else

	// otherwise iterate until index is found
	// tie previous node to the next node
	{
		curr := l.head
		var prev *NodeFloat64 = nil
		j := uint32(1)

		for j <= i {
			prev = curr
			curr = curr.Next
			j++
		}

		prev.Next = curr.Next
	}

	l.Length--

	return l
}

/*
*	traverse the list with a callback
 */
func (l *Float64Length32) Traverse(f func(*NodeFloat64)) *Float64Length32 {
	node := l.head

	for node != nil {
		f(node)
		node = node.Next
	}

	return l
}
