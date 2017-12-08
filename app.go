package main

import (
	"encoding/json"
	"fmt"
	"github.com/taddev/ajMonger/models"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	person := models.Person{"Bob", "Hope", "bob", "", nil, nil}

	password := []byte("badwolf")
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	person.Hash = string(hash)

	person.AddAddress(models.Address{"2121 4th Ave", "Rapid City", "South Dakota", 57702})
	person.AddAddress(models.Address{"2727 N Plaza Dr", "Rapid City", "South Dakota", 57702})

	person.AddEmail("bob@hope.com")
	person.AddEmail("tom@hanks.com")
	b, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}
