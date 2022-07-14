package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	mathFunction "week-3-homework-2-denizcamalan/Gopl/ch7/math"
)

func main() {
	//assert1()
	//assert2()
	//wrongWay()
	correctWay1(1)
	correctWay2(0.5)
	correctWay3(1)
}

func assert1() {
	var w io.Writer
	fmt.Printf("%T\n", w) // nil
	w = os.Stdout
	fmt.Printf("%T\n", w)
	x := w.(io.ReadWriter)
	fmt.Printf("%T\n", x)
	y := w.(*os.File)
	fmt.Printf("%T\n", y)

	//t := w.(os.Stdout) // Can't do that because os.Stdout is not a type, it is a variable of type *File
	// So let's get its type
	//fmt.Printf("%T\n", reflect.TypeOf(os.Stdout))
	//fmt.Printf("%s\n", reflect.TypeOf(os.Stdout))
	//typeOfStdout := reflect.TypeOf(os.Stdout)
	//t := w.(typeOfStdout) // Still can't do that because it is of type *reflect.rtype even though it has a string that holds a value of *os.File.

	// How about that?
	//z := w.(*bytes.Buffer)
	//fmt.Printf("%T\n", z)
	// Do the same thing after following assignment
	w = new(bytes.Buffer)
	z := w.(*bytes.Buffer) // *bytes.Buffer satisfies io.Writer by implementing its unique method Write(): func (b *Buffer) Write(p []byte) (n int, err error)
	fmt.Printf("%T\n", z)
}

func assert2() {
	var mf mathFunction.MathFunction
	fmt.Printf("%T\n", mf)

	mf = mathFunction.Sin{Name: "Sin"}
	fmt.Printf("%T\n", mf)
	t := mf.(mathFunction.MathFunction)
	fmt.Printf("%T\n\n", t)

	mf = mathFunction.Log{Name: "Log"}
	fmt.Printf("%T\n", mf)
	t = mf.(mathFunction.MathFunction)
	fmt.Printf("%T\n\n", t)

	mf = mathFunction.Cos{Name: "Cos"}
	fmt.Printf("%T\n", mf)
	t = mf.(mathFunction.MathFunction)

	fmt.Printf("%T\n\n\n", t)

	// How about that
	//t = mf.(io.ReadWriter) // Impossible conversion because io.ReadWriter can't be assigned to mathFunction.MathFunction
	//fmt.Printf("%T\n", t)

	var amf mathFunction.MathFunction
	fmt.Printf("%T\n", amf)
	amf = mathFunction.Cos{Name: "Cos"}
	fmt.Printf("%T\n", mf)
	t = amf.(mathFunction.Cos)
	fmt.Printf("%T\n\n", t)
	t = amf.(mathFunction.MathFunction)
	fmt.Printf("%T\n\n", t)
	// How about that
	t = amf.(mathFunction.Sin)
	fmt.Printf("%T\n\n", t)
}

func wrongWay() {
	var mf mathFunction.MathFunction = mathFunction.MathFunctionFactory()
	var mfType mathFunction.Sin = mf.(mathFunction.Sin) // Conversion (down-casting in some other languages)
	fmt.Println(mfType)
}

func correctWay1(arg float64) {
	var mf mathFunction.MathFunction = mathFunction.MathFunctionFactory()
	if mf == nil {
		fmt.Printf("It is nil.")
	} else if _, ok := mf.(mathFunction.Sin); ok {
		fmt.Printf("It is mathFunction.Sin")
		fmt.Printf("\nSin of %f is %f\n", arg, mf.Calculate(arg))
		var sinFunc mathFunction.Sin = mf.(mathFunction.Sin)
		fmt.Printf("Sin of %f is %f\n", arg, sinFunc.Calculate(arg))
	} else if _, ok := mf.(mathFunction.Cos); ok {
		fmt.Printf("It is mathFunction.Cos")
		fmt.Printf("\nCos of %f is %f\n", arg, mf.Calculate(arg))
		var cosFunc mathFunction.Cos = mf.(mathFunction.Cos)
		fmt.Printf("Cos of %f is %f\n", arg, cosFunc.Calculate(arg))
	} else if _, ok := mf.(mathFunction.Log); ok {
		fmt.Printf("It is mathFunction.Log")
		fmt.Printf("\nLog of %f is %f\n", arg, mf.Calculate(arg))
		var logFunc mathFunction.Log = mf.(mathFunction.Log)
		fmt.Printf("Log of %f is %f\n", arg, logFunc.Calculate(arg))
	} else {
		panic("shouldn't come here!")
	}
}

func correctWay2(arg float64) {
	var mf mathFunction.MathFunction = mathFunction.MathFunctionFactory()
	switch mf.(type) {
	case nil:
		fmt.Printf("It is nil.")
	case mathFunction.Sin:
		fmt.Printf("It is mathFunction.Sin and ")
		fmt.Printf("\nSin of %f is %f", arg, mf.Calculate(arg))
	case mathFunction.Cos:
		fmt.Printf("It is mathFunction.Cos and ")
		fmt.Printf("\nCos of %f is %f", arg, mf.Calculate(arg))
	case mathFunction.Log:
		fmt.Printf("It is mathFunction.Log and ")
		fmt.Printf("\nLog of %f is %f", arg, mf.Calculate(arg))
	default:
		panic("shouldn't come here!")
	}
}

func correctWay3(arg float64) {
	var mf mathFunction.MathFunction = mathFunction.MathFunctionFactory()
	switch mf.(type) {
	case nil:
		fmt.Printf("It is nil.")
	case mathFunction.Sin:
		{
			fmt.Printf("It is mathFunction.Sin.")
			fmt.Printf("\nSinus of %f is %f", arg, mf.Calculate(arg))
		}
	case mathFunction.Cos:
		{
			fmt.Printf("It is mathFunction.Cos.")
			fmt.Printf("\nCosinus of %f is %f", arg, mf.Calculate(arg))
		}
	case mathFunction.Log:
		{
			fmt.Printf("It is mathFunction.Log.")
			fmt.Printf("\nLog of %f is %f", arg, mf.Calculate(arg))
		}
	default:
		panic("shouldn't come here!")
	}
}
