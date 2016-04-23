package collision2d

import (
	"fmt"
)

//Polygon struct represents a polygon with position and edges in a counter-clockwise fashion.
type Polygon struct {
	Pos, Offset                        Vector
	Angle                              float64
	Points, CalcPoints, Edges, Normals []Vector
}

func (polygon Polygon) String() string {
	return fmt.Sprintf("{Pos:%sOffset:%sAngle: %f\nPoints: %s\nCalcPoints: %s}", polygon.Pos, polygon.Offset, polygon.Angle, polygon.Points, polygon.CalcPoints)
}

//NewPolygon creates a new polygon with pos, offset, angle and points.
//Points is an array of pairs of float64 values, that are mapped into Vectors with X and Y.
//The first value is X and the second is Y. See test to understand better.
func NewPolygon(pos, offset Vector, angle float64, points []float64) Polygon {
	var vectorPoints = make([]Vector, len(points)/2)
	for i := 0; i < len(points); i += 2 {
		vectorPoints[i/2] = NewVector(points[i], points[i+1])
	}
	polygon := Polygon{Pos: pos, Offset: offset, Angle: angle}
	return polygon.SetPoints(vectorPoints)
}

//SetPoints change the edges of the polygon and recauculate the rest of it's values.
func (polygon Polygon) SetPoints(points []Vector) Polygon {
	polygon.CalcPoints = make([]Vector, len(points))
	polygon.Edges = make([]Vector, len(points))
	polygon.Normals = make([]Vector, len(points))
	polygon.Points = make([]Vector, len(points))
	for i := 0; i < len(points); i++ {
		polygon.Points[i] = polygon.Points[i].Copy(points[i])
	}
	polygon.recalc()
	return polygon
}

//SetAngle changes the angle of the polygon
func (polygon Polygon) SetAngle(angle float64) Polygon {
	polygon.Angle = angle
	polygon.recalc()
	return polygon
}

//SetOffset changes the offset of the polygon
func (polygon Polygon) SetOffset(offset Vector) Polygon {
	polygon.Offset = offset
	polygon.recalc()
	return polygon
}

//Rotate rotates the polygon by angle in radian.
func (polygon Polygon) Rotate(angle float64) Polygon {
	points := polygon.Points
	for i := 0; i < len(points); i++ {
		points[i] = points[i].Rotate(angle)
	}
	polygon.recalc()
	return polygon
}

//Translate the polygon by x and y.
func (polygon Polygon) Translate(x, y float64) Polygon {
	points := polygon.Points
	for i := 0; i < len(points); i++ {
		points[i].X += x
		points[i].Y += y
	}
	polygon.recalc()
	return polygon
}

//GetAABB returns the axis-aligned bounding box of the polygon.
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
	box := NewBox(polygon.Pos.Clone().Add(NewVector(xMin, yMin)), xMax-xMin, yMax-yMin)
	return box.ToPolygon()
}

func (polygon *Polygon) recalc() {
	for i := 0; i < len(polygon.Points); i++ {
		polygon.CalcPoints[i] = polygon.CalcPoints[i].Copy(polygon.Points[i])
		polygon.CalcPoints[i].X += polygon.Offset.X
		polygon.CalcPoints[i].Y += polygon.Offset.Y
		if polygon.Angle != 0 {
			polygon.CalcPoints[i] = polygon.CalcPoints[i].Rotate(polygon.Angle)
		}
	}
	for i := 0; i < len(polygon.Points); i++ {
		var p2 Vector
		if i < len(polygon.Points)-1 {
			p2 = polygon.CalcPoints[i+1]
		} else {
			p2 = polygon.CalcPoints[0]
		}
		polygon.Edges[i] = polygon.Edges[i].Copy(p2).Sub(polygon.CalcPoints[i])
		polygon.Normals[i] = polygon.Normals[i].Copy(polygon.Edges[i]).Perp().Normalize()
	}
}
