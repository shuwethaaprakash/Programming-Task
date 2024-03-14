package main

import (
	"reflect"
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
		{
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
		},
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
