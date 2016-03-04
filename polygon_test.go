package collision2d_test

import (
	"github.com/Tarliton/collision2d"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestPolygonString(t *testing.T) {
	polygon := collision2d.Polygon{collision2d.Vector{5, 5}, collision2d.Vector{}, 0, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}}.SetPoints([]collision2d.Vector{collision2d.Vector{}, collision2d.Vector{10, 0}, collision2d.Vector{10, 10}, collision2d.Vector{0, 10}})
	output := string(polygon.String())
	assert.Equal(t, "{Pos:{X:5.000000, Y:5.000000}\nOffset:{X:0.000000, Y:0.000000}\nAngle: 0.000000\nPoints: [{X:0.000000, Y:0.000000}\n {X:10.000000, Y:0.000000}\n {X:10.000000, Y:10.000000}\n {X:0.000000, Y:10.000000}\n]}", output, "they should be equal")
}

func TestSetAngle(t *testing.T) {
	polygon := collision2d.Polygon{collision2d.Vector{5, 5}, collision2d.Vector{}, 0, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}}.SetPoints([]collision2d.Vector{collision2d.Vector{}, collision2d.Vector{10, 0}, collision2d.Vector{10, 10}, collision2d.Vector{0, 10}})
	polygon = polygon.SetAngle(math.Pi / 2)
	assert.Equal(t, math.Pi/2, polygon.Angle, "they should be equal")
}

func TestSetOffset(t *testing.T) {
	polygon := collision2d.Polygon{collision2d.Vector{5, 5}, collision2d.Vector{}, 0, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}}.SetPoints([]collision2d.Vector{collision2d.Vector{}, collision2d.Vector{10, 0}, collision2d.Vector{10, 10}, collision2d.Vector{0, 10}})
	polygon = polygon.SetOffset(collision2d.NewVector(50, 50))
	assert.Equal(t, collision2d.NewVector(50, 50), polygon.Offset, "they should be equal")
}

func TestRotate(t *testing.T) {
	polygon := collision2d.Polygon{collision2d.Vector{5, 5}, collision2d.Vector{}, 0, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}}.SetPoints([]collision2d.Vector{collision2d.Vector{}, collision2d.Vector{10, 0}, collision2d.Vector{10, 10}, collision2d.Vector{0, 10}})
	polygon = polygon.Rotate(math.Pi)
	assert.Equal(t, []collision2d.Vector{collision2d.Vector{X: -0, Y: 0}, collision2d.Vector{X: -10, Y: 1.2246467991473515e-15}, collision2d.Vector{X: -10.000000000000002, Y: -9.999999999999998}, collision2d.Vector{X: -1.2246467991473515e-15, Y: -10}}, polygon.Points, "they should be equal")
}

func TestTranslate(t *testing.T) {
	polygon := collision2d.Polygon{collision2d.Vector{5, 5}, collision2d.Vector{}, 0, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}}.SetPoints([]collision2d.Vector{collision2d.Vector{}, collision2d.Vector{10, 0}, collision2d.Vector{10, 10}, collision2d.Vector{0, 10}})
	polygon = polygon.Translate(5.5, 9.2)
	assert.Equal(t, []collision2d.Vector{collision2d.Vector{X: 5.5, Y: 9.2}, collision2d.Vector{X: 15.5, Y: 9.2}, collision2d.Vector{X: 15.5, Y: 19.2}, collision2d.Vector{X: 5.5, Y: 19.2}}, polygon.Points, "they should be equal")
}

func TestGetAABBPolygonOne(t *testing.T) {
	polygon := collision2d.Polygon{collision2d.Vector{5, 5}, collision2d.Vector{}, 0, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}}.SetPoints([]collision2d.Vector{collision2d.Vector{}, collision2d.Vector{10, 0}, collision2d.Vector{10, 10}, collision2d.Vector{0, 10}})
	polygonAABB := polygon.GetAABB()
	assert.Equal(t, polygon, polygonAABB, "they should be equal")
}

func TestGetAABBPolygonTwo(t *testing.T) {
	polygon := collision2d.Polygon{collision2d.Vector{5, 5}, collision2d.Vector{}, 0, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}, []collision2d.Vector{}}.SetPoints([]collision2d.Vector{collision2d.Vector{10, 10}, collision2d.Vector{0, 10}, collision2d.Vector{0, 0}, collision2d.Vector{10, 0}})
	polygonAABB := polygon.GetAABB()
	assert.Equal(t, collision2d.Polygon{Pos: collision2d.Vector{X: 5, Y: 5}, Offset: collision2d.Vector{X: 0, Y: 0}, Angle: 0, Points: []collision2d.Vector{collision2d.Vector{X: 0, Y: 0}, collision2d.Vector{X: 10, Y: 0}, collision2d.Vector{X: 10, Y: 10}, collision2d.Vector{X: 0, Y: 10}}, CalcPoints: []collision2d.Vector{collision2d.Vector{X: 0, Y: 0}, collision2d.Vector{X: 10, Y: 0}, collision2d.Vector{X: 10, Y: 10}, collision2d.Vector{X: 0, Y: 10}}, Edges: []collision2d.Vector{collision2d.Vector{X: 10, Y: 0}, collision2d.Vector{X: 0, Y: 10}, collision2d.Vector{X: -10, Y: 0}, collision2d.Vector{X: 0, Y: -10}}, Normals: []collision2d.Vector{collision2d.Vector{X: 0, Y: -1}, collision2d.Vector{X: 1, Y: -0}, collision2d.Vector{X: 0, Y: 1}, collision2d.Vector{X: -1, Y: -0}}}, polygonAABB, "they should be equal")
}
