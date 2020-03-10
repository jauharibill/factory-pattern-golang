package Shapes

import "strconv"

type Rectangle struct {
	Width, Height int64
}

func (rectangle Rectangle) Draw() string {
	return "This rectangle has width " + strconv.FormatInt(rectangle.Width, 10) + " & Height " + strconv.FormatInt(rectangle.Height, 10)
}
