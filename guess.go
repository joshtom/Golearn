package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main4() {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // Create a local random generator
	target := r.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin) // Create a budio reader which lets us read keyboard input
	success := false

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left")
		fmt.Println("Make a guess: ")

		input, err := reader.ReadString('\n') // Read what the user types up until they press enter

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)  // Remove the new line
		guess, err := strconv.Atoi(input) // Convert the input string to an integer

		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops. Your guess was LOW")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			break // Exit the loop
		}

		if !success {
			fmt.Println("Sorry, you didn't guess my number. It was: ", target)
		}
	}
}
