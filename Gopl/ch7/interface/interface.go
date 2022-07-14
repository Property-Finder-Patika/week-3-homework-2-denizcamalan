package main

import "fmt"

type Name struct {
	name string
}

func (n Name) Name() string {
	return n.name
}

type Department struct {
	Name
}

type Worker interface {
	work()
}

type HardWorker interface {
	Worker
	hardWork()
}

type Employee struct {
	Name
	hour int
}

func (e *Employee) work() {
	fmt.Println(e.Name, "is working.")
	e.hour++
}

type Manager struct {
	Employee
	Department
}

func (m *Manager) hardWork() {
	fmt.Println(m.Employee.Name, "is working hard.")
	m.hour += 2
}

func main() {
	var e Employee = Employee{Name{"Ali"}, 0}
	e.work()
	fmt.Printf("%#v\n", e)

	var m Manager = Manager{Employee{Name{"Nilay"}, 0}, Department{Name{"Sales"}}}
	m.hardWork()
	fmt.Printf("%#v", m)
}
