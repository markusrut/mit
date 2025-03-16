package utils

import (
	"fmt"
	"strings"
)

// ParseInput takes a string input and splits it into a slice of strings based on the provided delimiter.
func ParseInput(input string, delimiter string) []string {
	return strings.Split(input, delimiter)
}

// FormatOutput takes a string and returns it formatted with a prefix.
func FormatOutput(output string) string {
	return fmt.Sprintf("Output: %s", output)
}

// ```
// root
// ├── file1.go
// ├── folder1
// │   └── file1.go
// ├── folder2
// │   └── folder1
// │       └── file1.go
// ├── file2.mod
// ├── file3.sum
// └── file4.md
// ```
