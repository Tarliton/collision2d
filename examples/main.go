package main

import (
	"fmt"
	"github.com/Tarliton/collision2d"
)

func main() {
	circle1 := collision2d.Circle{collision2d.Vector{}, 20}
	circle2 := collision2d.Circle{collision2d.Vector{30, 0}, 30}
	response := new(collision2d.Response)
	colided := collision2d.TestCircleCircle(circle1, circle2, response)
	fmt.Println(colided)
	fmt.Println(response.Overlap)
	fmt.Println(response.OverlapV.X)
	fmt.Println(response.OverlapV.Y)
	fmt.Println("ae")

	point := collision2d.Vector{5, 5.321312}
	b := point
	b.X = 999999
	fmt.Println(b)
	circle := collision2d.Circle{point, 123123}
	polygon := new(collision2d.Polygon)
	box := new(collision2d.Box)
	fmt.Println(point)
	fmt.Println(circle)
	fmt.Println(polygon)
	fmt.Println(box)
	fmt.Println(point.Perp().Perp().Rotate(180))
}
