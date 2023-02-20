package main

import "fmt"

func main() {
	//OPERADORES

	//ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 1 / 2
	multiplicacao := 1 * 2
	restoDivisao := 1 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	//ATRIBUIÇÃO
	var variavel1 string = "texto"
	variavel2 := "texto"

	fmt.Println(variavel1, variavel2)

	//RELACIONAIS
	fmt.Println(soma > subtracao)
	fmt.Println(restoDivisao >= soma)
	fmt.Println(multiplicacao == soma)
	fmt.Println(soma != divisao)

	//LOGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(falso || verdadeiro)
	fmt.Println(!verdadeiro)

	//UNÁRIOS
	numero := 10
	numero++
	fmt.Println(numero)
	numero *= 10
	fmt.Println(numero)
	numero += 5
	fmt.Println(numero)

}
