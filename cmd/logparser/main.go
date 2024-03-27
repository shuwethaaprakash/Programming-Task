package logparser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/shuwethaaprakash/Programming-Task/internal/helper"
)

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
		fmt.Println("Cannot open log file: ", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
			return
		}
	}(file)

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
			lineCopy := line
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
				helper.IncreaseCount(ipCount, ip, &ipMutex)

				// Update URL count
				helper.IncreaseCount(urlCount, url, &urlMutex)

				// Update active IP count
				helper.IncreaseCount(activeIPs, ip, &activeIPMutex)

			}(lineCopy)
		}

	}

	// Wait for goroutines to complete
	wg.Wait()

	// Calculate number of unique IPs
	uniqueIPs := len(ipCount)

	// Find top 3 most visited URLs
	topURLs := helper.GetTopThree(urlCount)

	// Find top 3 most active IP addresses
	topIPs := helper.GetTopThree(activeIPs)

	// Print the results
	helper.PrintResults(uniqueIPs, urlCount, activeIPs, topURLs, topIPs)
}
