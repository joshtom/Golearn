package main

import (
	"fmt"
	"time"
)

func mains() {
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println("Hello, Go!", year)
}
