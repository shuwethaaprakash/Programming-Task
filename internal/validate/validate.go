package validate

import "regexp"

func IsValidRequest(req string) bool {
	pattern := `^(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}) - (?:\w+ )?-? ?\[(\d{2}\/\w{3}\/\d{4}:\d{2}:\d{2}:\d{2} [+\-]\d{4})\] "(GET|POST|PUT|DELETE|HEAD|OPTIONS|PATCH|CONNECT|TRACE) (?:https?:\/\/[^\/]+)?\/?(?:[^"]*\/)?[^"]* HTTP\/1\.1" \d+ \d+ "-" "Mozilla\/.*"(?: .*)?$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(req)
}
