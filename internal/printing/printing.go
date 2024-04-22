package printing

import (
	"fmt"
	"github.com/shuwethaaprakash/Programming-Task/internal/count"
)

func PrintResults(numIPs int, urls, IPs *count.SafeMap, topURLs []string, topIPs []string) {
	fmt.Printf("Number of unique IP addresses: %d\n", numIPs)
	fmt.Println("Top 3 most visited URLs:")
	for i, url := range topURLs {
		if key, ok := urls.Load(url); ok {
			fmt.Printf("%d. %s (%d vists)\n", i+1, url, key)
		} else {
			fmt.Printf("Error loading key from map %v", urls)
		}
	}
	fmt.Println("Top 3 most active IP addresses:")
	for i, ip := range topIPs {
		if key, ok := IPs.Load(ip); ok {
			fmt.Printf("%d. %s (%d requests)\n", i+1, ip, key)
		} else {
			fmt.Printf("Error loading key from map %v", IPs)
		}
	}
}
