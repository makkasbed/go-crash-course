package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

//greeting method(value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

//getMarried  (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	person1 := Person{
		firstName: "Adu",
		lastName:  "Asare",
		city:      "Amasaman",
		gender:    "F",
		age:       37,
	}
	fmt.Println(person1.firstName, person1.age, person1.gender)
	person1.hasBirthday()
	person1.getMarried("Acquah")
	fmt.Println(person1.greet())
}
