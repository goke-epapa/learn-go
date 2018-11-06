package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	// Using Order of fields
	alex := person{"Alex", "Anderson"}

	// Listing all the fields
	goke := person{firstName: "Goke", lastName: "Obasa"}

	// john initialised to a zero value
	var john person

	fmt.Println(alex)
	fmt.Println(goke)
	fmt.Println(john)

	// updating structs
	alex.firstName = "Alexander"
	alex.lastName = "Cleaopatra"

	// Formatting output
	fmt.Printf("%+v", alex)
}
