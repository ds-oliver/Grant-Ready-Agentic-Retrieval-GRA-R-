package security

import "regexp"

var (
	emailRegex = regexp.MustCompile(`(?i)[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}`)
	ssnRegex   = regexp.MustCompile(`\b\d{3}-\d{2}-\d{4}\b`)
)

// Sanitize redacts common PII patterns such as emails and SSN-like values.
func Sanitize(input string) string {
	redacted := emailRegex.ReplaceAllString(input, "[REDACTED]")
	redacted = ssnRegex.ReplaceAllString(redacted, "[REDACTED]")
	return redacted
}
