package main

import "fmt"

// person is a dummy struct
type person struct {
	Name string
}

func main() {
	var v interface{}
	p := person{"Sam"}
	v = &p                                   // assign the pointer to v
	fmt.Printf("Hello %p of typed %T", v, v) // print pointer value & type
}
