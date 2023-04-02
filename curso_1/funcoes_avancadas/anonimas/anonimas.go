package main

import "fmt"

func main() {
	fmt.Println("Funções anonimas...")

	func(nome string) {
		fmt.Println("Olá ...", nome)
	}("Luiz")

}
