package collision2d

import (
	"math"
)

//PointInCircle returns true if the point is inside the circle.
func PointInCircle(point Vector, circle Circle) bool {
	differenceV := NewVector(point.X, point.Y).Sub(circle.Pos)
	radiusSqr := circle.R * circle.R
	distanceSqr := differenceV.Len2()
	return distanceSqr <= radiusSqr
}

//PointInPolygon returns true if the point is inside a polygon.
func PointInPolygon(point Vector, polygon Polygon) bool {
	pointAsPolygon := NewBox(point.Clone(), 1, 1).ToPolygon()
	isInside, _ := TestPolygonPolygon(pointAsPolygon, polygon)
	return isInside
}

//TestCircleCircle returns true if the circles collide with each other.
func TestCircleCircle(circleA, circleB Circle) (isColliding bool, response Response) {
	response = NewResponse()
	differenceV := NewVector(circleB.Pos.X, circleB.Pos.Y).Sub(circleA.Pos)
	totalRadius := circleA.R + circleB.R
	totalRadiusSqr := totalRadius * totalRadius
	distanceSqr := differenceV.Len2()

	if distanceSqr > totalRadiusSqr {
		return false, response.NotColliding()
	}
	dist := math.Sqrt(distanceSqr)
	response.A = circleA
	response.B = circleB
	response.Overlap = totalRadius - dist
	response.OverlapN = response.OverlapN.Copy(differenceV.Normalize())
	response.OverlapV = response.OverlapV.Copy(differenceV.Normalize()).Scale(response.Overlap)
	response.AInB = circleA.R <= circleB.R && dist <= circleB.R-circleA.R
	response.BInA = circleB.R <= circleA.R && dist <= circleA.R-circleB.R

	return true, response
}

//TestPolygonCircle returns true if the polygon collides with the circle.
func TestPolygonCircle(polygon Polygon, circle Circle) (isColliding bool, response Response) {
	response = NewResponse()
	circlePos := NewVector(circle.Pos.X, circle.Pos.Y).Sub(polygon.Pos)
	radius := circle.R
	radius2 := radius * radius
	calcPoints := polygon.CalcPoints
	edge := NewVector(0, 0)
	point := NewVector(0, 0)

	for i := 0; i < len(calcPoints); i++ {
		var next int
		var prev int
		if i == len(calcPoints)-1 {
			next = 0
		} else {
			next = i + 1
		}
		if i == 0 {
			prev = len(calcPoints) - 1
		} else {
			prev = i - 1
		}
		overlap := 0.0
		overlapN := NewVector(0, 0)
		changedOverlapN := false
		edge = edge.Copy(polygon.Edges[i])
		point = point.Copy(circlePos).Sub(calcPoints[i])
		if point.Len2() > radius2 {
			response.AInB = false
		}
		region := voronoiRegion(edge, point)

		if region == leftVoronoiRegion {
			edge = edge.Copy(polygon.Edges[prev])
			point2 := NewVector(circlePos.X, circlePos.Y).Sub(calcPoints[prev])
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
func TestPolygonPolygon(polygonA, polygonB Polygon) (isColliding bool, response Response) {
	response = NewResponse()

	for i := 0; i < len(polygonA.CalcPoints); i++ {
		if isSeparatingAxis(polygonA.Pos, polygonB.Pos, polygonA.CalcPoints, polygonB.CalcPoints, polygonA.Normals[i], &response) {
			return false, response.NotColliding()
		}
	}

	for i := 0; i < len(polygonB.CalcPoints); i++ {
		if isSeparatingAxis(polygonA.Pos, polygonB.Pos, polygonA.CalcPoints, polygonB.CalcPoints, polygonB.Normals[i], &response) {
			return false, response.NotColliding()
		}
	}

	response.A = polygonA
	response.B = polygonB
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
	offsetV := NewVector(bPos.X, bPos.Y).Sub(aPos)
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
