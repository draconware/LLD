package main

type IShape interface {
	GetShape() string
}

type Square struct {
	name string
}

func NewSquareShape() *Square {
	return &Square{
		name: "Square shape",
	}
}

func (sq *Square) GetShape() string {
	return sq.name
}

type Rectangle struct {
	name string
}

func NewRectangleShape() *Rectangle {
	return &Rectangle{
		name: "Rectangle shape",
	}
}

func (rt *Rectangle) GetShape() string {
	return rt.name
}

type ShapeFactory struct{}

func (sf *ShapeFactory) GetShapeObject(shape string) IShape {
	switch shape {
	case "square":
		return NewSquareShape()
	case "rectangle":
		return NewRectangleShape()
	default:
		return nil
	}
}
