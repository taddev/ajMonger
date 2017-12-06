package main

import (
	"encoding/json"
	"fmt"
	"github.com/taddev/ajMonger/models"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	person := models.Person{"Bob", "Hope", "bob", ""}
	fmt.Printf("%v\n", person)

	password := []byte("badwolf")
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	person.Hash = string(hash)

	b, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)

	valid, _ := person.IsPasswordValid("badwolf1")
	if !valid {
		fmt.Println("Invalid")
	} else {
		fmt.Println("Valid")
	}
}
