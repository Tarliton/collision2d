package collision2d_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tarliton/collision2d"
	"testing"
)

func TestPointInCircle(t *testing.T) {
	point := collision2d.Vector{11, 11}
	circle := collision2d.Circle{collision2d.Vector{10, 10}, 5}
	result := collision2d.PointInCircle(point, circle)
	assert.Equal(t, true, result, "they should be equal")
}

func TestPointNotInCircle(t *testing.T) {
	point := collision2d.Vector{155, 11}
	circle := collision2d.Circle{collision2d.Vector{10, 10}, 5}
	result := collision2d.PointInCircle(point, circle)
	assert.Equal(t, false, result, "they should be equal")
}

func TestPointInPolygon(t *testing.T) {
	point := collision2d.Vector{35, 5}
	polygon := collision2d.Polygon{collision2d.Vector{30, 0}, collision2d.Vector{}, 0, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}}.SetPoints([]collision2d.Vector{collision2d.Vector{}, collision2d.Vector{30, 0}, collision2d.Vector{0, 30}})
	result := collision2d.PointInPolygon(point, polygon)
	assert.Equal(t, true, result, "they should be equal")
}

func TestPointNotInPolygon(t *testing.T) {
	point := collision2d.Vector{0, 0}
	polygon := collision2d.Polygon{collision2d.Vector{30, 0}, collision2d.Vector{}, 0, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}}.SetPoints([]collision2d.Vector{collision2d.Vector{}, collision2d.Vector{30, 0}, collision2d.Vector{0, 30}})
	result := collision2d.PointInPolygon(point, polygon)
	assert.Equal(t, false, result, "they should be equal")
}

func TestTestCircleCircle(t *testing.T) {
	circle1 := collision2d.Circle{collision2d.Vector{0, 0}, 20}
	circle2 := collision2d.Circle{collision2d.Vector{30, 0}, 20}
	response := new(collision2d.Response)
	result := collision2d.TestCircleCircle(circle1, circle2, response)
	assert.Equal(t, true, result, "they should be equal")
	assert.Equal(t, float64(10), response.Overlap, "they should be equal")
	assert.Equal(t, true, response.OverlapV.X == float64(10) && response.OverlapV.Y == float64(0), "they should be equal")
}

func TestTestPolygonCircle(t *testing.T) {
	polygon := collision2d.Polygon{collision2d.Vector{0, 0}, collision2d.Vector{}, 0, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}}.SetPoints([]collision2d.Vector{collision2d.Vector{}, collision2d.Vector{40, 0}, collision2d.Vector{40, 40}, collision2d.Vector{0, 40}})
	circle := collision2d.Circle{collision2d.Vector{50, 50}, 20}
	response := new(collision2d.Response)
	response.Clear()
	result := collision2d.TestPolygonCircle(polygon, circle, response)
	assert.Equal(t, true, result, "they should be equal")
	assert.Equal(t, float64(5.857864376269049), response.Overlap, "they should be equal")
	assert.Equal(t, float64(4.14213562373095), response.OverlapV.X, "they should be equal")
	assert.Equal(t, float64(4.14213562373095), response.OverlapV.Y, "they should be equal")
}

func TestTestCirclePolygon(t *testing.T) {
	polygon := collision2d.Polygon{collision2d.Vector{0, 0}, collision2d.Vector{}, 0, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}}.SetPoints([]collision2d.Vector{collision2d.Vector{}, collision2d.Vector{40, 0}, collision2d.Vector{40, 40}, collision2d.Vector{0, 40}})
	circle := collision2d.Circle{collision2d.Vector{50, 50}, 20}
	response := new(collision2d.Response)
	response.Clear()
	result := collision2d.TestCirclePolygon(circle, polygon, response)
	assert.Equal(t, true, result, "they should be equal")
	assert.Equal(t, float64(5.857864376269049), response.Overlap, "they should be equal")
	assert.Equal(t, float64(-4.14213562373095), response.OverlapV.X, "they should be equal")
	assert.Equal(t, float64(-4.14213562373095), response.OverlapV.Y, "they should be equal")
}

func TestTestPolygonPolygon(t *testing.T) {
	polygon := collision2d.Polygon{collision2d.Vector{0, 0}, collision2d.Vector{}, 0, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}}.SetPoints([]collision2d.Vector{collision2d.Vector{}, collision2d.Vector{40, 0}, collision2d.Vector{40, 40}, collision2d.Vector{0, 40}})
	response := new(collision2d.Response)
	response.Clear()
	result := collision2d.TestPolygonPolygon(polygon, polygon, response)
	assert.Equal(t, true, result, "they should be equal")
	assert.Equal(t, float64(40), response.Overlap, "they should be equal")
	assert.Equal(t, float64(0), response.OverlapV.X, "they should be equal")
	assert.Equal(t, float64(0), response.OverlapV.Y, "they should be equal")
}
