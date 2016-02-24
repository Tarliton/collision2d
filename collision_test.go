package sat_test

import (
	"github.com/stretchr/testify/assert"
	sat "github.com/tarliton/sat-go"
	"testing"
)

func TestPointInCircle(t *testing.T) {
	point := sat.Vector{11, 11}
	circle := sat.Circle{sat.Vector{10, 10}, 5}
	result := sat.PointInCircle(point, circle)
	assert.Equal(t, true, result, "they should be equal")
}

func TestPointNotInCircle(t *testing.T) {
	point := sat.Vector{155, 11}
	circle := sat.Circle{sat.Vector{10, 10}, 5}
	result := sat.PointInCircle(point, circle)
	assert.Equal(t, false, result, "they should be equal")
}

func TestPointInPolygon(t *testing.T) {
	point := sat.Vector{35, 5}
	polygon := sat.Polygon{sat.Vector{30, 0}, sat.Vector{}, 0, []sat.Vector{}, []sat.Vector{}, []sat.Vector{}, []sat.Vector{}}.SetPoints([]sat.Vector{sat.Vector{}, sat.Vector{30, 0}, sat.Vector{0, 30}})
	result := sat.PointInPolygon(point, polygon)
	assert.Equal(t, true, result, "they should be equal")
}

func TestPointNotInPolygon(t *testing.T) {
	point := sat.Vector{0, 0}
	polygon := sat.Polygon{sat.Vector{30, 0}, sat.Vector{}, 0, []sat.Vector{}, []sat.Vector{}, []sat.Vector{}, []sat.Vector{}}.SetPoints([]sat.Vector{sat.Vector{}, sat.Vector{30, 0}, sat.Vector{0, 30}})
	result := sat.PointInPolygon(point, polygon)
	assert.Equal(t, false, result, "they should be equal")
}

func TestTestCircleCircle(t *testing.T) {
	circle1 := sat.Circle{sat.Vector{0, 0}, 20}
	circle2 := sat.Circle{sat.Vector{30, 0}, 20}
	response := new(sat.Response)
	result := sat.TestCircleCircle(circle1, circle2, response)
	assert.Equal(t, true, result, "they should be equal")
	assert.Equal(t, float64(10), response.Overlap, "they should be equal")
	assert.Equal(t, true, response.OverlapV.X == float64(10) && response.OverlapV.Y == float64(0), "they should be equal")
}

func TestTestPolygonCircle(t *testing.T) {
	polygon := sat.Polygon{sat.Vector{0, 0}, sat.Vector{}, 0, []sat.Vector{}, []sat.Vector{}, []sat.Vector{}, []sat.Vector{}}.SetPoints([]sat.Vector{sat.Vector{}, sat.Vector{40, 0}, sat.Vector{40, 40}, sat.Vector{0, 40}})
	circle := sat.Circle{sat.Vector{50, 50}, 20}
	response := new(sat.Response)
	response.Clear()
	result := sat.TestPolygonCircle(polygon, circle, response)
	assert.Equal(t, true, result, "they should be equal")
	assert.Equal(t, float64(5.857864376269049), response.Overlap, "they should be equal")
	assert.Equal(t, float64(4.14213562373095), response.OverlapV.X, "they should be equal")
	assert.Equal(t, float64(4.14213562373095), response.OverlapV.Y, "they should be equal")
}

func TestTestCirclePolygon(t *testing.T) {
	polygon := sat.Polygon{sat.Vector{0, 0}, sat.Vector{}, 0, []sat.Vector{}, []sat.Vector{}, []sat.Vector{}, []sat.Vector{}}.SetPoints([]sat.Vector{sat.Vector{}, sat.Vector{40, 0}, sat.Vector{40, 40}, sat.Vector{0, 40}})
	circle := sat.Circle{sat.Vector{50, 50}, 20}
	response := new(sat.Response)
	response.Clear()
	result := sat.TestCirclePolygon(circle, polygon, response)
	assert.Equal(t, true, result, "they should be equal")
	assert.Equal(t, float64(5.857864376269049), response.Overlap, "they should be equal")
	assert.Equal(t, float64(-4.14213562373095), response.OverlapV.X, "they should be equal")
	assert.Equal(t, float64(-4.14213562373095), response.OverlapV.Y, "they should be equal")
}

func TestTestPolygonPolygon(t *testing.T) {
	polygon := sat.Polygon{sat.Vector{0, 0}, sat.Vector{}, 0, []sat.Vector{}, []sat.Vector{}, []sat.Vector{}, []sat.Vector{}}.SetPoints([]sat.Vector{sat.Vector{}, sat.Vector{40, 0}, sat.Vector{40, 40}, sat.Vector{0, 40}})
	response := new(sat.Response)
	response.Clear()
	result := sat.TestPolygonPolygon(polygon, polygon, response)
	assert.Equal(t, true, result, "they should be equal")
	assert.Equal(t, float64(40), response.Overlap, "they should be equal")
	assert.Equal(t, float64(0), response.OverlapV.X, "they should be equal")
	assert.Equal(t, float64(0), response.OverlapV.Y, "they should be equal")
}
