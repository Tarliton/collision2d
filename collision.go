package sat

import (
	"math"
)

func PointInCircle(p Vector, c Circle) bool {
	differenceV := Vector{}.Copy(p).Sub(c.Pos)
	radiusSqr := c.R * c.R
	distanceSqr := differenceV.Len2()
	return distanceSqr <= radiusSqr
}

func PointInPolygon(p Vector, poly Polygon) bool {
	polygon := Box{Vector{}, 1, 1}.ToPolygon()
	polygon.Pos.Copy(p)
	response := new(Response)
	result := TestPolygonPolygon(polygon, poly, response)
	return result
}

func TestCircleCircle(a, b Circle, response *Response) bool {
	differenceV := Vector{}.Copy(b.Pos).Sub(a.Pos)
	totalRadius := a.R + b.R
	totalRadiusSqr := totalRadius * totalRadius
	distanceSqr := differenceV.Len2()

	if distanceSqr > totalRadiusSqr {
		return false
	}
	if response != nil {
		dist := math.Sqrt(distanceSqr)
		response.A = a
		response.B = b
		response.Overlap = totalRadius - dist
		response.overlapN.Copy(differenceV.Normalize())
		response.overlapV.Copy(differenceV).Scale(response.Overlap)
		response.AInB = a.R <= b.R && dist <= b.R-a.R
		response.BInA = b.R <= a.R && dist <= a.R-b.R
	}
	return true
}

func TestPolygonCircle(polygon Polygon, circle Circle, response *Response) bool {
	circlePos := Vector{}.Copy(circle.Pos).Sub(polygon.Pos)
	radius := circle.R
	radius2 := radius * radius
	calcPoints := polygon.CalcPoints
	length := len(calcPoints)
	edge := Vector{}
	point := Vector{}

	for i := 0; i < length; i++ {
		var next int
		var prev int
		if i == length-1 {
			next = 0
		} else {
			next = i + 1
		}
		if i == 0 {
			prev = length - 1
		} else {
			prev = i - 1
		}
		overlap := 0.0
		overlapN := Vector{}
		edge.Copy(polygon.Edges[i])
		point.Copy(circlePos).Sub(calcPoints[i])
		if response != nil && point.Len2() > radius2 {
			response.AInB = false
		}
		region := voronoiRegion(edge, point)

		if region == LEFT_VORONOI_REGION {
			edge.Copy(polygon.Edges[prev])
			point2 := Vector{}.Copy(circlePos).Sub(calcPoints[prev])
			region := voronoiRegion(edge, point2)
			if region == RIGHT_VORONOI_REGION {
				dist := point.Len()
				if dist > radius {
					return false
				} else if response != nil {
					response.BInA = false
					overlapN.Copy(point.Normalize())
					overlap = radius - dist
				}
			}
		} else if region == RIGHT_VORONOI_REGION {
			edge.Copy(polygon.Edges[next])
			point.Copy(circlePos).Sub(calcPoints[next])
			region = voronoiRegion(edge, point)

			if region == LEFT_VORONOI_REGION {
				dist := point.Len()
				if dist > radius {
					return false
				} else if response != nil {
					response.BInA = false
					overlapN.Copy(point.Normalize())
					overlap = radius - dist
				}
			}
		} else {
			normal := edge.Perp().Normalize()
			dist := point.Dot(normal)
			distAbs := math.Abs(dist)
			if dist > 0 && distAbs > radius {
				return false
			} else if response != nil {
				overlapN.Copy(normal)
				overlap = radius - dist
				if dist >= 0 || overlap < 2*radius {
					response.BInA = false
				}
			}
		}

		if response != nil && math.Abs(overlap) < math.Abs(response.Overlap) {
			response.Overlap = overlap
			response.overlapN.Copy(overlapN)
		}
	}

	if response != nil {
		response.A = polygon
		response.B = circle
		response.overlapV.Copy(response.overlapN).Scale(response.Overlap)
	}
	return true
}

func TestCirclePolygon(circle Circle, polygon Polygon, response *Response) bool {
	result := TestPolygonCircle(polygon, circle, response)
	if result && response != nil {
		a := response.A
		aInB := response.AInB
		response.overlapN.Reverse()
		response.overlapV.Reverse()
		response.A = response.B
		response.B = a
		response.AInB = response.BInA
		response.BInA = aInB
	}
	return result
}

func TestPolygonPolygon(a, b Polygon, response *Response) bool {
	aPoints := a.CalcPoints
	aLen := len(aPoints)
	bPoints := b.CalcPoints
	bLen := len(bPoints)

	for i := 0; i < aLen; i++ {
		if isSeparatingAxis(a.Pos, b.Pos, aPoints, bPoints, a.Normals[i], response) {
			return false
		}
	}

	for i := 0; i < bLen; i++ {
		if isSeparatingAxis(a.Pos, b.Pos, aPoints, bPoints, b.Normals[i], response) {
			return false
		}
	}

	if response != nil {
		response.A = a
		response.B = b
		response.overlapV.Copy(response.overlapN).Scale(response.Overlap)
	}
	return true
}

func voronoiRegion(line, point Vector) int {
	len2 := line.Len2()
	dp := point.Dot(line)
	if dp < 0 {
		return LEFT_VORONOI_REGION
	} else if dp > len2 {
		return RIGHT_VORONOI_REGION
	} else {
		return MIDDLE_VORONOI_REGION
	}
}
