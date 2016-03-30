package collision2d_test

import (
	"github.com/Tarliton/collision2d"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestCollisionPointInCircle(t *testing.T) {
	point := collision2d.NewVector(11, 11)
	circle := collision2d.NewCircle(collision2d.NewVector(10, 10), 5)
	result := collision2d.PointInCircle(point, circle)
	assert.Equal(t, true, result, "they should be equal")
}

func TestCollisionPointNotInCircle(t *testing.T) {
	point := collision2d.NewVector(155, 11)
	circle := collision2d.NewCircle(collision2d.NewVector(10, 10), 5)
	result := collision2d.PointInCircle(point, circle)
	assert.Equal(t, false, result, "they should be equal")
}

func TestCollisionPointInPolygon(t *testing.T) {
	point := collision2d.NewVector(35, 5)
	polygonCorners := [...]float64{
		0, 0,
		30, 0,
		0, 30,
	}
	polygon := collision2d.NewPolygon(collision2d.NewVector(30, 0), collision2d.NewVector(0, 0), 0, polygonCorners[:])
	result := collision2d.PointInPolygon(point, polygon)
	assert.Equal(t, true, result, "they should be equal")
}

func TestCollisionPointNotInPolygon(t *testing.T) {
	point := collision2d.NewVector(0, 0)
	polygonCorners := [...]float64{
		0, 0,
		30, 0,
		0, 30,
	}
	polygon := collision2d.NewPolygon(collision2d.NewVector(30, 0), collision2d.NewVector(0, 0), 0, polygonCorners[:])
	result := collision2d.PointInPolygon(point, polygon)
	assert.Equal(t, false, result, "they should be equal")
}

func TestCollisionTestCircleCircle(t *testing.T) {
	circle1 := collision2d.NewCircle(collision2d.NewVector(0, 0), 20)
	circle2 := collision2d.NewCircle(collision2d.NewVector(30, 0), 20)
	result, response := collision2d.TestCircleCircle(circle1, circle2)
	assert.Equal(t, true, result, "they should be equal")
	assert.Equal(t, float64(10), response.Overlap, "they should be equal")
	assert.Equal(t, true, response.OverlapV.X == float64(10) && response.OverlapV.Y == float64(0), "they should be equal")
}

func TestCollisionTestNotCircleCircle(t *testing.T) {
	circle1 := collision2d.NewCircle(collision2d.NewVector(0, 0), 20)
	circle2 := collision2d.NewCircle(collision2d.NewVector(30, 50), 20)
	result, response := collision2d.TestCircleCircle(circle1, circle2)
	assert.Equal(t, false, result, "they should be equal")
	assert.Equal(t, -math.MaxFloat64, response.Overlap, "they should be equal")
	assert.Equal(t, float64(0), response.OverlapV.X, "they should be equal")
	assert.Equal(t, float64(0), response.OverlapV.Y, "they should be equal")
}

func TestCollisionTestPolygonCircle(t *testing.T) {
	polygonCorners := [...]float64{
		0, 0,
		40, 0,
		40, 40,
		0, 40,
	}
	polygon := collision2d.NewPolygon(collision2d.NewVector(0, 0), collision2d.NewVector(0, 0), 0, polygonCorners[:])
	circle := collision2d.NewCircle(collision2d.NewVector(50, 50), 20)
	result, response := collision2d.TestPolygonCircle(polygon, circle)
	assert.Equal(t, true, result, "they should be equal")
	assert.Equal(t, float64(5.857864376269049), response.Overlap, "they should be equal")
	assert.Equal(t, float64(4.14213562373095), response.OverlapV.X, "they should be equal")
	assert.Equal(t, float64(4.14213562373095), response.OverlapV.Y, "they should be equal")
}

func TestCollisionTestNotPolygonCircle(t *testing.T) {
	polygonCorners := [...]float64{
		0, 0,
		40, 0,
		40, 40,
		0, 40,
	}
	polygon := collision2d.NewPolygon(collision2d.NewVector(0, 0), collision2d.NewVector(0, 0), 0, polygonCorners[:])
	circle := collision2d.NewCircle(collision2d.NewVector(200, 200), 1)
	result, response := collision2d.TestPolygonCircle(polygon, circle)
	assert.Equal(t, false, result, "they should be equal")
	assert.Equal(t, -math.MaxFloat64, response.Overlap, "they should be equal")
	assert.Equal(t, float64(0), response.OverlapV.X, "they should be equal")
	assert.Equal(t, float64(0), response.OverlapV.Y, "they should be equal")
}

func TestCollisionTestCirclePolygon(t *testing.T) {
	polygonCorners := [...]float64{
		0, 0,
		40, 0,
		40, 40,
		0, 40,
	}
	polygon := collision2d.NewPolygon(collision2d.NewVector(0, 0), collision2d.NewVector(0, 0), 0, polygonCorners[:])
	circle := collision2d.NewCircle(collision2d.NewVector(50, 50), 20)
	result, response := collision2d.TestCirclePolygon(circle, polygon)
	assert.Equal(t, true, result, "they should be equal")
	assert.Equal(t, float64(5.857864376269049), response.Overlap, "they should be equal")
	assert.Equal(t, float64(-4.14213562373095), response.OverlapV.X, "they should be equal")
	assert.Equal(t, float64(-4.14213562373095), response.OverlapV.Y, "they should be equal")
}

func TestCollisionTestPolygonPolygon(t *testing.T) {
	polygonCorners := [...]float64{
		0, 0,
		40, 0,
		40, 40,
		0, 40,
	}
	polygon := collision2d.NewPolygon(collision2d.NewVector(0, 0), collision2d.NewVector(0, 0), 0, polygonCorners[:])
	result, response := collision2d.TestPolygonPolygon(polygon, polygon)
	assert.Equal(t, true, result, "they should be equal")
	assert.Equal(t, float64(40), response.Overlap, "they should be equal")
	assert.Equal(t, float64(0), response.OverlapV.X, "they should be equal")
	assert.Equal(t, float64(40), response.OverlapV.Y, "they should be equal")
}

var response collision2d.Response

func BenchmarkCollisionPointInCircle(b *testing.B) {
	circle := collision2d.NewCircle(collision2d.NewVector(0, 0), 20)
	point := collision2d.NewVector(1, 2)
	for i := 0; i < b.N; i++ {
		collision2d.PointInCircle(point, circle)
	}
}

func BenchmarkCollisionPointInPolygon(b *testing.B) {
	polygonCorners := [...]float64{
		0, 0,
		40, 0,
		40, 40,
		0, 40,
	}
	polygon := collision2d.NewPolygon(collision2d.NewVector(0, 0), collision2d.NewVector(0, 0), 0, polygonCorners[:])
	point := collision2d.NewVector(1, 2)
	for i := 0; i < b.N; i++ {
		collision2d.PointInPolygon(point, polygon)
	}
}

func BenchmarkCollisionTestCircleCircle(b *testing.B) {
	circle1 := collision2d.NewCircle(collision2d.NewVector(0, 0), 20)
	circle2 := collision2d.NewCircle(collision2d.NewVector(30, 0), 20)
	for i := 0; i < b.N; i++ {
		_, response = collision2d.TestCircleCircle(circle1, circle2)
	}
}

func BenchmarkCollisionTestPolygonCircle(b *testing.B) {
	circle := collision2d.NewCircle(collision2d.NewVector(0, 0), 20)
	polygonCorners := [...]float64{
		0, 0,
		40, 0,
		40, 40,
		0, 40,
	}
	polygon := collision2d.NewPolygon(collision2d.NewVector(0, 0), collision2d.NewVector(0, 0), 0, polygonCorners[:])
	for i := 0; i < b.N; i++ {
		_, response = collision2d.TestPolygonCircle(polygon, circle)
	}
}

func BenchmarkCollisionTestCirclePolygon(b *testing.B) {
	circle := collision2d.NewCircle(collision2d.NewVector(0, 0), 20)
	polygonCorners := [...]float64{
		0, 0,
		40, 0,
		40, 40,
		0, 40,
	}
	polygon := collision2d.NewPolygon(collision2d.NewVector(0, 0), collision2d.NewVector(0, 0), 0, polygonCorners[:])
	for i := 0; i < b.N; i++ {
		_, response = collision2d.TestCirclePolygon(circle, polygon)
	}
}

func BenchmarkCollisionTestPolygonPolygon(b *testing.B) {
	polygonCorners1 := [...]float64{
		0, 0,
		40, 0,
		40, 40,
		0, 40,
	}
	polygon1 := collision2d.NewPolygon(collision2d.NewVector(0, 0), collision2d.NewVector(0, 0), 0, polygonCorners1[:])
	polygonCorners2 := [...]float64{
		0, 0,
		40, 0,
		40, 40,
		0, 40,
	}
	polygon2 := collision2d.NewPolygon(collision2d.NewVector(15, 25), collision2d.NewVector(0, 0), 0, polygonCorners2[:])
	for i := 0; i < b.N; i++ {
		_, response = collision2d.TestPolygonPolygon(polygon1, polygon2)
	}
}
