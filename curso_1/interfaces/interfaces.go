package main

import (
	"fmt"
	"math"
)

// Interfaces são tipos que definem um comportamento
type forma interface {
	area() float64
}

// Função que recebe um tipo que implementa a interface forma
func escreverArea(f forma) {
	fmt.Printf("%f\n", f.area())
}

// Definindo tipos que implementam a interface forma
type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

// Implementando o método area para os tipos retangulo e circulo
func (r *retangulo) area() float64 {
	return r.altura * r.largura
}

func (c *circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	var r retangulo = retangulo{10, 15}
	var c circulo = circulo{5}

	escreverArea(&r)
	escreverArea(&c)
}
