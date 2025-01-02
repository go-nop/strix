package strix

import (
	"fmt"
)

// toInt is a function that converts a string to uint.
func toInt(s string) (int, error) {
	var result int
	var negative bool
	for i := 0; i < len(s); i++ {
		if i == 0 && s[i] == '-' {
			negative = true
			continue
		}
		if s[i] < '0' || s[i] > '9' {
			return 0, fmt.Errorf("invalid character: %c", s[i])
		}
		result = result*10 + int(s[i]-'0')
	}
	if negative {
		result = -result
	}
	return result, nil
}

// toLower is a function that converts a rune to lowercase.
func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

// toUpper is a function that converts a rune to uppercase.
func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 32
	}
	return r
}

func splitByte(s string, sep byte) []string {
	n := len(s)
	estimatedParts := 1 + n/2
	result := make([]string, 0, estimatedParts)

	start := 0
	for {
		i := IndexByte(s[start:], sep)
		if i == -1 {
			break
		}
		result = append(result, s[start:start+i])
		start += i + 1
	}
	result = append(result, s[start:])
	return result
}

func splitOptimized(s, sep string) []string {
	if sep == "" {
		return []string{s}
	}

	var result []string
	idx := 0
	for {
		i := Index(s[idx:], sep)
		if i == -1 {
			result = append(result, s[idx:])
			break
		}
		result = append(result, s[idx:idx+i])
		idx += i + len(sep)
	}

	// Preallocate slice size to avoid excessive reallocations
	// Estimate the number of slices to be added.
	estimateSize := len(result)
	if estimateSize > 5 {
		result = make([]string, 0, estimateSize)
	}

	return result
}
