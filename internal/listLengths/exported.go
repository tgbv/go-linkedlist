package listLengths

/*
*	contains list Length types
*
*
*
 */

// can't have negative Length
// all ought to be unsigned
type List8Length struct {
	Length uint8
}

type List16Length struct {
	Length uint16
}

type List32Length struct {
	Length uint32
}

type List64Length struct {
	Length uint64
}

// f = float
type List32fLength struct {
	Length float32
}

type List64fLength struct {
	Length float64
}

type List64cLength struct {
	Length complex64
}

// c = complex
type List128cLength struct {
	Length complex128
}
