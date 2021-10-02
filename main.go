package main

import (
	"fmt"
	solver "solver/src"
)

func main() {
	b := solver.BuildBoard()
	b.Print()

	for i := 0; i <= 81; i++ {
		didSomething := b.Solve()

		if didSomething {
			b.Print()
		} else {
			fmt.Print("Unable to solve further \n")
			return
		}
	}
	//
	//for i := 0; i < 5; i++ {
	//	b.Solve()
	//	b.Print()
	//}
}
