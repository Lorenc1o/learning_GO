package main

import "fmt"

func main() {
	// Map
	statePopulations := map[string]int{ // we have to specify the type of the key and the value
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations)

	// make function
	statePopulations2 := make(map[string]int)
	statePopulations2["California"] = 39250017 // To add a key-value pair, we use the key in square brackets
	statePopulations2["Texas"] = 27862596
	statePopulations2["Florida"] = 20612439
	statePopulations2["New York"] = 19745289
	statePopulations2["Pennsylvania"] = 12802503
	statePopulations2["Illinois"] = 12801539
	statePopulations2["Ohio"] = 11614373
	fmt.Println(statePopulations2)

	// To get the value of a key, we use the key in square brackets
	fmt.Println(statePopulations["California"])
	// The return order of the keys is not guaranteed

	// To delete a key, we use the delete function
	delete(statePopulations, "Ohio")
	fmt.Println(statePopulations)

	// To check if a key exists, we use the value, ok idiom
	pop, ok := statePopulations["Ohio"]
	fmt.Println(pop, ok) // 0 false
	pop, ok = statePopulations["California"]
	fmt.Println(pop, ok) // 39250017 true

	// To get the length of a map, we use the len function
	fmt.Println("Length:", len(statePopulations))

	// When you copy a map, you copy the reference to the map
	statePopulations3 := statePopulations
	statePopulations3["Ohio"] = 8
	fmt.Println(statePopulations)
	fmt.Println(statePopulations3)
}
