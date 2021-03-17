package main

import (
	"fmt"

	"github.com/tgbv/go-linkedlist/pkg"
	"github.com/tgbv/go-linkedlist/pkg/listTypesDoubly"
	"github.com/tgbv/go-linkedlist/pkg/listTypesSingly"
)

func main() {
	// int8 of length32
	int8 := pkg.MakeSinglyListInt8Length32()
	int8.Add(22).Add(44).Add(100).DeleteAtIndex(0).Traverse(func(n *listTypesSingly.Int8Node) {
		fmt.Println(n.Value)
	})

	// int32 of length32
	int32 := pkg.MakeSinglyListInt32Length32()
	int32.Add(20000).Add(44).Add(100).DeleteAtIndex(2).Traverse(func(n *listTypesSingly.Int32Node) {
		fmt.Println(n.Value)
	})

	// string of length32
	str := pkg.MakeSinglyListStringLength32()
	str.Add("aaa").Add("bbb").Add("ccc").DeleteAtIndex(2).Traverse(func(n *listTypesSingly.StringNode) {
		fmt.Println(n.Value)
	})

	// string of length32
	float := pkg.MakeSinglyListFloat64Length32()
	float.Add(33.0).Add(22.0).Add(12).DeleteAtIndex(2).Traverse(func(n *listTypesSingly.NodeFloat64) {
		fmt.Println(n.Value)
	})

	// doubly list of length 32 containing int32 values
	dInt32 := pkg.MakeDoublyListInt32Length32()
	dInt32.Add(22).Add(11).Add(55).DeleteAtIndex(1).Traverse(func(n *listTypesDoubly.NodeInt32) {
		fmt.Println(n)
	})

}
