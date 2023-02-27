package main

import "fmt"

func main() {
	grade1 := 97
	grade2 := 85
	grade3 := 93
	fmt.Printf("Grades: %v, %v, %v\n", grade1, grade2, grade3)

	// Array
	var grades [3]int = [3]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Grades: %v, %v, %v\n", grades[0], grades[1], grades[2])
	// Arrays are stored in contiguous memory locations

	// Array length is part of the type
	var students [3]string = [3]string{"Lisa", "Ahmed", "Samantha"}
	fmt.Printf("Number of students: %v\n", len(students))

	// Array literal
	grades2 := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades2)

	// Array index
	fmt.Printf("First student: %v\n", students[0])
	fmt.Printf("Second student: %v\n", students[1])
	fmt.Printf("Third student: %v\n", students[2])

	// Array index out of bounds
	// fmt.Printf("Fourth student: %v\n", students[3]) // panic: runtime error: index out of range [3] with length 3

	// Matrix
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Printf("Identity matrix: %v\n", identityMatrix)

	// Copying arrays is done by value
	var copiedArray [3]int = grades
	copiedArray[0] = 100
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Copied array: %v\n", copiedArray)

	// If we want to copy the array by reference, we need to use & to get the address of the array
	var copiedArray2 *[3]int = &grades
	copiedArray2[0] = 100
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Copied array: %v\n", copiedArray2)
}
