package main

import (
	"fmt"
	"github.com/Tarliton/collision2d"
)

func main() {
	circle1 := collision2d.Circle{collision2d.Vector{}, 20}
	circle2 := collision2d.Circle{collision2d.Vector{30, 0}, 30}
	colided, response := collision2d.TestCircleCircle(circle1, circle2)
	fmt.Println(colided)
	fmt.Println(response)

	point := collision2d.Vector{5, 5.321312}
	b := point
	b.X = 999999
	fmt.Println(b)
	circle := collision2d.Circle{point, 123123}
	polygon := new(collision2d.Polygon)
	box := new(collision2d.Box)
	fmt.Print(point)
	fmt.Print(point)
	fmt.Println("--------------")
	fmt.Println(circle)
	fmt.Println("--------------")
	fmt.Println(polygon)
	fmt.Println("--------------")
	fmt.Println(box)
	fmt.Println("--------------")
	fmt.Println(point.Perp().Perp().Rotate(180))
}
