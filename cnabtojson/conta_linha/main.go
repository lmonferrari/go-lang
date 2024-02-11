package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filename := "../bemol/CON01012023160800.rem"
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	linhas := strings.Split(string(content), "\n")

	for i, linha := range linhas {
		tamanho := len(linha)
		if tamanho < 500 {
			fmt.Printf("Linha: %d - Caracteres: %d\n", i-1, tamanho)
		}
	}

}
