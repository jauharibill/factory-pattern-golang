package Shapes

import "strconv"

type Circle struct {
	Diameter int64
}

func (circle Circle) Draw() string {
	return "This Circle has Diameter " + strconv.FormatInt(circle.Diameter, 10)
}
