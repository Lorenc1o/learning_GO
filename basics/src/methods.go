package main

import "fmt"

// We can define a method on a type
type greeter struct {
	greeting string
	name     string
}

// The receiver is passed by value
// If we want to pass by reference, we can use a pointer
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func (g *greeter) setName(name string) {
	g.name = name
}

func main() {
	g := greeter{"Hello", "GO!"}
	g.greet()
	g.setName("Gopher")
	g.greet()
}
