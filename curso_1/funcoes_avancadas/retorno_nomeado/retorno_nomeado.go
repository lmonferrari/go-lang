package main

import "fmt"

func calculosMatematicos(n1 int, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2

	return //ele ja entende que terá que pegar as duas variaveis acima e retornar
}

func main() {
	numero1 := 2
	numero2 := 3

	fmt.Println(calculosMatematicos(numero1, numero2))
}
