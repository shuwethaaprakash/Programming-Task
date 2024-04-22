package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/shuwethaaprakash/Programming-Task/internal/file"
)

func TestReadLines(t *testing.T) {
	tests := []struct {
		name        string
		filePath    string
		batchSize   int
		expected    [][]string
		expectedErr error
	}{
		{
			name:      "Valid file with batchSize 2",
			filePath:  "data/test_data.txt",
			batchSize: 2,
			expected: [][]string{
				{"line 1", "line 2"},
				{"line 3", "line 4"},
				{"line 5"},
			},
			expectedErr: nil,
		},
		{
			name:        "Invalid file path",
			filePath:    "invalid_path.txt",
			batchSize:   2,
			expected:    nil,
			expectedErr: fmt.Errorf("open invalid_path.txt: no such file or directory"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			batches, err := file.ReadLines(tt.filePath, tt.batchSize)

			// Check for error
			if (err == nil && tt.expectedErr != nil) || (err != nil && tt.expectedErr == nil) {
				t.Fatalf("Expected error: %v, got: %v", tt.expectedErr, err)
			}
			if err != nil && err.Error() != tt.expectedErr.Error() {
				t.Fatalf("Expected error: %v, got: %v", tt.expectedErr, err)
			}

			// Check batches
			if !reflect.DeepEqual(batches, tt.expected) {
				t.Errorf("Expected batches: %v, got: %v", tt.expected, batches)
			}
		})
	}
}
