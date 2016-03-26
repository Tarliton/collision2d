package collision2d_test

import (
	"github.com/Tarliton/collision2d"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoxString(t *testing.T) {
	box := collision2d.NewBox(collision2d.NewVector(0, 0), 0, 0)
	output := string(box.String())
	assert.Equal(t, "{Pos:{X:0.000000, Y:0.000000}\nWidth:0.000000\nHeight:0.000000}", output, "Box string should be all zeroed.")
}

func TestBoxNewBox(t *testing.T) {
	box := collision2d.NewBox(collision2d.NewVector(-5, 1), 20, 30)
	assert.Equal(t, float64(-5), box.Pos.X, "Box X position should be -5.")
	assert.Equal(t, float64(1), box.Pos.Y, "Box X position should be 1.")
	assert.Equal(t, float64(20), box.W, "Box width should be 20.")
	assert.Equal(t, float64(30), box.H, "Box height should be 30.")
}

func TestBoxToPolygon(t *testing.T) {
	box := collision2d.NewBox(collision2d.NewVector(50, -20), 15, 20)
	polygon := box.ToPolygon()
	assert.Equal(t, collision2d.Polygon{Pos: collision2d.Vector{X: 50, Y: -20}, Offset: collision2d.Vector{X: 0, Y: 0}, Angle: 0, Points: []collision2d.Vector{collision2d.Vector{X: 0, Y: 0}, collision2d.Vector{X: 15, Y: 0}, collision2d.Vector{X: 15, Y: 20}, collision2d.Vector{X: 0, Y: 20}}, CalcPoints: []collision2d.Vector{collision2d.Vector{X: 0, Y: 0}, collision2d.Vector{X: 15, Y: 0}, collision2d.Vector{X: 15, Y: 20}, collision2d.Vector{X: 0, Y: 20}}, Edges: []collision2d.Vector{collision2d.Vector{X: 15, Y: 0}, collision2d.Vector{X: 0, Y: 20}, collision2d.Vector{X: -15, Y: 0}, collision2d.Vector{X: 0, Y: -20}}, Normals: []collision2d.Vector{collision2d.Vector{X: 0, Y: -1}, collision2d.Vector{X: 1, Y: -0}, collision2d.Vector{X: 0, Y: 1}, collision2d.Vector{X: -1, Y: -0}}}, polygon, "Polygon should be like that.")
}

var polygonBoxTest collision2d.Polygon

func BenchmarkBoxToPolygon(b *testing.B) {
	box := collision2d.NewBox(collision2d.NewVector(50, -20), 15, 20)
	for i := 0; i < b.N; i++ {
		polygonBoxTest = box.ToPolygon()
	}
}
