package helper

import (
	"fmt"
	"sort"
	"sync"
)

func IncreaseCount(m map[string]int, key string, mutex *sync.Mutex) {
	mutex.Lock()
	(m)[key]++
	mutex.Unlock()
}

func GetTopThree(m map[string]int) []string {
	// Create an array to hold the top 3 items
	var topThree []string

	for url := range m {
		topThree = append(topThree, url)
	}

	// Sort items in descending order (use sort stable func instead)
	sort.SliceStable(topThree, func(i, j int) bool {
		return (m)[topThree[i]] > (m)[topThree[j]]
	})

	// Get the top 3 items
	if len(topThree) > 3 {
		topThree = topThree[:3]
	}

	return topThree
}

func PrintResults(numIPs int, urls map[string]int, IPs map[string]int, topURLs []string, topIPs []string) {
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
