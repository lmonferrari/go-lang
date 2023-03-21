package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Loops...")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0

	for i < 10 {
		i++
		fmt.Println(i + 10)
	}

	a := []int{30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	for i, value := range a {
		fmt.Println(i, value)
	}

	usuario := map[string]string{
		"nome":      "Luiz",
		"sobrenome": "Monferrari",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Loop eterno")
		time.Sleep(time.Second)
	}
}
