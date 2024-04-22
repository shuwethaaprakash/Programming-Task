package test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/shuwethaaprakash/Programming-Task/internal/count"
	"github.com/shuwethaaprakash/Programming-Task/internal/printing"
)

func TestPrintResults(t *testing.T) {
	// Create URLs ands IPs test data
	numIPs := 3

	urls := &count.SafeMap{}
	urls.Store("https://test.com/page1", 10)
	urls.Store("https://test.com/", 8)
	urls.Store("https://test.com/home", 5)

	IPs := &count.SafeMap{}
	IPs.Store("192.168.0.1", 20)
	IPs.Store("10.0.0.1", 15)
	IPs.Store("172.0.0.1", 10)

	topURLs := []string{"https://test.com/page1", "https://test.com/", "https://test.com/home"}
	topIPs := []string{"192.168.0.1", "10.0.0.1", "172.0.0.1"}

	// Redirect stdout to buffer
	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w
	defer func() {
		os.Stdout = old
		err := w.Close()
		if err != nil {
			return
		}
	}()

	// Call Results function
	printing.PrintResults(numIPs, urls, IPs, topURLs, topIPs)

	// Get output
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		return
	}
	closeErr := r.Close()
	if closeErr != nil {
		return
	}

	// Define expected output
	expectedOutput := fmt.Sprintf("Number of unique IP addresses: %d\n", numIPs) +
		"Top 3 most visited URLs:\n" +
		"1. https://test.com/page1 (10 visits)\n" +
		"2. https://test.com/ (8 visits)\n" +
		"3. https://test.com/home (5 visits)\n" +
		"Top 3 most active IP addresses:\n" +
		"1. 192.168.0.1 (20 requests)\n" +
		"2. 10.0.0.1 (15 requests)\n" +
		"3. 172.0.0.1 (10 requests)\n"

	// Compare output
	if buf.String() != expectedOutput {
		t.Errorf("Unexpected output:\n%s", buf.String())
	}
}
