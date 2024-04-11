package main

import (
	"fmt"
	"os"

	"github.com/shuwethaaprakash/Programming-Task/internal/file"
	"github.com/shuwethaaprakash/Programming-Task/internal/helper"
	"github.com/shuwethaaprakash/Programming-Task/internal/log"
)

const chunkSize = 500

func main() {
	// Check that a log file is given
	if len(os.Args) < 2 {
		fmt.Println("Too few arguments.\n Usage: go run main.go <log_file.txt>")
		return
	}

	// Open given log file
	filePath := os.Args[1]
	lines, readErr := file.ReadLines(filePath)
	if readErr != nil {
		fmt.Println("Error reading log file:", readErr)
		return
	}

	// Process log lines
	uniqueIPs, urlCount, activeIPs, processErr := log.ProcessChunks(lines, chunkSize)
	if processErr != nil {
		fmt.Println("Error processing log lines:", processErr)
		return
	}

	// Find top 3 most visited URLs
	topURLs := helper.GetTopThree(urlCount)

	// Find top 3 most active IP addresses
	topIPs := helper.GetTopThree(activeIPs)

	// Print the results
	helper.PrintResults(uniqueIPs, urlCount, activeIPs, topURLs, topIPs)
}
