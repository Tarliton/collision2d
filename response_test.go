package collision2d_test

import (
	"github.com/Tarliton/collision2d"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestResponseString(t *testing.T) {
	response := collision2d.NewResponse()
	expected := string("Response:\n{A: %!s(<nil>)\nB: %!s(<nil>)\nOverlap: 179769313486231570814527423731704356798070567525844996598917476803157260780028538760589558632766878171540458953514382464234321326889464182768467546703537516986049910576551282076245490090389328944075868508455133942304583236903222948165808559332123348274797826204144723168738177180919299881250404026184124858368.000000\nOverlapN: {X:0.000000, Y:0.000000}\nOverlapV: {X:0.000000, Y:0.000000}\nAInB: true, BInA: true}")
	output := string(response.String())
	assert.Equal(t, expected, output, "they should be equal")
}

func TestResponseNotColliding(t *testing.T) {
	response := collision2d.NewResponse()
	notCollidingResponse := response.NotColliding()
	assert.Equal(t, nil, notCollidingResponse.A, "they should be equal")
	assert.Equal(t, nil, notCollidingResponse.B, "they should be equal")
	assert.Equal(t, -math.MaxFloat64, notCollidingResponse.Overlap, "they should be equal")
	assert.Equal(t, collision2d.NewVector(0, 0), notCollidingResponse.OverlapN, "they should be equal")
	assert.Equal(t, collision2d.NewVector(0, 0), notCollidingResponse.OverlapV, "they should be equal")
	assert.Equal(t, false, notCollidingResponse.AInB, "they should be equal")
	assert.Equal(t, false, notCollidingResponse.BInA, "they should be equal")
}
