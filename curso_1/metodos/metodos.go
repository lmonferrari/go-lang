package main

import "fmt"

// criando um tipo de dado
type usuario struct {
	nome  string
	idade uint16
}

// método de usuário
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário: Nome %s - idade %d\n", u.nome, u.idade)
}

func (u usuario) maiorIdade() bool {
	return u.idade >= 18
}

func main() {
	var u usuario = usuario{"Joãozinho", 30}
	u.salvar()
	fmt.Println("Maior de idade?", u.maiorIdade())
}
