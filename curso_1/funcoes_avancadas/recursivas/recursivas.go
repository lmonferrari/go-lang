package main

import "fmt"

func fibonacci(posicao int) int {
	if posicao < 0 {
		return -1
	}

	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções recursivas...")

	posicao := -1
	fmt.Println(fibonacci(posicao))
}
