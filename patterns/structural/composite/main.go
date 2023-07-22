// This pattern is about creating tree-like structures where individual components (leaves) and composition of
// components (composites) are treated uniformly. This is implemented by defining a common interface for both
// the individual objects and the composition of objects.

package main

import "fmt"

// Department Interface for both the individual objects and the composition of objects
type Department interface {
	PrintDepartmentName()
}

type BaseDepartment struct {
	name string
}

func (bd *BaseDepartment) PrintDepartmentName() {
	fmt.Println(bd.name)
}

type CompositeDepartment struct {
	BaseDepartment
	subDepartments []Department
}

func (cd *CompositeDepartment) PrintDepartmentName() {
	cd.BaseDepartment.PrintDepartmentName()
	for _, d := range cd.subDepartments {
		d.PrintDepartmentName()
	}
}

func (cd *CompositeDepartment) AddDepartment(department Department) {
	cd.subDepartments = append(cd.subDepartments, department)
}

func main() {
	marketing := &BaseDepartment{name: "Marketing"}
	sales := &BaseDepartment{name: "Sales"}

	usa := &CompositeDepartment{
		BaseDepartment: BaseDepartment{name: "USA"},
	}
	usa.AddDepartment(marketing)
	usa.AddDepartment(sales)

	global := &CompositeDepartment{
		BaseDepartment: BaseDepartment{name: "Global"},
	}
	global.AddDepartment(usa)

	global.PrintDepartmentName()
}
