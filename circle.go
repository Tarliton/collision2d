package sat

import ()

type Circle struct {
	Pos Vector
	R   float64
}

func (circle Circle) getAABB() Polygon {
	r := circle.R
	vector := Vector{r, r}
	corner := circle.Pos.Sub(vector)
	polygon := Box{corner, r * 2, r * 2}.ToPolygon()
	return polygon
}
