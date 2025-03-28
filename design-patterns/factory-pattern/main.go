package main

import "fmt"

func main() {
	factory := ShapeFactory{}

	shape1 := factory.GetShapeObject("square")
	shape2 := factory.GetShapeObject("rectangle")

	fmt.Println("Shape 1: ", shape1.GetShape())
	fmt.Println("Shape 2: ", shape2.GetShape())
}
