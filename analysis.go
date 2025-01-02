package strix

// CharFrequency is a function that counts the frequency of each character in a string.
func CharFrequency(s string) map[rune]int {
	m := make(map[rune]int, len(s))
	for _, r := range s {
		m[r]++
	}
	return m
}

// CharFrequencyCaseInsensitive is a function that counts the frequency of each character in a string case-insensitively.
func CharFrequencyCaseInsensitive(s string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range s {
		m[toLower(r)]++
	}
	return m
}

// CharFrequencyAlpha is a function that counts the frequency of each alphabetic character in a string.
func CharFrequencyAlpha(s string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range s {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			m[r]++
		}
	}
	return m
}

// CharFrequencyNumeric is a function that counts the frequency of each numeric character in a string.
func CharFrequencyNumeric(s string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range s {
		if r >= '0' && r <= '9' {
			m[r]++
		}
	}
	return m
}

// HasPrefix is a function that checks if a string has a prefix.
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// HasSuffix is a function that checks if a string has a suffix.
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// Index finds the index of the first occurrence of substr in s using the Knuth-Morris-Pratt algorithm.
func Index(s, substr string) int {
	if len(substr) == 0 {
		return 0
	}
	n, m := len(s), len(substr)
	if m > n {
		return -1
	}

	for i := 0; i <= n-m; i++ {
		if s[i] == substr[0] {
			match := true
			for j := 1; j < m; j++ {
				if s[i+j] != substr[j] {
					match = false
					break
				}
			}
			if match {
				return i
			}
		}
	}
	return -1
}

// IndexAny is a function that finds the index of the first occurrence of any character in chars in s.
func IndexAny(s, chars string) int {
	for i, r := range s {
		for _, c := range chars {
			if r == c {
				return i
			}
		}
	}
	return -1
}

// IndexByte is a function that finds the index of the first occurrence of c in s.
func IndexByte(s string, c byte) int {
	for i, b := range []byte(s) {
		if b == c {
			return i
		}
	}
	return -1
}

// IndexFunc is a function that finds the index of the first occurrence of a character that satisfies the function f in s.
func IndexFunc(s string, f func(rune) bool) int {
	for i := 0; i < len(s); i++ {
		if f(rune(s[i])) {
			return i
		}
	}
	return -1
}

// IndexRune is a function that finds the index of the first occurrence of r in s.
func IndexRune(s string, r rune) int {
	for i := 0; i < len(s); i++ {
		if rune(s[i]) == r {
			return i
		}
	}
	return -1
}

// LastIndex is a function that finds the index of the last occurrence of substr in s.
func LastIndex(s, substr string) int {
	if len(substr) == 0 {
		return len(s)
	}
	for i := len(s) - len(substr); i >= 0; i-- {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

// LastIndexAny is a function that finds the index of the last occurrence of any character in chars in s.
func LastIndexAny(s, chars string) int {
	for i := len(s) - 1; i >= 0; i-- {
		for _, c := range chars {
			if rune(s[i]) == c {
				return i
			}
		}
	}
	return -1
}

// LastIndexByte is a function that finds the index of the last occurrence of c in s.
func LastIndexByte(s string, c byte) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			return i
		}
	}
	return -1
}

// LastIndexFunc is a function that finds the index of the last occurrence of a character that satisfies the function f in s.
func LastIndexFunc(s string, f func(rune) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if f(rune(s[i])) {
			return i
		}
	}
	return -1
}

// LastIndexRune is a function that finds the index of the last occurrence of r in s.
func LastIndexRune(s string, r rune) int {
	for i := len(s) - 1; i >= 0; i-- {
		if rune(s[i]) == r {
			return i
		}
	}
	return -1
}
