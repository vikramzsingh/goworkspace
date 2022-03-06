package main

import "fmt"
 
// type Jedi interface {
//     HasForce() bool
// 	M2() string
// }
 
// // treated as class.
// type Knight struct {
// 	Name string
// 	Age  int
// }

// // will through error
// // var _ Jedi = (*Knight)(nil)
  

// func (k *Knight) HasForce() bool {
// 	fmt.Println("Structure k", k)
// 	return true
// }

// func (k *Knight) M2() string {
// 	return fmt.Sprint("done")
// }

// func UseInterface(iface Jedi) {
// 	fmt.Printf("%T\n", iface)
// 	fmt.Println("running by interface implementation:", iface.HasForce())
// 	fmt.Println(iface.M2())
// }

// func main() {
// 	k1 := Knight{
// 		"Vikram",
// 		24,
// 	}
// 	k1.HasForce() // running HasForce by type( as method alloted to type), just like method defined in class 
// 	UseInterface(&k1) // running HasForce by interface ( as anybody has HasForce() is also of jedi type, i.e type Knight have HasForce() method)
// }

// #error
// cannot use k1 (type Knight) as type Jedi in argument to UseInterface:
//         Knight does not implement Jedi (HasForce method has pointer receiver)

// NOTE:- 
// but &Knight(Pointer to Knight) implements jedi INTERFACE.


// EXAMPLE#2

// Employee is an interface for printing employee details
type Employee interface {
	PrintName(name string)
	PrintSalary(basic int, tax int) int
}

// Emp user-defined type
type Emp int

// PrintName method to print employee name
func (e Emp) PrintName(name string) {
	fmt.Println("Employee Id:\t", e)
	fmt.Println("Employee Name:\t", name)
}

// PrintSalary method to calculate employee salary
func (e Emp) PrintSalary(basic int, tax int) int {
	var salary = (basic * tax) / 100
	return basic - salary
}

func main() {
	var e1 Employee
	e1 = Emp(1)
	e1.PrintName("John Doe")
	fmt.Println("Employee Salary:", e1.PrintSalary(25000, 5))
}