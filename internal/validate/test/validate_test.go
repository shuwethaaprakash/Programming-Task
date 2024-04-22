package test

import (
	"testing"

	"github.com/shuwethaaprakash/Programming-Task/internal/validate"
)

func TestIsValidRequest(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected bool
		wantErr  bool
	}{
		{
			name: "Valid Requests",
			input: []string{
				`127.0.0.1 - - [01/Jan/2022:12:00:00 +0000] "GET /index.html HTTP/1.1" 200 1234 "-" "Mozilla/5.0"`,
				`192.168.0.1 - admin [19/Dec/2021:09:30:00 +0200] "POST /submit-form HTTP/1.1" 201 3575 "-" "Mozilla/5.0"`,
				`72.44.32.10 - - [19/Jul/2023:15:49:48 +0500] "PUT /error/ HTTP/1.1" 400 3574 "-" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/536.5 Chrome/19.0.1084.9 Safari/536.5"`,
				`72.44.32.10 - - [09/Jul/2018:15:48:07 +0200] "TRACE / HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (compatible; MSIE 10.6; Windows NT 6.1; Trident/5.0; InfoPath.2; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET CLR 2.0.50727)" blah blah`,
			},
			expected: true,
		},
		{
			name: "Invalid Requests",
			input: []string{
				`Invalid`,
				`GET /index.html HTTP/1.1`,
				`192.168.0.1 - - [01/Jan/2022:12:00:00 +0000] "INVALID /submit-form HTTP/1.1" 404 1234 "-" "Mozilla/5.0"`,
				`1111 - - [01/Jan/2022:12:00:00 +0000] "GET /index.html HTTP/1.1" 200 1234 "-" "Mozilla/5.0"`,
				`127.0.0.1 [01/Jan/2022:12:00:00 +0000] "POST /index.html HTTP/1.1" 1 1234 "-" "Mozilla/5.0"`,
				`127.0.0.1 - - [01/Jan/2022:12:00:00 +0000] "FAKE /index.html HTTP/1.1" 200 1234 "-" "Mozilla/5.0"`,
				`127.0.0.1 - [01/Jan/2022:12:00:00 +0000] "POST / HTWP/1.1" 200 1234 "-" "Mozilla/5.0"`,
			},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, req := range tt.input {
				got := validate.IsValidRequest(req)
				if got != tt.expected {
					t.Errorf("IsValidRequest(), want = %v, got = %v, request = %v", tt.expected, got, req)
				}
			}
		})
	}
}
