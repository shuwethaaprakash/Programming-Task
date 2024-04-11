package main

import (
	"reflect"
	"sync"
	"testing"

	"github.com/shuwethaaprakash/Programming-Task/internal/helper"
)

func TestIncreaseCount(t *testing.T) {
	tests := []struct {
		name   string
		input  map[string]int
		key    string
		result map[string]int
	}{
		{
			name:   "Increase count existing IP",
			input:  map[string]int{"168.41.191.40": 1, "72.44.32.10": 2},
			key:    "168.41.191.40",
			result: map[string]int{"168.41.191.40": 2, "72.44.32.10": 2},
		},
		{
			name:   "Increase count new IP",
			input:  map[string]int{"168.41.191.40": 1, "72.44.32.10": 2},
			key:    "177.71.128.21",
			result: map[string]int{"168.41.191.40": 1, "72.44.32.10": 2, "177.71.128.21": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mutex := &sync.Mutex{}
			helper.IncreaseCount(tt.input, tt.key, mutex)
			for k, v := range tt.result {
				if tt.input[k] != v {
					t.Errorf("Expected %s: %d, got %s: %d", k, v, k, tt.input[k])
				}
			}
		})
	}
}

func TestGetTopThree(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []string
	}{
		{
			name:     "Test with 3 items",
			input:    map[string]int{"a": 3, "b": 2, "c": 1},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "Test with more than 3 items",
			input:    map[string]int{"a": 5, "b": 3, "c": 2, "d": 4, "e": 1},
			expected: []string{"a", "d", "b"},
		},
		{
			name:     "Test with less than 3 items",
			input:    map[string]int{"a": 5, "b": 3},
			expected: []string{"a", "b"},
		},
		{
			name:     "Test with empty map",
			input:    map[string]int{},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := helper.GetTopThree(tt.input)

			if len(tt.input) == 0 && len(result) == 0 {
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
