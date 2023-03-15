package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Inválido"
	}

}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda"
	case numero == 3:
		return "Terça"
	case numero == 4:
		return "Quarta"
	case numero == 5:
		return "Quinta"
	case numero == 6:
		return "Sexta"
	case numero == 7:
		return "Sabado"
	default:
		return "Inválido"
	}
}

func diaDaSemana3(numero int) string {
	var dia_da_semana string

	switch {
	case numero == 1:
		dia_da_semana = "Domingo"
	case numero == 2:
		dia_da_semana = "Segunda"
	case numero == 3:
		dia_da_semana = "Terça"
	case numero == 4:
		dia_da_semana = "Quarta"
	case numero == 5:
		dia_da_semana = "Quinta"
	case numero == 6:
		dia_da_semana = "Sexta"
	case numero == 7:
		dia_da_semana = "Sabado"
	default:
		dia_da_semana = "Inválido"
	}

	return dia_da_semana
}

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

	fmt.Println(diaDaSemana(2))

	dia := 3
	fmt.Println(diaDaSemana2(dia))

	dia = 4
	fmt.Println(diaDaSemana3(dia))
}
