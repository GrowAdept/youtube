// methods (pointer receivers)
package main

import (
	"fmt"
)

type employee struct {
	name   string
	salary float64
}

// receives pointer (memory address) of original struct
func (e *employee) doesChangeStr() {
	e.salary = e.salary * 1.05
	fmt.Println("memory address of e.salary:", &e.salary)
	fmt.Println("e.salary:", e.salary)
}

//creates copy of type and passes it to the function, not the original
func (e employee) doesNotChangeStr() {
	fmt.Println("memory address of arguement e.salary:", &e.salary)
	e.salary = e.salary * 1.05
	fmt.Println("e.salary:", e.salary)
}

func main() {

	var jim = employee{"Jim Wilson", 70000}
	fmt.Println("memory address of jim.salary:", &jim.salary)
	jim.doesNotChangeStr()
	fmt.Println("jim.salary:", jim.salary)

	/*
		var jim = employee{"Jim Wilson", 70000}
		fmt.Println("memory address of jim.salary:", &jim.salary)
		jim.doesChangeStr()
		fmt.Println("jim.salary:", jim.salary)
	*/

	/*
		var jane = employee{"Judith Swanson", 90000}
		var mark = employee{"Mark Lincoln", 44000}
		var beth = employee{"Beth Jacobs", 50000}
		var kelly = employee{"Kelly Smith", 55000}
		var ben = employee{"Ben Anderson", 70000}

		var workforce = []*employee {&jane, &mark, &beth, &kelly, &ben}
		payRaise(workforce, .1)
		fmt.Printf("%.0f", beth.salary)
	*/
}

func payRaise(wf []*employee, prctIncr float64) {
	for _, v := range wf {
		v.salary = v.salary * (1 + prctIncr)
	}
}
