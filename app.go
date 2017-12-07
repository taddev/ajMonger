package main

import (
	"encoding/json"
	"fmt"
	"github.com/taddev/ajMonger/models"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	person := models.Person{"Bob", "Hope", "bob", "", nil}
//	fmt.Printf("%v\n", person)

	password := []byte("badwolf")
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	person.Hash = string(hash)

	//valid, _ := person.IsPasswordValid("badwolf1")
	//if !valid {
//		fmt.Println("Invalid")
//	} else {
//		fmt.Println("Valid")
//	}

	person.AddAddress(models.Address{"2121 4th Ave", "Rapid City", "South Dakota", 57702})
	person.AddAddress(models.Address{"2727 N Plaza Dr", "Rapid City", "South Dakota", 57702})

	b, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}
