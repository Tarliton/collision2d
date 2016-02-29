package main

import (
	"fmt"
	"github.com/tarliton/collision2d"
)

func main() {
	circle1 := sat.Circle{sat.Vector{}, 20}
	circle2 := sat.Circle{sat.Vector{30, 0}, 30}
	response := new(sat.Response)
	colided := sat.TestCircleCircle(circle1, circle2, response)
	fmt.Println(colided)
	fmt.Println(response.Overlap)
	fmt.Println(response.OverlapV.X)
	fmt.Println(response.OverlapV.Y)
	fmt.Println("ae")

	point := sat.Vector{5, 5.321312}
	b := point
	b.X = 999999
	fmt.Println(b)
	circle := sat.Circle{point, 123123}
	polygon := new(sat.Polygon)
	box := new(sat.Box)
	fmt.Println(point)
	fmt.Println(circle)
	fmt.Println(polygon)
	fmt.Println(box)
	fmt.Println(point.Perp().Perp().Rotate(180))
}
