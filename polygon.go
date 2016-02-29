package collision2d

type Polygon struct {
	Pos, Offset                        Vector
	Angle                              float64
	Points, CalcPoints, Edges, Normals []Vector
}

func (polygon Polygon) SetPoints(points []Vector) Polygon {
	calcPoints := polygon.CalcPoints
	edges := polygon.Edges
	normals := polygon.Normals
	for i := 0; i < len(points); i++ {
		calcPoints = append(calcPoints, Vector{})
		edges = append(edges, Vector{})
		normals = append(normals, Vector{})
	}
	polygon.Points = points
	polygon.CalcPoints = calcPoints
	polygon.Edges = edges
	polygon.Normals = normals
	polygon.recalc()
	return polygon
}

func (polygon Polygon) SetAngle(angle float64) Polygon {
	polygon.Angle = angle
	polygon.recalc()
	return polygon
}

func (polygon Polygon) SetOffset(offset Vector) Polygon {
	polygon.Offset = offset
	polygon.recalc()
	return polygon
}

func (polygon Polygon) Rotate(angle float64) Polygon {
	points := polygon.Points
	for i := 0; i < len(points); i++ {
		points[i].Rotate(angle)
	}
	polygon.recalc()
	return polygon
}

func (polygon Polygon) Translate(x, y float64) Polygon {
	points := polygon.Points
	for i := 0; i < len(points); i++ {
		points[i].X += x
		points[i].Y += y
	}
	polygon.recalc()
	return polygon
}

func (polygon Polygon) GetAABB() Polygon {
	calcPoints := polygon.CalcPoints
	xMin := calcPoints[0].X
	yMin := calcPoints[0].Y
	xMax := calcPoints[0].X
	yMax := calcPoints[0].Y
	for i := 1; i < len(calcPoints); i++ {
		point := calcPoints[i]
		if point.X < xMin {
			xMin = point.X
		} else if point.X > xMax {
			xMax = point.X
		}

		if point.Y < yMin {
			yMin = point.Y
		} else if point.Y > yMax {
			yMax = point.Y
		}
	}
	box := Box{polygon.Pos.Clone().Add(Vector{xMin, yMin}), xMax - xMin, yMax - yMin}
	return box.ToPolygon()
}

func (polygon *Polygon) recalc() {
	calcPoints := polygon.CalcPoints
	edges := polygon.Edges
	normals := polygon.Normals
	points := polygon.Points
	offset := polygon.Offset
	angle := polygon.Angle
	length := len(points)
	for i := 0; i < length; i++ {
		calcPoint := calcPoints[i].Copy(points[i])
		calcPoint.X += offset.X
		calcPoint.Y += offset.Y
		if angle != 0 {
			calcPoint.Rotate(angle)
		}
		calcPoints[i] = calcPoints[i].Copy(calcPoint)
	}
	for i := 0; i < length; i++ {
		p1 := calcPoints[i]
		var p2 Vector
		if i < length-1 {
			p2 = calcPoints[i+1]
		} else {
			p2 = calcPoints[0]
		}
		edges[i] = edges[i].Copy(p2).Sub(p1)
		normals[i] = normals[i].Copy(edges[i]).Perp().Normalize()
	}
	polygon.CalcPoints = calcPoints
	polygon.Edges = edges
	polygon.Points = points
	polygon.Normals = normals
	polygon.Offset = offset
	polygon.Angle = angle
}
