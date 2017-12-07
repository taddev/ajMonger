package models

import (
	"golang.org/x/crypto/bcrypt"
)

type Person struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
	Hash      string `json:"hash,omitempty"`
	Addresses Addresses `json:"addresses"`
}

func (p *Person) IsPasswordValid(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(p.Hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Person) AddAddress(address Address) Addresses {
	p.Addresses = append(p.Addresses, address)
	return p.Addresses
}
