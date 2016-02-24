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
	polygon.Pos = polygon.Pos.Copy(p)
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
		response.OverlapN = response.OverlapN.Copy(differenceV.Normalize())
		response.OverlapV = response.OverlapV.Copy(differenceV.Normalize()).Scale(response.Overlap)
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
		changedOverlapN := false
		edge = edge.Copy(polygon.Edges[i])
		point = point.Copy(circlePos).Sub(calcPoints[i])
		if response != nil && point.Len2() > radius2 {
			response.AInB = false
		}
		region := voronoiRegion(edge, point)

		if region == LEFT_VORONOI_REGION {
			edge = edge.Copy(polygon.Edges[prev])
			point2 := Vector{}.Copy(circlePos).Sub(calcPoints[prev])
			region2 := voronoiRegion(edge, point2)
			if region2 == RIGHT_VORONOI_REGION {
				dist := point.Len()
				if dist > radius {
					return false
				} else if response != nil {
					response.BInA = false
					overlapN = overlapN.Copy(point.Normalize())
					changedOverlapN = true
					overlap = radius - dist
				}
			}
		} else if region == RIGHT_VORONOI_REGION {
			edge = edge.Copy(polygon.Edges[next])
			point = point.Copy(circlePos).Sub(calcPoints[next])
			region2 := voronoiRegion(edge, point)

			if region2 == LEFT_VORONOI_REGION {
				dist := point.Len()
				if dist > radius {
					return false
				} else if response != nil {
					response.BInA = false
					overlapN = overlapN.Copy(point.Normalize())
					changedOverlapN = true
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
				overlapN = overlapN.Copy(normal)
				changedOverlapN = true
				overlap = radius - dist
				if dist >= 0 || overlap < 2*radius {
					response.BInA = false
				}
			}
		}
		if changedOverlapN && response != nil && math.Abs(overlap) < math.Abs(response.Overlap) {
			response.Overlap = overlap
			response.OverlapN = response.OverlapN.Copy(overlapN)
		}
	}

	if response != nil {
		response.A = polygon
		response.B = circle
		response.OverlapV = response.OverlapV.Copy(response.OverlapN).Scale(response.Overlap)
	}
	return true
}

func TestCirclePolygon(circle Circle, polygon Polygon, response *Response) bool {
	result := TestPolygonCircle(polygon, circle, response)
	if result && response != nil {
		a := response.A
		aInB := response.AInB
		response.OverlapN = response.OverlapN.Reverse()
		response.OverlapV = response.OverlapV.Reverse()
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
		response.OverlapV.Copy(response.OverlapN).Scale(response.Overlap)
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

func isSeparatingAxis(aPos, bPos Vector, aPoints, bPoints []Vector, axis Vector, response *Response) bool {
	offsetV := Vector{}.Copy(bPos).Sub(aPos)
	projectedOffset := offsetV.Dot(axis)
	rangeA := flattenPointsOn(aPoints, axis)
	rangeB := flattenPointsOn(bPoints, axis)
	rangeB[0] += projectedOffset
	rangeB[1] += projectedOffset
	if rangeA[0] > rangeB[1] || rangeB[0] > rangeA[1] {
		return true
	}

	if response != nil {
		overlap := 0.0
		if rangeA[0] < rangeB[0] {
			response.AInB = false
			if rangeA[1] < rangeB[1] {
				overlap = rangeA[1] - rangeB[0]
				response.BInA = false
			} else {
				option1 := rangeA[1] - rangeB[0]
				option2 := rangeB[1] - rangeA[0]
				if option1 < option2 {
					overlap = option1
				} else {
					overlap = -option2
				}
			}
		} else {
			response.BInA = false
			if rangeA[1] > rangeB[1] {
				overlap = rangeA[0] - rangeB[1]
				response.AInB = false
			} else {
				option1 := rangeA[1] - rangeB[0]
				option2 := rangeB[1] - rangeA[0]
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
			response.OverlapN.Copy(axis)
			if overlap < 0 {
				response.OverlapN.Reverse()
			}
		}
	}
	return false
}

func flattenPointsOn(points []Vector, normal Vector) []float64 {
	result := []float64{0, 0, 0, 0, 0, 0, 0}
	min := math.MaxFloat64
	max := -math.MaxFloat64
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
	result[0] = min
	result[1] = max
	return result
}
