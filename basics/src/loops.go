package main

import (
	"fmt"
)

func main() {
	// A for statement is a way to repeat a block of code
	// It has three components: init statement, condition expression and post statement
	// The init statement is executed before the first iteration
	// The condition expression is evaluated before every iteration
	// If the condition expression is true, the block of code is executed
	// The post statement is executed at the end of every iteration
	// The init statement and post statement are optional
	// The condition expression is mandatory
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// If the init statement and post statement are omitted, we have a while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// If the condition expression is omitted, we have an infinite loop
	// We can break out of the loop using the break keyword
	for {
		fmt.Println("This will print forever")
		break
	}

	// We can skip the current iteration using the continue keyword
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	// We can use a for loop to iterate over an array
	nums := []int{2, 3, 4}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	// We can use a for loop to iterate over a slice
	nums2 := []int{2, 3, 4}
	for i, num := range nums2 {
		fmt.Println(i, num)
	}

	// We can use a for loop to iterate over a map
	emails := map[string]string{"Bob": "Plumber", "Alice": "Teacher"}
	for name, job := range emails {
		fmt.Println(name, job)
	}

	// We can use labels to break out of nested loops
	for i := 0; i < 5; i++ {
	Loop:
		for j := 0; j < 5; j++ {
			if i == 3 {
				break Loop
			}
			fmt.Println(i, j)
		}
	}
}
