// Description: This file contains the comparison functions.

package strix

// Compare is a function that compares two strings.
func Compare(s1, s2 string) int {
	if s1 < s2 {
		return -1
	}
	if s1 > s2 {
		return 1
	}
	return 0
}

// Contains is a function that checks if a string contains a substring.
func Contains(s, substr string) bool {
	if len(substr) == 0 {
		return true // Substring kosong selalu ada dalam string
	}
	if len(substr) > len(s) {
		return false // Substring lebih panjang dari string tidak mungkin ada
	}

	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i] == substr[0] {
			match := true
			for j := 1; j < len(substr); j++ {
				if s[i+j] != substr[j] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}

// ContainsAny is a function that checks if a string contains any of the substrings.
func ContainsAny(s string, substrs ...string) bool {
	for _, substr := range substrs {
		if Contains(s, substr) {
			return true
		}
	}
	return false
}

// ContainsFunc is a function that checks if a string contains a substring based on a function.
func ContainsFunc(s string, f func(rune) bool) bool {
	for _, r := range s {
		if f(r) {
			return true
		}
	}
	return false
}

// ContainsRune is a function that checks if a string contains a rune.
func ContainsRune(s string, r rune) bool {
	return ContainsFunc(s, func(r1 rune) bool {
		return r1 == r
	})
}

// ContainsAll is a function that checks if a string contains all of the substrings.
func ContainsAll(s string, substrs ...string) bool {
	for _, substr := range substrs {
		if !Contains(s, substr) {
			return false
		}
	}
	return true
}

// Equal is a function that checks if two strings are equal.
func Equal(s1, s2 string) bool {
	return s1 == s2
}

// EqualFold is a function that checks if two strings are equal ignoring case.
func EqualFold(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false // Panjang string berbeda, langsung return false
	}

	for i := 0; i < len(s1); i++ {
		// Bandingkan dengan bitwise operasi untuk ASCII (A-Z atau a-z)
		if (s1[i] | 32) != (s2[i] | 32) {
			return false
		}
	}
	return true
}

// StartsWith is a function that checks if a string starts with a substring.
func StartsWith(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
}

// EndsWith is a function that checks if a string ends with a substring.
func EndsWith(s, substr string) bool {
	return len(s) >= len(substr) && s[len(s)-len(substr):] == substr
}

// IsAnagram is a function that checks if a string is an anagram.
func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[rune]int)
	for _, r := range s1 {
		m[r]++
	}
	for _, r := range s2 {
		m[r]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

// IsPermutation is a function that checks if a string is a permutation.
func IsPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[rune]int)
	for _, r := range s1 {
		m[r]++
	}
	for _, r := range s2 {
		m[r]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

// IsSubstring is a function that checks if a string is a substring.
func IsSubstring(s, substr string) bool {
	return Contains(s, substr)
}

// IsPrefix is a function that checks if a string is a prefix.
func IsPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// IsSuffix is a function that checks if a string is a suffix.
func IsSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// IsSubstringCaseInsensitive is a function that checks if a string is a case-insensitive substring.
func IsSubstringCaseInsensitive(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if ToLowerCase(s[i:i+len(substr)]) == ToLowerCase(substr) {
			return true
		}
	}
	return false
}

// IsPrefixCaseInsensitive is a function that checks if a string is a case-insensitive prefix.
func IsPrefixCaseInsensitive(s, prefix string) bool {
	return len(s) >= len(prefix) && ToLowerCase(s[:len(prefix)]) == ToLowerCase(prefix)
}

// IsSuffixCaseInsensitive is a function that checks if a string is a case-insensitive suffix.
func IsSuffixCaseInsensitive(s, suffix string) bool {
	return len(s) >= len(suffix) && ToLowerCase(s[len(s)-len(suffix):]) == ToLowerCase(suffix)
}

// LongestCommonPrefix is a function that returns the longest common prefix of strings.
func LongestCommonPrefix(strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for _, s := range strs[1:] {
			if i >= len(s) || s[i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// LongestCommonSuffix is a function that returns the longest common suffix of strings.
func LongestCommonSuffix(strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for _, s := range strs[1:] {
			if i >= len(s) || s[len(s)-1-i] != strs[0][len(strs[0])-1-i] {
				return strs[0][len(strs[0])-i:]
			}
		}
	}
	return strs[0]
}

// LongestCommonSubstring is a function that returns the longest common substring of strings.
func LongestCommonSubstring(strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := i + 1; j <= len(strs[0]); j++ {
			substr := strs[0][i:j]
			for _, s := range strs[1:] {
				if !IsSubstring(s, substr) {
					return strs[0][i : j-1]
				}
			}
		}
	}
	return strs[0]
}

// IsSubsequence is a function that checks if a string is a subsequence.
func IsSubsequence(s, substr string) bool {
	i := 0
	for _, r := range substr {
		found := false
		for i < len(s) {
			if rune(s[i]) == r {
				found = true
				i++
				break
			}
			i++
		}
		if !found {
			return false
		}
	}
	return true
}

// ContainsPattern is a function that checks if a string contains a pattern.
func ContainsPattern(s, pattern string) bool {
	for i := 0; i <= len(s)-len(pattern); i++ {
		if IsSubsequence(s[i:], pattern) {
			return true
		}
	}
	return false
}

// AreFrequencyEqual is a function that checks if the frequency of characters in two strings are equal.
func AreFrequencyEqual(s1, s2 string) bool {
	m1 := CharFrequency(s1)
	m2 := CharFrequency(s2)
	if len(m1) != len(m2) {
		return false
	}
	for r, count1 := range m1 {
		count2, ok := m2[r]
		if !ok || count1 != count2 {
			return false
		}
	}
	return true
}
