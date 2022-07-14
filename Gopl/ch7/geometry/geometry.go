package main

import (
	"fmt"
	myerror "week-3-homework-2-denizcamalan/Gopl/ch7/error/myError"
)

type Shape interface {
	draw()
}

type Rectangle struct {
	short, long int
}

func (r *Rectangle) draw() {
	fmt.Printf("Drawing  rectangle %d by %d\n", r.short, r.long)
}

func New(x, y int) (Rectangle, error) {
	var rectangle Rectangle
	if x < 0 {
		err := myerror.NewNegativeArgumentError("argument is negative", float64(x))
		return rectangle, err
	}
	if y < 0 {
		err := myerror.NewNegativeArgumentError("argument is negative", float64(y))
		return rectangle, err
	}
	rectangle = Rectangle{x, y}
	return rectangle, nil
}

func main() {
	rect1, err := New(2, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		rect1.draw()
	}

	rect2, err := New(-2, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		rect2.draw()
	}
}
