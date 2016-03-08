package collision2d

import (
	"fmt"
)

//Circle is a struct that represents a circle with a position and a raidus.
type Circle struct {
	Pos Vector
	R   float64
}

func (circle Circle) String() string {
	return fmt.Sprintf("{Pos:%sRadius: %f}", circle.Pos, circle.R)
}

//NewCircle create a new circle with vector pos as center and radius r
func NewCircle(pos Vector, r float64) Circle {
	return Circle{Pos: pos, R: r}
}

//GetAABB returns the axis-aligned bounding box of the circle.
func (circle Circle) GetAABB() Polygon {
	r := circle.R
	vector := NewVector(r, r)
	corner := circle.Pos.Sub(vector)
	polygon := Box{corner, r * 2, r * 2}.ToPolygon()
	return polygon
}
