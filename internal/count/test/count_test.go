package test

import (
	"reflect"
	"sync"
	"testing"

	"github.com/shuwethaaprakash/Programming-Task/internal/count"
)

func TestIncreaseCount(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		expected int
	}{
		{
			name:     "Increase count new IP",
			key:      "168.41.191.40",
			expected: 1,
		},
		{
			name:     "Increase count existing IP",
			key:      "168.41.191.40",
			expected: 2,
		},
	}

	m := &count.SafeMap{
		Map: sync.Map{},
		Mu:  sync.Mutex{},
	}

	for _, tt := range tests {
		// Call IncreaseCount function
		count.IncreaseCount(m, tt.key)

		// Retrieve the value from the map
		value, _ := m.Load(tt.key)
		res := value.(int)

		// Check if the actual value matches the expected value
		if res != tt.expected {
			t.Errorf("For key %s, expected value %d, but got %d", tt.key, tt.expected, res)
		}
	}
}

func TestGetTopThree(t *testing.T) {
	tests := []struct {
		name     string
		input    count.SafeMap
		expected []string
	}{
		{
			name: "Test with 3 items",
			input: count.SafeMap{
				Map: sync.Map{},
				Mu:  sync.Mutex{},
			},
			expected: []string{"a", "b", "c"},
		},
		{
			name: "Test with more than 3 items",
			input: count.SafeMap{
				Map: sync.Map{},
				Mu:  sync.Mutex{},
			},
			expected: []string{"a", "d", "b"},
		},
		{
			name: "Test with less than 3 items",
			input: count.SafeMap{
				Map: sync.Map{},
				Mu:  sync.Mutex{},
			},
			expected: []string{"b", "a"},
		},
		{
			name: "Test with empty map",
			input: count.SafeMap{
				Map: sync.Map{},
				Mu:  sync.Mutex{},
			},
			expected: []string{},
		},
	}

	// Add input to maps
	tests[0].input.Store("a", 3)
	tests[0].input.Store("b", 2)
	tests[0].input.Store("c", 1)

	tests[1].input.Store("a", 5)
	tests[1].input.Store("b", 3)
	tests[1].input.Store("c", 2)
	tests[1].input.Store("d", 4)
	tests[1].input.Store("e", 1)

	tests[2].input.Store("a", 3)
	tests[2].input.Store("b", 5)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.input.Mu.Lock()
			defer tt.input.Mu.Unlock()

			res := count.GetTopThree(&tt.input)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Test case %q failed: expected %v, got %v", tt.name, tt.expected, res)
			}
		})
	}
}
