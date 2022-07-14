package main

import (
	"fmt"
	"syscall"
)

func main() {
	var err error = syscall.Errno(0)
	fmt.Println(err)
	err = syscall.Errno(1)
	fmt.Println(err)
	err = syscall.Errno(2)
	fmt.Println(err)
	err = syscall.Errno(3)
	fmt.Println(err)
}
