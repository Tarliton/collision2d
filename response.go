package sat

import (
	"math"
)

type Response struct {
	A, B               interface{}
	Overlap            float64
	OverlapN, OverlapV Vector
	AInB, BInA         bool
}

func (response *Response) Clear() {
	response.AInB = true
	response.BInA = true
	response.Overlap = math.MaxFloat64
}
