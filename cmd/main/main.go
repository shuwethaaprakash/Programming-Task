package main

import (
	"fmt"
	"os"

	"github.com/shuwethaaprakash/Programming-Task/internal/count"
	"github.com/shuwethaaprakash/Programming-Task/internal/file"
	"github.com/shuwethaaprakash/Programming-Task/internal/log"
	"github.com/shuwethaaprakash/Programming-Task/internal/printing"
)

const batchSize = 500

func main() {
	// Check that a log file is given
	if len(os.Args) < 2 {
		fmt.Println("Too few arguments.\n Usage: go run main.go <log_file.txt>")
		return
	}

	// Open given log file
	filePath := os.Args[1]
	batches, readErr := file.ReadLines(filePath, batchSize)
	if readErr != nil {
		fmt.Println("Error reading log file:", readErr)
		return
	}

	// Process log lines
	uniqueIPs, urlCount, activeIPs, processErr := log.ProcessChunks(batches)
	if processErr != nil {
		fmt.Println("Error processing log lines:", processErr)
		return
	}

	// Find top 3 most visited URLs
	topURLs := count.GetTopThree(urlCount)

	// Find top 3 most active IP addresses
	topIPs := count.GetTopThree(activeIPs)

	// Print the results
	printing.PrintResults(uniqueIPs, urlCount, activeIPs, topURLs, topIPs)
}
