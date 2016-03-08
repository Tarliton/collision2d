package collision2d_test

import (
	"github.com/Tarliton/collision2d"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestPointInCircle(t *testing.T) {
	point := collision2d.NewVector(11, 11)
	circle := collision2d.NewCircle(collision2d.NewVector(10, 10), 5)
	result := collision2d.PointInCircle(point, circle)
	assert.Equal(t, true, result, "they should be equal")
}

func TestPointNotInCircle(t *testing.T) {
	point := collision2d.NewVector(155, 11)
	circle := collision2d.NewCircle(collision2d.NewVector(10, 10), 5)
	result := collision2d.PointInCircle(point, circle)
	assert.Equal(t, false, result, "they should be equal")
}

func TestPointInPolygon(t *testing.T) {
	point := collision2d.NewVector(35, 5)
	polygonCorners := []float64{
		0, 0,
		30, 0,
		0, 30,
	}
	polygon := collision2d.NewPolygon(collision2d.NewVector(30, 0), collision2d.NewVector(0, 0), 0, polygonCorners)
	result := collision2d.PointInPolygon(point, polygon)
	assert.Equal(t, true, result, "they should be equal")
}

func TestPointNotInPolygon(t *testing.T) {
	point := collision2d.NewVector(0, 0)
	polygonCorners := []float64{
		0, 0,
		30, 0,
		0, 30,
	}
	polygon := collision2d.NewPolygon(collision2d.NewVector(30, 0), collision2d.NewVector(0, 0), 0, polygonCorners)
	result := collision2d.PointInPolygon(point, polygon)
	assert.Equal(t, false, result, "they should be equal")
}

func TestTestCircleCircle(t *testing.T) {
	circle1 := collision2d.NewCircle(collision2d.NewVector(0, 0), 20)
	circle2 := collision2d.NewCircle(collision2d.NewVector(30, 0), 20)
	result, response := collision2d.TestCircleCircle(circle1, circle2)
	assert.Equal(t, true, result, "they should be equal")
	assert.Equal(t, float64(10), response.Overlap, "they should be equal")
	assert.Equal(t, true, response.OverlapV.X == float64(10) && response.OverlapV.Y == float64(0), "they should be equal")
}

func TestTestNotCircleCircle(t *testing.T) {
	circle1 := collision2d.NewCircle(collision2d.NewVector(0, 0), 20)
	circle2 := collision2d.NewCircle(collision2d.NewVector(30, 50), 20)
	result, response := collision2d.TestCircleCircle(circle1, circle2)
	assert.Equal(t, false, result, "they should be equal")
	assert.Equal(t, -math.MaxFloat64, response.Overlap, "they should be equal")
	assert.Equal(t, float64(0), response.OverlapV.X, "they should be equal")
	assert.Equal(t, float64(0), response.OverlapV.Y, "they should be equal")
}

func TestTestPolygonCircle(t *testing.T) {
	polygonCorners := []float64{
		0, 0,
		40, 0,
		40, 40,
		0, 40,
	}
	polygon := collision2d.NewPolygon(collision2d.NewVector(0, 0), collision2d.NewVector(0, 0), 0, polygonCorners)
	circle := collision2d.NewCircle(collision2d.NewVector(50, 50), 20)
	result, response := collision2d.TestPolygonCircle(polygon, circle)
	assert.Equal(t, true, result, "they should be equal")
	assert.Equal(t, float64(5.857864376269049), response.Overlap, "they should be equal")
	assert.Equal(t, float64(4.14213562373095), response.OverlapV.X, "they should be equal")
	assert.Equal(t, float64(4.14213562373095), response.OverlapV.Y, "they should be equal")
}

func TestTestNotPolygonCircle(t *testing.T) {
	polygonCorners := []float64{
		0, 0,
		40, 0,
		40, 40,
		0, 40,
	}
	polygon := collision2d.NewPolygon(collision2d.NewVector(0, 0), collision2d.NewVector(0, 0), 0, polygonCorners)
	circle := collision2d.NewCircle(collision2d.NewVector(200, 200), 1)
	result, response := collision2d.TestPolygonCircle(polygon, circle)
	assert.Equal(t, false, result, "they should be equal")
	assert.Equal(t, -math.MaxFloat64, response.Overlap, "they should be equal")
	assert.Equal(t, float64(0), response.OverlapV.X, "they should be equal")
	assert.Equal(t, float64(0), response.OverlapV.Y, "they should be equal")
}

func TestTestCirclePolygon(t *testing.T) {
	polygonCorners := []float64{
		0, 0,
		40, 0,
		40, 40,
		0, 40,
	}
	polygon := collision2d.NewPolygon(collision2d.NewVector(0, 0), collision2d.NewVector(0, 0), 0, polygonCorners)
	circle := collision2d.NewCircle(collision2d.NewVector(50, 50), 20)
	result, response := collision2d.TestCirclePolygon(circle, polygon)
	assert.Equal(t, true, result, "they should be equal")
	assert.Equal(t, float64(5.857864376269049), response.Overlap, "they should be equal")
	assert.Equal(t, float64(-4.14213562373095), response.OverlapV.X, "they should be equal")
	assert.Equal(t, float64(-4.14213562373095), response.OverlapV.Y, "they should be equal")
}

func TestTestPolygonPolygon(t *testing.T) {
	polygonCorners := []float64{
		0, 0,
		40, 0,
		40, 40,
		0, 40,
	}
	polygon := collision2d.NewPolygon(collision2d.NewVector(0, 0), collision2d.NewVector(0, 0), 0, polygonCorners)
	result, response := collision2d.TestPolygonPolygon(polygon, polygon)
	assert.Equal(t, true, result, "they should be equal")
	assert.Equal(t, float64(40), response.Overlap, "they should be equal")
	assert.Equal(t, float64(0), response.OverlapV.X, "they should be equal")
	assert.Equal(t, float64(0), response.OverlapV.Y, "they should be equal")
}
