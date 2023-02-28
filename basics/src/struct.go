package main

import (
	"fmt"
	"reflect"
)

type person struct {
	firstName string
	lastName  string
	age       int
	// A tag is a string that can be attached to a field
	country string `required max:100` // We can use tags to add metadata to our struct fields
}

func main() {
	// Struct
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.age = 30
	alex.country = "USA"
	fmt.Println(alex)
	fmt.Printf("%+v\n", alex)

	// To get the tag of a field, we use the reflect package
	t := reflect.TypeOf(person{})
	field, _ := t.FieldByName("country")
	fmt.Println(field.Tag)

	// Struct literal
	jim := person{firstName: "Jim", lastName: "Party", age: 50, country: "USA"}
	fmt.Printf("%+v\n", jim)

	// Struct fields
	fmt.Println("First name:", jim.firstName)
	fmt.Println("Last name:", jim.lastName)
	fmt.Println("Age:", jim.age)
	fmt.Println("Country:", jim.country)

	// Anonymous struct
	person1 := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "John",
		lastName:  "Doe",
		age:       25,
	}
	fmt.Printf("%+v\n", person1)

	// Structs are stored in contiguous memory locations
	// Structs are comparable
	// Structs are passed by value
	// Structs are mutable
}
