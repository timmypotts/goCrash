package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName string
	// city string
	// gender string
	// age int

	firstName, lastName, city, gender string
	age int
}

// Greeting method
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old."
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	fmt.Println("Structs")
	// init person
	person1 := Person{firstName: "Tim", lastName: "Potts", city: "Denver", gender: "m", age: 25}

	person2 := Person{"Sav", "Lankford", "Denver", "f", 24}

	fmt.Println(person1, person2)

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())

	person1.hasBirthday()
	fmt.Println(person1.greet())

	person2.getMarried("Potts")
	fmt.Println(person2.greet())

}