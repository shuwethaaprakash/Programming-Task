package file

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filePath string, batchSize int) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	var batches [][]string
	scanner := bufio.NewScanner(file)
	batch := make([]string, 0, batchSize)

	for scanner.Scan() {
		line := scanner.Text()
		batch = append(batch, line)
		if len(batch) == batchSize {
			batches = append(batches, batch)
			batch = make([]string, 0, batchSize)
		}
	}

	if scanErr := scanner.Err(); err != nil {
		return nil, scanErr
	}

	// Append any remaining lines
	if len(batch) > 0 {
		batches = append(batches, batch)
	}

	return batches, nil
}
