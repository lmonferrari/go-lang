package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
}

func main() {

	var randomNum int

	if randomNum = rand.Intn(10); randomNum == 0 {
		randomNum += 1
	}
	endpoint := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", randomNum)
	log.Printf("Endpoint: %s", endpoint)

	response, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Failed to fech user:", err.Error())
		return
	}

	defer response.Body.Close()

	var user User
	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		fmt.Println("Failed to parse", err.Error())
		return
	}

	fmt.Printf("User: %d\nName: %s\nUsername: %s\nEmail: %s", user.ID, user.Name, user.Username, user.Email)
}
