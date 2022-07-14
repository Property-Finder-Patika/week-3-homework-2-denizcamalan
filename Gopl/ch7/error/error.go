package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	error1()
	//error2()
}

func error1() {

	//var err1 errors.error // Can't say this! error is not exported!
	var err1 error = errors.New("a new error") // error is a built-in or predeclared identifier
	//err1 := errors.New("a new error") // Can use this way too
	fmt.Printf("%s => %T\n", err1, err1)

	var err2 error = fmt.Errorf("another error occurred at: %v", time.Now())
	fmt.Printf("%s => %T\n", err2, err2)

	var err3 error = &GrumblingError{"grumbling..."}
	fmt.Printf("%s => %T\n", err3, err3)
}

func error2() {
	var selim Human = Human{"Selim", true}
	s1, err1 := selim.greet()
	if err1 != nil {
		fmt.Printf("%s\n", err1)
	} else {
		fmt.Printf("%s\n", s1)
	}

	var ilyas Human = Human{"ilyas", false}
	s2, err2 := ilyas.greet()
	if err2 != nil {
		fmt.Printf("%s\n", err2)
	} else {
		fmt.Printf("%s\n", s2)
	}
}

type greet interface {
	greet() (string, error)
}

type Human struct {
	name     string
	cheerful bool
}

func (h *Human) greet() (string, error) {
	if h.cheerful {
		return "hello " + h.name + " :)", nil
	} else {
		var g GrumblingError = GrumblingError{"grumbling..."}
		return "", &g
	}
}

type GrumblingError struct {
	message string
}

func (g *GrumblingError) Error() string {
	return g.message
}
