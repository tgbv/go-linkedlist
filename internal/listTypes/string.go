package listTypes

import (
	"github.com/tgbv/go-linkedlist/internal/listLengths"
	"github.com/tgbv/go-linkedlist/internal/nodeTypesSingly"
)

/*
*	the list type
	formula: data withheld + max Length of nodes
*/
type StringLength32Node = nodeTypesSingly.NodeString
type StringLength32 struct {
	listLengths.List32Length
	head *StringLength32Node
	tail *StringLength32Node
}

/*
*	add int64 value to singly int64 list
	O(1)
*/
func (l *StringLength32) Add(v string) *StringLength32 {
	// check list Length
	// if < 1 then add to start
	if l.Length < 1 {
		l.head = &StringLength32Node{Value: v}
		l.tail = l.head
	} else

	// if > 1 append to tail Next
	{
		l.tail.Next = &StringLength32Node{Value: v}
		l.tail = l.tail.Next
	}

	l.Length++

	return l
}

/*
*	traverse the list with a callback
	O(n)
*/
func (l *StringLength32) Traverse(f func(*StringLength32Node)) *StringLength32 {
	node := l.head

	for node != nil {
		f(node)
		node = node.Next
	}

	return l
}

/*
*	return node at index
	indexes start at 0

	O(n)
*/
func (l *StringLength32) Get(i uint32) *StringLength32Node {
	// check for index overflow
	if i < 0 || i > l.Length-1 {
		panic("Index overflow !")
	}

	node := l.head
	j := uint32(1)

	for j <= i {
		node = node.Next
		j++
	}

	return node
}

/*
*	delete head
	O(1)
*/
func (l *StringLength32) DeleteHead() *StringLength32 {
	// first check Length. there may not be something to delete :)
	if l.Length < 1 {
		return l
	}

	l.head = l.head.Next

	l.Length--

	return l
}

/*
*	delete node at index
	O(n)
*/
func (l *StringLength32) DeleteAtIndex(i uint32) *StringLength32 {
	// check for index overflow
	if i < 0 || i > l.Length-1 {
		panic("Index overflow !")
	}

	// if Length is 0 do nothing
	if l.Length == 0 {
		return l
	}

	// if index is 0 chop the head
	if i == 0 {
		return l.DeleteHead()
	} else

	// otherwise iterate and delete appropriate node
	{
		var (
			prev *StringLength32Node
			curr *StringLength32Node = l.head
			j    uint32              = 1
		)

		for j <= i {
			prev = curr
			curr = prev.Next
			j++
		}

		prev.Next = curr.Next
	}

	l.Length--

	return l
}
