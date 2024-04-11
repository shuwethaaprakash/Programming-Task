package log

import (
	"fmt"
	"strings"
	"sync"

	"github.com/shuwethaaprakash/Programming-Task/internal/helper"
)

func ProcessChunks(lines []string, chunkSize int) (int, map[string]int, map[string]int, error) {
	// Initialise maps for counting
	ipCount := make(map[string]int)
	urlCount := make(map[string]int)
	activeIPs := make(map[string]int)

	// Create a wait group to process the lines concurrently
	var wg sync.WaitGroup

	// Create individual mutex
	var ipMutex sync.Mutex
	var urlMutex sync.Mutex
	var activeIPMutex sync.Mutex

	// Process log lines in chunks
	for i := 0; i < len(lines); i += chunkSize {
		end := i + chunkSize
		if end > len(lines) {
			end = len(lines)
		}
		chunk := lines[i:end]

		// Process chunk of log lines concurrently
		for _, line := range chunk {
			lineCopy := line
			wg.Add(1)
			go func(line string) {
				defer wg.Done()
				processLine(line, ipCount, urlCount, activeIPs, &ipMutex, &urlMutex, &activeIPMutex)
			}(lineCopy)
		}
	}

	// Wait for goroutines to complete
	wg.Wait()

	// Calculate number of unique IPs
	uniqueIPs := len(ipCount)

	return uniqueIPs, urlCount, activeIPs, nil
}

func processLine(line string, ipCount, urlCount, activeIPs map[string]int, ipMutex, urlMutex, activeIPMutex *sync.Mutex) {

	// Check if line is in a valid format
	if !helper.IsValidRequest(line) {
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
	helper.IncreaseCount(ipCount, ip, ipMutex)

	// Update URL count
	helper.IncreaseCount(urlCount, url, urlMutex)

	// Update active IP count
	helper.IncreaseCount(activeIPs, ip, activeIPMutex)
}
