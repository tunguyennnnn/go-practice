package main

import (
	"fmt"
)

type ContactInfo struct {
	phone   string
	address string
}

type Person struct {
	firstName   string
	lastName    string
	contactInfo ContactInfo
}

func (p *Person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	jim := Person{
		firstName: "Jim",
		lastName:  "Cammel",
		contactInfo: ContactInfo{
			phone:   "514 473 1234",
			address: "abcde",
		},
	}
	jimPointer := &jim
	jim.updateName("kin")
	fmt.Println(jimPointer)
	fmt.Println(jim)
}
