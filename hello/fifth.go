package main

import "fmt"

func main() {

	age := 18

	if age >= 18 {
		fmt.Println("You can vote!")
	} else {
		fmt.Println("No, you can't vote!")
	}

	switch age {
	case 17:
		fmt.Println("Under age")
	case 18:
		fmt.Println("Don't chase girls!")
	default:
		fmt.Println("I don't know :D")
	}
}
