package main

import "fmt"

func main() {
	students := map[string]int{
		"Alice": 90,
		"Bob":   85,
		"Carol": 92,
	}

	// Add/Update
	students["David"] = 88
	students["Alice"] = 95

	// Check and read
	if grade, ok := students["Bob"]; ok {
		fmt.Printf("Bob's grade: %d\n", grade)
	}

	delete(students, "Carol")

	// Iterate
	fmt.Println("\nAll students:")

	for name, grade := range students {
		fmt.Printf("%s: %d\n", name, grade)
	}

	// Length
	fmt.Printf("\nTotal students: %d\n", len(students))
}
