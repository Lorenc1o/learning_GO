package main

import "fmt"

// Go does not have inheritance
// Instead, we can compose structs
type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal   // We can embed the Animal struct in the Bird struct
	SpeedKPH float32
	CanFly   bool
}

func main() {
	b := Bird{}
	b.Name = "Emu" // We can access the embedded fields using the dot notation
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)
}
