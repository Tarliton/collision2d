package main

import (
	"fmt"
	sat "github.com/tarliton/sat-go"
)

func main() {
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
