package main

import "fmt"

type Shape interface {
	getType() string
	accept(Visitor)
}

type Visitor interface {
	visitForSquare(*Square)
	visitForRectange(*Rectangle)
}

type Square struct {
	side int
}

func (sq *Square) getType() string {
	return "Square"
}
func (sq *Square) accept(v Visitor) {
	v.visitForSquare(sq)
}

type Rectangle struct {
	length  int
	breadth int
}

func (rt *Rectangle) getType() string {
	return "Rectangle"
}
func (rt *Rectangle) accept(v Visitor) {
	v.visitForRectange(rt)
}

type AreaCalculator struct{}

func (ac *AreaCalculator) visitForRectange(r *Rectangle) {
	fmt.Println("Calculating area for rectangle.")
}
func (ac *AreaCalculator) visitForSquare(sq *Square) {
	fmt.Println("Calculating area for square.")
}

type MiddlePointsCalculator struct{}

func (ac *MiddlePointsCalculator) visitForRectange(r *Rectangle) {
	fmt.Println("Calculating MiddlePoints for rectangle.")
}
func (ac *MiddlePointsCalculator) visitForSquare(sq *Square) {
	fmt.Println("Calculating MiddlePoints for square.")
}
