package collision2d

type Box struct {
	Pos  Vector
	W, H float64
}

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
