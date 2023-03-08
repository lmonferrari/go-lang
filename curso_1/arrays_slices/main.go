package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays - Slices")

	var array1 [5]int
	fmt.Println(array1)

	for i := range array1 {
		array1[i] = i
	}

	for _, value := range array1 {
		fmt.Println(value)
	}

	array2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	for _, value := range array2 {
		fmt.Println(value)
	}

	slice1 := []int{1, 20, 304, 343, 23, 12}

	for _, value := range slice1 {
		fmt.Print(value, " ")
	}
	fmt.Println()
	slice1 = append(slice1, 902)

	for _, value := range slice1 {
		fmt.Print(value, " ")
	}
	fmt.Println()

	fmt.Println("Variável Array1: ", reflect.TypeOf(array1))
	fmt.Println("Variável Slice1: ", reflect.TypeOf(slice1))

	slice2 := array2[1:3]
	fmt.Println(slice2)

	slice2[1] = 10
	fmt.Println(slice2)
}
