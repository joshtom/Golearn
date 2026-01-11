package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Result holds the fetch result from a URL
type Result struct {
	URL      string
	Status   string
	Duration time.Duration
}

// fetchURL simulates fetching data from a URL
func fetchURL(url string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()

	// Simulate network latency (100-500ms)
	delay := time.Duration(100+rand.Intn(400)) * time.Millisecond
	time.Sleep(delay)

	// Send result back through channel
	results <- Result{
		URL:      url,
		Status:   "200 OK",
		Duration: time.Since(start),
	}
}

func mainGoRoutine() {
	urls := []string{
		"https://api.github.com",
		"https://api.twitter.com",
		"https://api.facebook.com",
		"https://api.google.com",
		"https://api.amazon.com",
	}

	// Create a channel to receive results

	results := make(chan Result, len(urls))

	// WaitGroup to wait for all foroutines to finish
	var wg sync.WaitGroup

	fmt.Println("Starting concurrent fetches...")
	start := time.Now()

	// Launch a goroutine for each URL
	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, results, &wg)
	}

	// Close results channel when all goroutines complete
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("✓ %s - %s (%v)\n", result.URL, result.Status, result.Duration)
	}

	fmt.Printf("\nAll fetches completed in %v\n", time.Since(start))
	fmt.Println("(Sequential would have taken ~1.5-2.5 seconds)")
}
