package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@mail.com",
			zipCode: 94000,
		},
	}
	fmt.Printf("%+v", alex)
}
