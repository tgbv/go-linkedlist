package listTypesDoubly

import (
	"github.com/tgbv/go-linkedlist/internal/listLengths"
	"github.com/tgbv/go-linkedlist/internal/nodeTypesDoubly"
)

/*
*	contains a doubly linked list structure of length uint32 holding int32 values
*
 */

type NodeInt32 = nodeTypesDoubly.NodeInt32
type Int32Length32 struct {
	listLengths.List32Length
	Head *NodeInt32
	Tail *NodeInt32
}

/*
*	adds a node
	O(1)
*/
func (l *Int32Length32) Add(v int32) *Int32Length32 {
	// store ref to a new node
	node := &NodeInt32{Value: v}

	// check list length
	if l.Length == 0 {
		l.Head = node
		l.Tail = l.Head
	} else

	// otherwise assign tail.next->new; new.prev->tail; tail->tail.next
	{
		l.Tail.Next = node
		node.Prev = l.Tail
		l.Tail = l.Tail.Next
	}

	l.Length++

	return l
}

/*
*	get node at index
	O(n/2)
*/
func (l *Int32Length32) Get(i uint32) *NodeInt32 {
	// check index boundaries
	if i < 0 || i > l.Length-1 {
		panic("Index out of boundaries!")
	}

	// for optimization purposes depending on index and list length we'll go forward or backwards list interval
	// so there will always be O(n/2) time complexity, which is better than iterating the whole list - O(n)
	le := l.Length

	if i < le/2 {
		curr := l.Head
		j := uint32(1)

		for j <= i {
			curr = curr.Next
			j++
		}

		return curr
	} else {
		curr := l.Tail
		j := int32(le - 2) /* -1 because indexes start at 0; -1 again because I've already iterated the tail */

		for int32(i) <= j {
			curr = curr.Prev
			j--
		}

		return curr
	}
}

/*
*	chop the head :)
	O(1)
*/
func (l *Int32Length32) DeleteHead() *Int32Length32 {
	// we may not have to delete anything
	if l.Length == 0 {
		return l
	} else {
		l.Head = l.Head.Next
		l.Head.Prev = nil
	}

	l.Length--

	return l
}

/*
*	delete node at index
	O(n/2)
*/
func (l *Int32Length32) DeleteAtIndex(i uint32) *Int32Length32 {
	// check index boundaries
	if i < 0 || i > l.Length-1 {
		panic("Index out of boundaries!")
	}

	if l.Length == 0 {
		return l.DeleteHead()
	} else

	// retrieve the node at index
	// assign prev to next pointer; and next.prev to prev
	{
		node := l.Get(i)

		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}

	l.Length--

	return l
}

/*
*	traverse
	O(n)
*/
func (l *Int32Length32) Traverse(f func(*NodeInt32)) *Int32Length32 {
	// get head and while node isn't null iterate
	node := l.Head

	for node != nil {
		f(node)
		node = node.Next
	}

	return l
}
