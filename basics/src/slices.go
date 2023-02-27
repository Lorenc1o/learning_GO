package main

import "fmt"

func main() {
	// Slice
	grades := []int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	// The difference between an array and a slice is that a slice does not have a fixed length
	// A slice is a descriptor for a contiguous segment of an underlying array
	// and provides access to a numbered sequence of elements from that array
	// A slice is a three word data structure: the pointer, the length, and the capacity
	// - The length is the number of elements referred to by the slice
	// - The capacity is the number of elements in the underlying array, counting from the first element in the slice
	// - The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s)
	// - The zero value of a slice is nil
	// - A slice can be re-sliced, creating a new slice value that points to the same array
	// - A slice expression s[lo:hi] creates a slice with low index lo and high index hi-1, thus excluding the element at index hi
	grades2 := grades[1:2]
	fmt.Printf("Grades: %v\n", grades2)

	// Slice literal
	grades3 := []int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades3)

	// When you copy a slice, you copy the descriptor, so both point to the same array
	// If you modify the elements of one slice, you will see the changes in the other slice as well
	grades4 := grades3
	grades4[0] = 100
	fmt.Printf("Grades: %v\n", grades3)
	fmt.Printf("Grades: %v\n", grades4)

	// make function
	grades5 := make([]int, 3) // len = 3, cap = 3
	fmt.Printf("Grades: %v\n", grades5)
	grades6 := make([]int, 3, 5) // len = 3, cap = 5
	fmt.Printf("Grades: %v\n", grades6)

	// append function
	grades5 = append(grades5, 100)
	fmt.Printf("Grades: %v\n", grades5)

	// append can be used to append multiple elements
	grades5 = append(grades5, 100, 99, 98)
	fmt.Printf("Grades: %v\n", grades5)

	// append can be used to append a slice to another slice
	grades5 = append(grades5, grades6...) // ... is the spread operator
	fmt.Printf("Grades: %v\n", grades5)

	// shift
	grades8 := grades5[1:] // remove the first element
	fmt.Printf("Grades: %v\n", grades8)
	grades8 = grades5[:len(grades5)-1] // remove the last element
	fmt.Printf("Grades: %v\n", grades8)
	// to remove an element at a specific index, we need to use append
	grades8 = append(grades8[:3], grades8[4:]...) // remove the element at index 3
	fmt.Printf("Grades: %v\n", grades8)
	// This does not work because the slices are pointing to the same array
}
