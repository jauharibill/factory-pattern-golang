package main

import (
	"FactoryPattern/Shapes"
	"fmt"
)

func getShape(shape Shapes.Shape) {

	fmt.Println(shape.Draw())

}
