package Shapes

import "strconv"

type Square struct {
	X, Y int64
}

func (square Square) Draw() string {
	return "This square has x " + strconv.FormatInt(square.X, 10) + " and y " + strconv.FormatInt(square.Y, 10)
}