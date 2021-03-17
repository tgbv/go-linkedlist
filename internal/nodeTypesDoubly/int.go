package nodeTypesDoubly

/*
*	contains the node types for doubly linked lists
*
 */

type NodeInt8 struct {
	Value int8
	Next  *NodeInt8
	Prev  *NodeInt8
}

type NodeInt16 struct {
	Value int16
	Next  *NodeInt16
	Prev  *NodeInt16
}

type NodeInt32 struct {
	Value int32
	Next  *NodeInt32
	Prev  *NodeInt32
}

type NodeInt64 struct {
	Value int64
	Next  *NodeInt64
	Prev  *NodeInt64
}
