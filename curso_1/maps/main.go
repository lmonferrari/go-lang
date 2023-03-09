package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps...")

	user1 := map[int]string{
		1: "Luiz",
		2: "Luke",
	}

	fmt.Println(user1[1], user1[2])

	user2 := map[int]map[string]string{
		1: {
			"primeiro": "Luiz",
			"ultimo":   "Monferrari",
		},
		2: {
			"primeiro": "Luke",
			"ultimo":   "Skywalker",
		},
	}

	fmt.Println(user2[1]["primeiro"], user2[2]["primeiro"])

	// excluindo elemento
	delete(user2, 1)

	fmt.Println(user2)

	// adicionando elemento

	user2[1] = map[string]string{
		"primeiro": "Luiz",
		"ultimo":   "Monferrari",
	}
	fmt.Println(user2)

}
