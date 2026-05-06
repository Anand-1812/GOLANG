package main

import (
	"fmt"
	"methods/geometry"
	"methods/list"
)

func main() {
	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{X: 4, Y: 6}

	distanceFromP := p.Distance
	fmt.Println("Distance from p is ", distanceFromP(q))

	distFunc := geometry.Point.Distance
	fmt.Println("Distance (Expr):", distFunc(p, q))

	var myList *list.IntList
	fmt.Println("Sum of nil list: ", myList.Sum())

	p.ScaleBy(2)
	fmt.Printf("Scaled p: %+v\n", p)
}

