package main

import "fmt"

type user struct {
	nome     string
	idade    uint8
	endereco address
}

type address struct {
	logradouro string
	numero     uint
}

func main() {
	fmt.Println("Structs")

	var usuario1 user
	usuario1.nome = "Luiz"
	usuario1.idade = 34
	usuario1.endereco.logradouro = "Rua 10"
	usuario1.endereco.numero = 32
	fmt.Println(usuario1, usuario1.nome, usuario1.idade, usuario1.endereco.logradouro, usuario1.endereco.numero)

	usuario2 := user{"Luke", 7, address{"Rua 20", 32}}
	fmt.Println(usuario2, usuario2.nome, usuario2.idade, usuario2.endereco.logradouro, usuario2.endereco.numero)

	usuario3 := user{nome: "Bruce"}
	fmt.Println(usuario3, usuario3.nome)
}
