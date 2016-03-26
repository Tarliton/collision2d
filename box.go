package collision2d

import (
	"fmt"
)

//Box is a simple box with position, width and height.
type Box struct {
	Pos  Vector
	W, H float64
}

func (box Box) String() string {
	return fmt.Sprintf("{Pos:%sWidth:%f\nHeight:%f}", box.Pos, box.W, box.H)
}

//NewBox create a new box with vector pos as center and width w and height h
func NewBox(pos Vector, w, h float64) Box {
	return Box{Pos: pos, W: w, H: h}
}

//ToPolygon returns a new polygon whose edges are the edges of the box.
func (box Box) ToPolygon() Polygon {
	polygonCorners := [...]float64{
		0, 0,
		box.W, 0,
		box.W, box.H,
		0, box.H,
	}
	return NewPolygon(box.Pos.Clone(), NewVector(0, 0), 0, polygonCorners[:])
}
