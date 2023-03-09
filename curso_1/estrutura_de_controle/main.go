package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle...")

	numero := 12

	if numero >= 10 {
		fmt.Println("Número maior ou igual 10")
	} else {
		fmt.Println("Número menor que 10")
	}

	// if init

	// o escopo da variável outroNumero fica limitada ao if
	if outroNumero := 12; outroNumero >= 10 {
		fmt.Println("Número maior ou igual 10")
	} else {
		fmt.Println("Número menor que 10")
	}

}
