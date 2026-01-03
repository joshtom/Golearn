package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)

	if err != nil {
		return 0, err
	}

	return number, nil
}

func tocelcius() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := getFloat()

	if err != nil {
		log.Fatal(err) // If an error is returned, we log it and exit
	}

	celsius := (fahrenheit - 32) * 5 / 9 // Convert temperature to Celsius
	fmt.Printf("%.2f°F is %.2f°C\n", fahrenheit, celsius)
}
