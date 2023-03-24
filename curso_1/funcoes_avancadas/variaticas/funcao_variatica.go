package main

import "fmt"

func soma(numeros ...int) int {
	soma := 0

	for _, value := range numeros {
		soma += value
	}
	return soma
}

func main() {
	fmt.Println(soma(1, 2, 3, 4, 5))
}
