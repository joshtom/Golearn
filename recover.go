package main

import (
	"fmt"
	"log"
)

// Simulates processing a user request that might panic
func processRequest(userID int) {
	if userID == 0 {
		panic("Invalid user ID: cannot be zero")
	}
	if userID < 0 {
		panic(fmt.Sprintf("invalid user ID: %d is negative", userID))
	}
	fmt.Printf("Successfully processed request for user ID: %d\n", userID)
}

func safeHandler(userID int) {
	defer func() {
		if r := recover(); r != nil {
			//  Log the error but don't crash the server
			log.Printf("Recovered from panic: %v", r)
			fmt.Printf("Request for user %d failed gracefully \n", userID)
		}
	}()
	processRequest(userID)
}

func recoverMain() {
	fmt.Println("Server Started")

	userIDs := []int{42, 0, 15, -5, 100}

	for _, id := range userIDs {
		fmt.Printf("Handling request for user %d \n", id)
		safeHandler(id)
		fmt.Println()
	}

	fmt.Println("Server Still Running")
}
