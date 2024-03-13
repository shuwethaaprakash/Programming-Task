package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
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

	// Read the file line by line ==> load concurrently AND in chunks to improve speed
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Extract IP and URL
		slice := strings.Fields(line)
		ip := slice[0]
		url := slice[6]

		// Update IP count
		ipCount[ip]++

		// Update URL count
		urlCount[url]++

		// Update active IP count
		activeIPs[ip]++
	}

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
