package main

import "fmt"

type contact struct {
	email	string
	phone	string
}

type person struct {
	firstName	string
	lastName 	string
	contact
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateFirstName(firstName string) {
	p.firstName = firstName
}

func main() {
	me := person{firstName: "Jakob", lastName: "Rolik", contact: contact{
		email: "rolik.jakob@gmail.com",
		phone: "+421 948 509 880",
	}}

	me.print()

	me.updateFirstName("Jacob")
	me.print()
}
