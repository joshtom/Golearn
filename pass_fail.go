package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main2() {
	fmt.Println("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err) // Report the error and stop the program
	}
	fmt.Println(input)
}
