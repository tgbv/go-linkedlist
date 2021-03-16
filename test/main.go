package main

import (
	"fmt"

	"github.com/tgbv/go-linkedlist/internal/listTypesSingly"
	"github.com/tgbv/go-linkedlist/pkg"
)

func main() {
	// int8 of length32
	int8 := pkg.MakeSinglyListInt8Length32()
	int8.Add(22).Add(44).Add(100).DeleteAtIndex(0).Traverse(func(n *listTypesSingly.Int8Length32Node) {
		fmt.Println(n.Value)
	})

	// int32 of length32
	int32 := pkg.MakeSinglyListInt32Length32()
	int32.Add(20000).Add(44).Add(100).DeleteAtIndex(2).Traverse(func(n *listTypesSingly.Int32Length32Node) {
		fmt.Println(n.Value)
	})

	// string of length32
	str := pkg.MakeSinglyListStringLength32()
	str.Add("aaa").Add("bbb").Add("ccc").DeleteAtIndex(2).Traverse(func(n *listTypesSingly.StringLength32Node) {
		fmt.Println(n.Value)
	})

}
