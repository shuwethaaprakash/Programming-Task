package test

import (
	"testing"

	"github.com/shuwethaaprakash/Programming-Task/internal/count"
	"github.com/shuwethaaprakash/Programming-Task/internal/log"
)

func TestProcessChunks(t *testing.T) {
	tests := []struct {
		name           string
		batches        [][]string
		expectedUnique int
		expectedURLs   []string
		expectedIPs    []string
	}{
		{
			name: "Test one batch valid log lines",
			batches: [][]string{
				{`192.168.0.1 - - [10/Apr/2024:09:15:45 +0000] "GET /home HTTP/1.1" 200 1234 "-" "Mozilla/5.0"`,
					`192.168.0.2 - - [10/Apr/2024:09:15:45 +0000] "GET /about HTTP/1.1" 201 3575 "-" "Mozilla/5.0"`},
			},
			expectedUnique: 2,
			expectedURLs:   []string{"/home", "/about"},
			expectedIPs:    []string{"192.168.0.1", "192.168.0.2"},
		},
		{
			name: "Test invalid log lines",
			batches: [][]string{
				{`192.168.0.1 - - [10/Apr/2024:09:15:45 +0000] "GET /home HTTP/1.1" 200 1234 "-" "Mozilla/5.0"`,
					`192.168.0.2 - - [10/Apr/2024:09:15:45 +0000] "INVALID REQUEST" 201 3575 "-" "Mozilla/5.0"`},
			},
			expectedUnique: 1,
			expectedURLs:   []string{"/home"},
			expectedIPs:    []string{"192.168.0.1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call ProcessChunks with the test batches
			uniqueIPs, urlCount, activeIPs, err := log.ProcessChunks(tt.batches)
			if err != nil {
				t.Fatalf("Error processing chunks: %v", err)
			}

			// Check number of unique IPs
			if uniqueIPs != tt.expectedUnique {
				t.Errorf("Expected unique IPs: %d, got: %d", tt.expectedUnique, uniqueIPs)
			}

			// Check top 3 URLs
			topURLs := count.GetTopThree(urlCount)
			if !compareSlices(topURLs, tt.expectedURLs) {
				t.Errorf("Expected top URLs: %v, got: %v", tt.expectedURLs, topURLs)
			}

			// Check top 3 active IPs
			topIPs := count.GetTopThree(activeIPs)
			if !compareSlices(topIPs, tt.expectedIPs) {
				t.Errorf("Expected top IPs: %v, got: %v", tt.expectedIPs, topIPs)
			}
		})
	}
}

func compareSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
