package collision2d_test

import (
	"github.com/Tarliton/collision2d"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCircleString(t *testing.T) {
	circle := collision2d.Circle{collision2d.Vector{10, 10}, 5}
	output := string(circle.String())
	assert.Equal(t, "{Pos:{X:10.000000, Y:10.000000}\nRadius: 5.000000}", output, "they should be equal")
}

func TestGetAABB(t *testing.T) {
	circle := collision2d.Circle{collision2d.Vector{10, 10}, 5}
	polygon := collision2d.Polygon{collision2d.Vector{5, 5}, collision2d.Vector{}, 0, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}}.SetPoints([]collision2d.Vector{collision2d.Vector{}, collision2d.Vector{10, 0}, collision2d.Vector{10, 10}, collision2d.Vector{0, 10}})
	assert.Equal(t, polygon, circle.GetAABB(), "they should be equal")
}
