package log

import (
	"fmt"
	"strings"
	"sync"

	"github.com/shuwethaaprakash/Programming-Task/internal/count"
	"github.com/shuwethaaprakash/Programming-Task/internal/validate"
)

func ProcessChunks(batches [][]string) (int, *count.SafeMap, *count.SafeMap, error) {
	// Initialise maps for counting
	ipCount := count.SafeMap{}
	urlCount := count.SafeMap{}
	activeIPs := count.SafeMap{}

	// Create a wait group to process the lines concurrently
	var wg sync.WaitGroup

	// Process log lines in batches
	for _, batch := range batches {
		wg.Add(1)
		go func(batch []string) {
			defer wg.Done()
			for _, line := range batch {
				processLine(line, &ipCount, &urlCount, &activeIPs)
			}
		}(batch)
	}

	// Wait for goroutines to complete
	wg.Wait()

	// Calculate number of unique IPs
	uniqueIPs := count.GetMapLength(&ipCount)

	return uniqueIPs, &urlCount, &activeIPs, nil
}

func processLine(line string, ipCount, urlCount, activeIPs *count.SafeMap) {

	// Check if line is in a valid format
	if !validate.IsValidRequest(line) {
		fmt.Printf("Error: Request %q is not in the correct format\n", line)
		return
	}

	// Extract IP and URL
	slice := strings.Fields(line)

	if len(slice) < 7 {
		fmt.Println("Error reading log file: request incorrect length")
		return
	}

	ip := slice[0]
	url := slice[6]

	// Update IP count
	count.IncreaseCount(ipCount, ip)

	// Update URL count
	count.IncreaseCount(urlCount, url)

	// Update active IP count
	count.IncreaseCount(activeIPs, ip)
}
