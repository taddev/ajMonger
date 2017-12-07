package models

type Address struct {
	Address string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     int	`json:"zip"`
}

type Addresses []Address

