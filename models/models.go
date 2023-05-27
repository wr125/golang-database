package models

import "encoding/json"

type Address struct {
	City     string
	Borough  string
	Country  string
	Postcode string
}
type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}
