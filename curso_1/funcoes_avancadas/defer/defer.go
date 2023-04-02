package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1...")
}

func funcao2() {
	fmt.Println("Executando a função 2...")
}

func alunoEstaAprovado(n1 int, n2 int) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Somente exemplo de defer acima")
	media := (n1 + n2) / 2

	if media < 6 {
		return false
	}

	return true
}

func main() {
	fmt.Println("Função Defer...")
	funcao1()
	funcao2()

	fmt.Println("defer...")
	defer funcao1() // adiando a execução desta função até o ultimo momento possível
	funcao2()

	n1 := 10
	n2 := 12
	fmt.Println(alunoEstaAprovado(n1, n2))
}
