package main

import "fmt"

// define a struct

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Value receiver (similar to Getter in C#)
func (p Person) greet() string {
	return "Hello, my name is: " + p.firstName + " " + p.lastName
}

// Pointer receiver
func (p *Person) hasBirthday() {
	p.age++
}

// pointer receiver
func (p *Person) getMarried(spouseName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseName
	}
}

func main() {
	// Init Struct
	person := Person{firstName: "Jog", lastName: "Sam", city: "Abuja", gender: "M", age: 23}
	//Alternative
	person1 := Person{"Jane", "Ashley", "Somalia", "F", 23}

	fmt.Println(person)
	// Get single prop

	fmt.Println(person.age)
	person.age++

	fmt.Println(person)
	fmt.Println(person1)

	//Value receiver
	fmt.Println(person1.greet())

	// Pointer to reciever
	person1.hasBirthday()
	fmt.Println(person1)

	// Change the last name for Jane
	person1.getMarried("Ahmed")
	fmt.Println(person1.greet())
}
