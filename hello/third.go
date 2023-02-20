package main

import "fmt"

func main() {

	x := 10
	changeValue(&x) /* ponteiro */
	fmt.Println(x)

}

/* ponteiro serve para acessar o endereço da variável que não tem acesso ao escopo*/

func changeValue(x *int) {
	*x = 20
}
