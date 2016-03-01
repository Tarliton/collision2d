package collision2d

import (
	"math"
)

//Response contains the information about an collision test.
type Response struct {
	A, B               interface{}
	Overlap            float64
	OverlapN, OverlapV Vector
	AInB, BInA         bool
}

//NewResponse is used to create a new response when necessary.
func NewResponse() *Response {
	return &Response{Overlap: math.MaxFloat64, AInB: true, BInA: true}
}
