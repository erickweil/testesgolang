package main

import "fmt"

func maiin() {
	var a int = 1
	var b int = 3
	var astar *int = &a

	fmt.Println("A:", a, " A*", *astar, "-->", astar)

	*astar = *astar + 1

	fmt.Println("A:", a, " A*", *astar, "-->", astar)

	astar = &b

	fmt.Println("A:", a, " A*", *astar, "-->", astar)
}
