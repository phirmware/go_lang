package main

import "fmt"

type contactInfo struct {
	address string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	john := person{
		firstName: "john",
		lastName:  "paul",
		contact: contactInfo{
			address: "Alaska",
			zipCode: 12333,
		},
	}

	john.updatePerson("Tobi")
	john.print()
}

func (p *person) updatePerson(firstName string) {
	(*p).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
