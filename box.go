package collision2d

import (
	"fmt"
)

//Box is a simple box with position, width and heigth.
type Box struct {
	Pos  Vector
	W, H float64
}

func (box Box) String() string {
	return fmt.Sprintf("{Pos:%sWidth:%f\nHeight:%f}", box.Pos, box.W, box.H)
}

//ToPolygon returns a new polygon whose edges are the edges of the box.
func (box Box) ToPolygon() Polygon {
	pos := box.Pos
	w := box.W
	h := box.H
	vector := Vector{pos.X, pos.Y}
	points := []Vector{}
	points = append(points, Vector{})
	points = append(points, Vector{w, 0})
	points = append(points, Vector{w, h})
	points = append(points, Vector{0, h})
	polygon := Polygon{vector, Vector{}, 0, []Vector{}, []Vector{}, []Vector{}, []Vector{}}.SetPoints(points)
	return polygon
}
