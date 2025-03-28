package main

import "fmt"

func main() {
	sq := &Square{side: 2}
	rt := &Rectangle{length: 2, breadth: 2}

	areaCalculator := &AreaCalculator{}
	sq.accept(areaCalculator)
	rt.accept(areaCalculator)

	fmt.Println()

	middlePointCalculator := &MiddlePointsCalculator{}
	sq.accept(middlePointCalculator)
	rt.accept(middlePointCalculator)
}
