package main

import (
	"fmt"
	mathFunction "week-3-homework-2-denizcamalan/Gopl/ch7/math"
)

// That's about dynamic dispatching when calling methods on an interface type in Go
func main() {

	var mf mathFunction.MathFunction = mathFunction.MathFunctionFactory()
	val := mf.Calculate(0.784) // That's pi/4
	fmt.Printf("%f\n", val)    // Which method is called? sin, cos or log
	//fmt.Printf("%T", mf)
}
