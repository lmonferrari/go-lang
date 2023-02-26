package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	age        uint8
}

type student struct {
	person
	school string
	course string
}

func main() {
	fmt.Println("Herança")

	user1 := student{person{"Luke", "Skywalker", 25}, "Jedi Academy", "Padawn"}
	fmt.Println(user1.first_name, user1.last_name, user1.age, user1.school, user1.course)
}
