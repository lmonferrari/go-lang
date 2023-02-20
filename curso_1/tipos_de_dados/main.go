package main

import (
	"errors"
	"fmt"
)

func main() {
	var inteiro1 int16 = 100
	var inteiro2 int32 = 2000000000
	var inteiro3 uint32 = 1000000000
	var inteiro4 int64 = 1000000000000000000

	// int32 = rune - geralmente usado para representação de caracteres
	var inteiro5 rune = 1223434

	// byte = uint8
	var inteiro6 byte = 123

	fmt.Println(inteiro1, inteiro2, inteiro3, inteiro4, inteiro5, inteiro6)

	var real1 float32 = 63446.45345
	var real2 float64 = 352345.36453456

	fmt.Println(real1, real2)

	var texto string = "string em go"

	// variavel com valor zero
	var zero int
	fmt.Println(texto, zero)

	var variavelboleana bool = false
	fmt.Println(variavelboleana)

	// Variavel do tipo erro
	var erro error
	fmt.Println(erro)

	erro = errors.New("Erro interno.")
	fmt.Println(erro)

}
