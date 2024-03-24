package main

import (
	"os/exec"
	"reflect"
	"strings"
	"sync"
	"testing"
)

func TestIncreaseCount(t *testing.T) {
	type args struct {
		IP string
	}
	tests := []struct {
		name     string
		arg      args
		expected int
	}{
		{"First IP count", args{"192.168.1.1"}, 1},
		//{"Additional IP count", args{"192.168.1.1"}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := make(map[string]int)
			var mutex sync.Mutex

			increaseCount(&m, tt.arg.IP, &mutex)
			if m[tt.arg.IP] != tt.expected {
				t.Errorf("Error: increaseCount() for %s, want: %d, got: %d", tt.arg.IP, tt.expected, m[tt.arg.IP])
			}
		})
	}
}

func TestGetTopThree(t *testing.T) {
	type args struct {
		ipMap map[string]int
	}
	tests := []struct {
		name     string
		arg      args
		expected []string
	}{
		{
			name: "Placing addresses in descending order",
			arg: args{
				ipMap: map[string]int{
					"google.com":   10,
					"test.com":     100,
					"facebook.com": 1,
					"youtube.com":  1000,
				},
			},
			expected: []string{"youtube.com", "test.com", "google.com"},
		},
		/* {
			name: "Addresses with the same number of requests",
			arg: args{
				ipMap: map[string]int{
					"google.com":   1,
					"test.com":     1,
					"youtube.com":  0,
					"facebook.com": 1,
				},
			},
			expected: []string{"google.com", "test.com", "facebook.com"},
		},
		{
			name: "Keeping equal elements in original order",
			arg: args{
				ipMap: map[string]int{
					"google.com":   20,
					"test.com":     10,
					"youtube.com":  20,
					"facebook.com": 1,
				},
			},
			expected: []string{"google.com", "youtube.com", "test.com"},
		}, */
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := getTopThree(&tt.arg.ipMap)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Error: getTopThree() want: %v, got: %v", tt.expected, res)
			}
		})
	}
}

// Note attempted to test integration, however cannot compare due to utilising maps when iterating
func TestIntegration(t *testing.T) {
	tests := []struct {
		name          string
		fileName      string
		expectedLines []string
	}{
		{
			name:     "Provided log file",
			fileName: "programming-task-example-data.log",
			expectedLines: []string{
				"Number of unique IP addresses: 11",
				"Top 3 most visited URLs:",
				"1. /docs/manage-websites/ (2 visits)",
				"2. /asset.css (1 visits)",
				"3. /blog/category/community/ (1 visits)",
				"Top 3 most active IP addresses:",
				"1. 168.41.191.40 (4 requests)",
				"2. 177.71.128.21 (3 requests)",
				"3. 72.44.32.10 (3 requests)",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Create a command to execute the function
			cmd := exec.Command("go", "run", "main.go", tt.fileName)
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Errorf("Error: could not read file, expected %v as output.", tt.expectedLines)
			}

			// Convert output to string
			res := strings.Split(strings.TrimSpace(string(output)), "\n")

			// Compare the result to expected output
			for i, line := range res {
				// Cannot compare lines that have the same number of requests/visits
				if i == 3 || i == 4 || i == 7 || i == 8 {
					continue
				}
				if line != tt.expectedLines[i] {
					t.Errorf("Error: output does not match expected result. Wanted: %s, got: %s", tt.expectedLines[i], line)
				}
			}

		})
	}
}
