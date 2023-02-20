package main

import "fmt"

func main() {

	var evenNumber [6]int
	j := 0

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			evenNumber[j] = i
			j++
		}
	}

	fmt.Println(evenNumber)

	for _, value := range evenNumber {
		fmt.Println(value)
	}
}
