package collision2d_test

import (
	"github.com/Tarliton/collision2d"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCircleString(t *testing.T) {
	circle := collision2d.NewCircle(collision2d.NewVector(10, 10), 5)
	output := string(circle.String())
	assert.Equal(t, "{Pos:{X:10.000000, Y:10.000000}\nRadius: 5.000000}", output, "they should be equal")
}

func TestGetAABBCircle(t *testing.T) {
	circle := collision2d.NewCircle(collision2d.NewVector(10, 10), 5)
	polygonCorners := []float64{
		0, 0,
		10, 0,
		10, 10,
		0, 10,
	}
	polygon := collision2d.NewPolygon(collision2d.NewVector(5, 5), collision2d.NewVector(0, 0), 0, polygonCorners)
	assert.Equal(t, polygon, circle.GetAABB(), "they should be equal")
}

func TestNewCircle(t *testing.T) {
	circle := collision2d.NewCircle(collision2d.NewVector(5, 8), 20)
	assert.Equal(t, float64(5), circle.Pos.X, "they should be equal")
	assert.Equal(t, float64(8), circle.Pos.Y, "they should be equal")
	assert.Equal(t, float64(20), circle.R, "they should be equal")
}
