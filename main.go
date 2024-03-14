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
	IP  string
	URL string
}

const chunkSize = 500

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

	// Create individual mutex
	var ipMutex sync.Mutex
	var urlMutex sync.Mutex
	var activeIPMutex sync.Mutex

	for {
		// Create a "chunk" to read
		lines := make([]string, 0, chunkSize)
		for i := 0; i < chunkSize && scanner.Scan(); i++ {
			lines = append(lines, scanner.Text())
		}

		if len(lines) == 0 {
			break
		}

		for _, line := range lines {
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
				increaseCount(&ipCount, ip, &ipMutex)

				// Update URL count
				increaseCount(&urlCount, url, &urlMutex)

				// Update active IP count
				increaseCount(&activeIPs, ip, &activeIPMutex)

			}(line)
		}

	}

	// Wait for goroutines to complete
	wg.Wait()

	// Calculate number of unique IPs
	uniqueIPs := len(ipCount)

	// Find top 3 most visited URLs
	topURLs := getTopThree(&urlCount)

	// Find top 3 most active IP addresses
	topIPs := getTopThree(&activeIPs)

	// Print the results
	printResults(uniqueIPs, urlCount, activeIPs, topURLs, topIPs)
}

func increaseCount(m *map[string]int, key string, mutex *sync.Mutex) {
	mutex.Lock()
	(*m)[key]++
	mutex.Unlock()
}

func getTopThree(m *map[string]int) []string {
	// Create an array to hold the top 3 items
	var topThree []string

	for url := range *m {
		topThree = append(topThree, url)
	}

	// Sort items in descending order (use sort stable func instead)
	sort.SliceStable(topThree, func(i, j int) bool {
		return (*m)[topThree[i]] > (*m)[topThree[j]]
	})

	// Get the top 3 items
	if len(topThree) > 3 {
		topThree = topThree[:3]
	}

	return topThree
}

func printResults(numIPs int, urls map[string]int, IPs map[string]int, topURLs []string, topIPs []string) {
	fmt.Printf("Number of unique IP addresses: %d\n", numIPs)
	fmt.Println("Top 3 most visited URLs:")
	for i, url := range topURLs {
		fmt.Printf("%d. %s (%d vists)\n", i+1, url, urls[url])
	}
	fmt.Println("Top 3 most active IP addresses:")
	for i, ip := range topIPs {
		fmt.Printf("%d. %s (%d requests)\n", i+1, ip, IPs[ip])
	}
}
