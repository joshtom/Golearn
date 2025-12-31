package main

import (
	"fmt"
	"log"
)

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a height of %.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil
}

func main() {
	// var amount, total float64
	amount, err := paintNeeded(4.2, -3.0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(err)
	fmt.Printf("%.2f liters needed\n", amount)
}

// Next concept to learn
// - Arrays and slices
// - Maps
// - Structs
// - Defined Types
// - Encapsulation and Embedding
// - Interfaces
// - Recovering from failures
// - Goroutines and Channels
// - Code quality Assurance - Automated Testing
// - Responding to requests: Web apps
// - A pattern to follow: HTML Templates
// - Understanding os.OpenFile: Opening files
