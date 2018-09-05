package main

import "fmt"

type Employee struct {
	name     string
	salary   int
	currency string
}

// displaySalary  方法作为Employee 作为接收器类型
func (e Employee) displaySalary(score int) int {
	fmt.Println(e.name, e.salary, e.currency)
	return e.salary + score
}

func main() {
	emp1 := Employee{"yang", 8888, "$"}

	fmt.Println(emp1.displaySalary(8))

}
