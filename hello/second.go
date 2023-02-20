package main

import "fmt"

func main() {

	const pi float64 = 3.14159265359

	fmt.Println("Constant pi value: ", pi)

	var (
		varA = 2
		varB = 3
	)

	fmt.Println(varA, varB)

	x, y := 14, 15

	fmt.Println(x, y)

	var name string = "Luiz Monferrari"

	fmt.Println(len(name))
}
