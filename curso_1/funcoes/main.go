package main

import "fmt"

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	// variavel do tipo função
	var f = func(n1 int8, n2 int8) int8 {
		return n1 - n2
	}

	fmt.Println(f(20, 10))

	valor1, valor2 := calculosMatematicos(20, 20)

	fmt.Println(valor1, valor2)

}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// mais de um retorno
func calculosMatematicos(n1 int8, n2 int8) (int8, int8) {
	return (n1 + n2), (n1 / n2)
}
