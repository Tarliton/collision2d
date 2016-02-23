package sat

import ()

type Vector struct {
	x, y float64
}

type Circle struct {
	pos Vector
	r   float64
}

type Polygon struct {
	pos, offset Vector
	angle       float64
	points      []Vector
}

type Box struct {
	pos  Vector
	w, h float64
}
