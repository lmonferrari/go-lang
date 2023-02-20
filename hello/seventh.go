package main

import (
	"fmt"
)

func main() {
	Student := make(map[string]int)

	Student["Luiz"] = 35
	Student["Ana Clara"] = 3

	fmt.Println(Student)
	fmt.Println(Student["Luiz"])

	delete(Student, "Ana Clara")
	fmt.Println(Student)

	super_hero := map[string]map[string]string{
		"Superman": map[string]string{
			"RealName": "Clark Kent",
			"City":     "Metropolis",
		},
		"Batman": map[string]string{
			"RealName": "Bruce Wayne",
			"City":     "Gotham City",
		},
	}

	if temp, hero := super_hero["Superman"]; hero {
		fmt.Println(temp["RealName"], temp["City"])
	}

}
