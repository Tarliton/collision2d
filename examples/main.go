package main

import (
	"fmt"
	sat "github.com/tarliton/sat-go"
)

func main() {
	point := new(sat.Vector)
	circle := new(sat.Circle)
	polygon := new(sat.Polygon)
	box := new(sat.Box)
	fmt.Println(point)
	fmt.Println(circle)
	fmt.Println(polygon)
	fmt.Println(box)
}
