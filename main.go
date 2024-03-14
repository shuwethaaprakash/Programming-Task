package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
)

type Request struct {
	IP   string
	DATE string
	URL  string
}

func main() {
	// Check that a log file is given
	if len(os.Args) < 2 {
		fmt.Println("Too few arguments.\n Usage: go run main.go <log_file.txt>")
		return
	}

	// Open given log file
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Initialise maps for counting
	ipCount := make(map[string]int)
	urlCount := make(map[string]int)
	activeIPs := make(map[string]int)

	// Read the file line by line
	scanner := bufio.NewScanner(file)

	// Create a wait group to process the lines concurrently
	var wg sync.WaitGroup

	// Initialise mutex to avoid data race condition
	var mutex = sync.RWMutex{}

	for scanner.Scan() {
		wg.Add(1)
		go func(line string) {
			defer wg.Done()

			// Extract IP and URL
			slice := strings.Fields(line)

			if len(slice) < 7 {
				return
			}

			ip := slice[0]
			url := slice[6]

			// Update IP count
			mutex.Lock()
			ipCount[ip]++
			mutex.Unlock()

			// Update URL count
			mutex.Lock()
			urlCount[url]++
			mutex.Unlock()

			// Update active IP count
			mutex.Lock()
			activeIPs[ip]++
			mutex.Unlock()

		}(scanner.Text())

	}

	// Wait for goroutines to complete
	wg.Wait()

	// Calculate number of unique IPs
	uniqueIPs := len(ipCount)

	// Find top 3 most visited URLs
	var topURLs []string
	for url := range urlCount {
		topURLs = append(topURLs, url)
	}
	sort.SliceStable(topURLs, func(i, j int) bool { //use sort stable func instead
		return urlCount[topURLs[i]] > urlCount[topURLs[j]]
	})

	if len(topURLs) > 3 {
		topURLs = topURLs[:3]
	}

	// Find top 3 most active IP addresses
	var topIPs []string
	for ip := range activeIPs {
		topIPs = append(topIPs, ip)
	}
	sort.SliceStable(topIPs, func(i, j int) bool {
		return activeIPs[topIPs[i]] > activeIPs[topIPs[j]]
	})

	if len(topIPs) > 3 {
		topIPs = topIPs[:3]
	}

	// Print the results
	fmt.Printf("Number of unique IP addresses: %d\n", uniqueIPs)
	fmt.Println("Top 3 most visited URLs:")
	for i, url := range topURLs {
		fmt.Printf("%d. %s (%d vists)\n", i+1, url, urlCount[url])
	}
	fmt.Println("Top 3 most active IP addresses:")
	for i, ip := range topIPs {
		fmt.Printf("%d. %s (%d requests)\n", i+1, ip, activeIPs[ip])
	}
}
