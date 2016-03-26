package collision2d

import (
	"math"
)

//PointInCircle returns true if the point is inside the circle.
func PointInCircle(p Vector, c Circle) bool {
	differenceV := Vector{}.Copy(p).Sub(c.Pos)
	radiusSqr := c.R * c.R
	distanceSqr := differenceV.Len2()
	return distanceSqr <= radiusSqr
}

//PointInPolygon returns true if the point is inside a polygon.
func PointInPolygon(p Vector, poly Polygon) bool {
	polygon := NewBox(NewVector(0, 0), 1, 1).ToPolygon()
	polygon.Pos = polygon.Pos.Copy(p)
	isInside, _ := TestPolygonPolygon(polygon, poly)
	return isInside
}

//TestCircleCircle returns true if the circles collide with each other.
func TestCircleCircle(a, b Circle) (isColliding bool, response Response) {
	response = NewResponse()
	differenceV := Vector{}.Copy(b.Pos).Sub(a.Pos)
	totalRadius := a.R + b.R
	totalRadiusSqr := totalRadius * totalRadius
	distanceSqr := differenceV.Len2()

	if distanceSqr > totalRadiusSqr {
		return false, response.NotColliding()
	}
	dist := math.Sqrt(distanceSqr)
	response.A = a
	response.B = b
	response.Overlap = totalRadius - dist
	response.OverlapN = response.OverlapN.Copy(differenceV.Normalize())
	response.OverlapV = response.OverlapV.Copy(differenceV.Normalize()).Scale(response.Overlap)
	response.AInB = a.R <= b.R && dist <= b.R-a.R
	response.BInA = b.R <= a.R && dist <= a.R-b.R

	return true, response
}

//TestPolygonCircle returns true if the polygon collides with the circle.
func TestPolygonCircle(polygon Polygon, circle Circle) (isColliding bool, response Response) {
	response = NewResponse()
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
		changedOverlapN := false
		edge = edge.Copy(polygon.Edges[i])
		point = point.Copy(circlePos).Sub(calcPoints[i])
		if point.Len2() > radius2 {
			response.AInB = false
		}
		region := voronoiRegion(edge, point)

		if region == leftVoronoiRegion {
			edge = edge.Copy(polygon.Edges[prev])
			point2 := Vector{}.Copy(circlePos).Sub(calcPoints[prev])
			region2 := voronoiRegion(edge, point2)
			if region2 == rightVoronoiRegion {
				dist := point.Len()
				if dist > radius {
					return false, response.NotColliding()
				}
				response.BInA = false
				overlapN = overlapN.Copy(point.Normalize())
				changedOverlapN = true
				overlap = radius - dist
			}
		} else if region == rightVoronoiRegion {
			edge = edge.Copy(polygon.Edges[next])
			point = point.Copy(circlePos).Sub(calcPoints[next])
			region2 := voronoiRegion(edge, point)

			if region2 == leftVoronoiRegion {
				dist := point.Len()
				if dist > radius {
					return false, response.NotColliding()
				}
				response.BInA = false
				overlapN = overlapN.Copy(point.Normalize())
				changedOverlapN = true
				overlap = radius - dist
			}
		} else {
			normal := edge.Perp().Normalize()
			dist := point.Dot(normal)
			distAbs := math.Abs(dist)
			if dist > 0 && distAbs > radius {
				return false, response.NotColliding()
			}
			overlapN = overlapN.Copy(normal)
			changedOverlapN = true
			overlap = radius - dist
			if dist >= 0 || overlap < 2*radius {
				response.BInA = false
			}
		}
		if changedOverlapN && math.Abs(overlap) < math.Abs(response.Overlap) {
			response.Overlap = overlap
			response.OverlapN = response.OverlapN.Copy(overlapN)
		}
	}

	response.A = polygon
	response.B = circle
	response.OverlapV = response.OverlapV.Copy(response.OverlapN).Scale(response.Overlap)

	return true, response
}

//TestCirclePolygon returns true if the circle collides with the polygon.
func TestCirclePolygon(circle Circle, polygon Polygon) (isColliding bool, response Response) {
	result, response := TestPolygonCircle(polygon, circle)
	if result {
		a := response.A
		aInB := response.AInB
		response.OverlapN = response.OverlapN.Reverse()
		response.OverlapV = response.OverlapV.Reverse()
		response.A = response.B
		response.B = a
		response.AInB = response.BInA
		response.BInA = aInB
	}
	return result, response
}

//TestPolygonPolygon returns true if the polygons collide with each other.
func TestPolygonPolygon(a, b Polygon) (isColliding bool, response Response) {
	response = NewResponse()
	aPoints := a.CalcPoints
	aLen := len(aPoints)
	bPoints := b.CalcPoints
	bLen := len(bPoints)

	for i := 0; i < aLen; i++ {
		if isSeparatingAxis(a.Pos, b.Pos, aPoints, bPoints, a.Normals[i], &response) {
			return false, response.NotColliding()
		}
	}

	for i := 0; i < bLen; i++ {
		if isSeparatingAxis(a.Pos, b.Pos, aPoints, bPoints, b.Normals[i], &response) {
			return false, response.NotColliding()
		}
	}

	response.A = a
	response.B = b
	response.OverlapV = response.OverlapV.Copy(response.OverlapN).Scale(response.Overlap)

	return true, response
}

func voronoiRegion(line, point Vector) int {
	len2 := line.Len2()
	dp := point.Dot(line)
	if dp < 0 {
		return leftVoronoiRegion
	} else if dp > len2 {
		return rightVoronoiRegion
	} else {
		return middleVoronoiRegion
	}
}

func isSeparatingAxis(aPos, bPos Vector, aPoints, bPoints []Vector, axis Vector, response *Response) bool {
	offsetV := Vector{}.Copy(bPos).Sub(aPos)
	projectedOffset := offsetV.Dot(axis)
	minA, maxA := flattenPointsOn(aPoints, axis)
	minB, maxB := flattenPointsOn(bPoints, axis)
	minB += projectedOffset
	maxB += projectedOffset
	if minA > maxB || minB > maxA {
		return true
	}

	overlap := 0.0
	if minA < minB {
		response.AInB = false
		if maxA < maxB {
			overlap = maxA - minB
			response.BInA = false
		} else {
			option1 := maxA - minB
			option2 := maxB - minA
			if option1 < option2 {
				overlap = option1
			} else {
				overlap = -option2
			}
		}
	} else {
		response.BInA = false
		if maxA > maxB {
			overlap = minA - maxB
			response.AInB = false
		} else {
			option1 := maxA - minB
			option2 := maxB - minA
			if option1 < option2 {
				overlap = option1
			} else {
				overlap = -option2
			}
		}
	}

	absOverlap := math.Abs(overlap)
	if absOverlap < response.Overlap {
		response.Overlap = absOverlap
		response.OverlapN = response.OverlapN.Copy(axis)
		if overlap < 0 {
			response.OverlapN = response.OverlapN.Reverse()
		}
	}
	return false
}

func flattenPointsOn(points []Vector, normal Vector) (min, max float64) {
	min = math.MaxFloat64
	max = -math.MaxFloat64
	length := len(points)
	for i := 0; i < length; i++ {
		dot := points[i].Dot(normal)
		if dot < min {
			min = dot
		}
		if dot > max {
			max = dot
		}
	}
	return min, max
}
