package main

import "FactoryPattern/Shapes"

func main() {

	circle := Shapes.Circle{Diameter: 5}
	rectangle := Shapes.Rectangle{Width: 6, Height: 7}

	getShape(circle)
	getShape(rectangle)

}
