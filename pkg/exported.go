package pkg

import "github.com/tgbv/go-linkedlist/internal/listTypesSingly"

/*
*	this one acts as bootstrap for all required modules from within ../internal
 */

func MakeSinglyListInt8Length32() listTypesSingly.Int8Length32 {
	return listTypesSingly.Int8Length32{}
}

func MakeSinglyListInt32Length32() listTypesSingly.Int32Length32 {
	return listTypesSingly.Int32Length32{}
}

func MakeSinglyListStringLength32() listTypesSingly.StringLength32 {
	return listTypesSingly.StringLength32{}
}