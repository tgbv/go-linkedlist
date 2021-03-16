package nodeTypesSingly

/*
*	contains all base node types
*
*	each node is categorized based on the Value it can hold
*
 */

/*
*	base node holding bool Values
 */
type NodeBool struct {
	Value bool
	Next  *NodeBool
}

/*
*	base node holding int8 Values
 */
type NodeInt8 struct {
	Value int8
	Next  *NodeInt8
}

/*
*	base node holding int16 Values
 */
type NodeInt16 struct {
	Value int16
	Next  *NodeInt16
}

/*
*	base node holding int32 Values
 */
type NodeInt32 struct {
	Value int32
	Next  *NodeInt32
}

/*
*	base node holding int64 Values
 */
type NodeInt64 struct {
	Value int64
	Next  *NodeInt64
}

/*
*	base node holding float32 Values
 */
type NodeFloat32 struct {
	Value float32
	Next  *NodeFloat32
}

/*
*	base node holding float64 Values
 */
type NodeFloat64 struct {
	Value float64
	Next  *NodeFloat64
}

/*
*	base node holding complex64 Values
 */
type NodeComplex64 struct {
	Value complex64
	Next  *NodeComplex64
}

/*
*	base node holding complex128 Values
 */
type NodeComplex128 struct {
	Value complex128
	Next  *NodeComplex128
}

/*
*	base node holding string Values
 */
type NodeString struct {
	Value string
	Next  *NodeString
}

/*
*	base node holding struct Values
 */
type NodeStruct struct {
	Value struct{}
	Next  *NodeStruct
}
