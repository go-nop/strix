package strix_test

import (
	"testing"

	"github.com/go-nop/strix"
)

func TestCompare(t *testing.T) {
	tests := []struct {
		s1, s2   string
		expected int
	}{
		{"a", "b", -1},
		{"b", "a", 1},
		{"a", "a", 0},
	}

	for _, test := range tests {
		result := strix.Compare(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("Compare(%q, %q) = %v; want %v", test.s1, test.s2, result, test.expected)
		}
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		s, substr string
		expected  bool
	}{
		{"hello world", "world", true},
		{"hello world", "worlds", false},
		{"hello world", "", true},
	}

	for _, test := range tests {
		result := strix.Contains(test.s, test.substr)
		if result != test.expected {
			t.Errorf("Contains(%q, %q) = %v; want %v", test.s, test.substr, result, test.expected)
		}
	}
}

func TestContainsAny(t *testing.T) {
	tests := []struct {
		s        string
		substrs  []string
		expected bool
	}{
		{"hello world", []string{"world", "earth"}, true},
		{"hello world", []string{"earth", "mars"}, false},
		{"hello world", []string{}, false},
	}

	for _, test := range tests {
		result := strix.ContainsAny(test.s, test.substrs...)
		if result != test.expected {
			t.Errorf("ContainsAny(%q, %v) = %v; want %v", test.s, test.substrs, result, test.expected)
		}
	}
}

func TestContainsFunc(t *testing.T) {
	tests := []struct {
		s        string
		f        func(rune) bool
		expected bool
	}{
		{"hello world", func(r rune) bool { return r == 'w' }, true},
		{"hello world", func(r rune) bool { return r == 'x' }, false},
	}

	for _, test := range tests {
		result := strix.ContainsFunc(test.s, test.f)
		if result != test.expected {
			t.Errorf("ContainsFunc(%q) = %v; want %v", test.s, result, test.expected)
		}
	}
}

func TestContainsRune(t *testing.T) {
	tests := []struct {
		s        string
		r        rune
		expected bool
	}{
		{"hello world", 'w', true},
		{"hello world", 'x', false},
	}

	for _, test := range tests {
		result := strix.ContainsRune(test.s, test.r)
		if result != test.expected {
			t.Errorf("ContainsRune(%q, %q) = %v; want %v", test.s, test.r, result, test.expected)
		}
	}
}

func TestContainsAll(t *testing.T) {
	tests := []struct {
		s        string
		substrs  []string
		expected bool
	}{
		{"hello world", []string{"hello", "world"}, true},
		{"hello world", []string{"hello", "earth"}, false},
		{"hello world", []string{}, true},
	}

	for _, test := range tests {
		result := strix.ContainsAll(test.s, test.substrs...)
		if result != test.expected {
			t.Errorf("ContainsAll(%q, %v) = %v; want %v", test.s, test.substrs, result, test.expected)
		}
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		s1, s2   string
		expected bool
	}{
		{"hello", "hello", true},
		{"hello", "world", false},
	}

	for _, test := range tests {
		result := strix.Equal(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("Equal(%q, %q) = %v; want %v", test.s1, test.s2, result, test.expected)
		}
	}
}

func TestEqualFold(t *testing.T) {
	tests := []struct {
		s1, s2   string
		expected bool
	}{
		{"hello", "HELLO", true},
		{"hello", "world", false},
	}

	for _, test := range tests {
		result := strix.EqualFold(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("EqualFold(%q, %q) = %v; want %v", test.s1, test.s2, result, test.expected)
		}
	}
}

func TestStartsWith(t *testing.T) {
	tests := []struct {
		s, substr string
		expected  bool
	}{
		{"hello world", "hello", true},
		{"hello world", "world", false},
	}

	for _, test := range tests {
		result := strix.StartsWith(test.s, test.substr)
		if result != test.expected {
			t.Errorf("StartsWith(%q, %q) = %v; want %v", test.s, test.substr, result, test.expected)
		}
	}
}

func TestEndsWith(t *testing.T) {
	tests := []struct {
		s, substr string
		expected  bool
	}{
		{"hello world", "world", true},
		{"hello world", "hello", false},
	}

	for _, test := range tests {
		result := strix.EndsWith(test.s, test.substr)
		if result != test.expected {
			t.Errorf("EndsWith(%q, %q) = %v; want %v", test.s, test.substr, result, test.expected)
		}
	}
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s1, s2   string
		expected bool
	}{
		{"listen", "silent", true},
		{"hello", "world", false},
	}

	for _, test := range tests {
		result := strix.IsAnagram(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("IsAnagram(%q, %q) = %v; want %v", test.s1, test.s2, result, test.expected)
		}
	}
}

func TestIsPermutation(t *testing.T) {
	tests := []struct {
		s1, s2   string
		expected bool
	}{
		{"abc", "bca", true},
		{"abc", "def", false},
	}

	for _, test := range tests {
		result := strix.IsPermutation(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("IsPermutation(%q, %q) = %v; want %v", test.s1, test.s2, result, test.expected)
		}
	}
}

func TestIsSubstring(t *testing.T) {
	tests := []struct {
		s, substr string
		expected  bool
	}{
		{"hello world", "world", true},
		{"hello world", "earth", false},
	}

	for _, test := range tests {
		result := strix.IsSubstring(test.s, test.substr)
		if result != test.expected {
			t.Errorf("IsSubstring(%q, %q) = %v; want %v", test.s, test.substr, result, test.expected)
		}
	}
}

func TestIsPrefix(t *testing.T) {
	tests := []struct {
		s, prefix string
		expected  bool
	}{
		{"hello world", "hello", true},
		{"hello world", "world", false},
	}

	for _, test := range tests {
		result := strix.IsPrefix(test.s, test.prefix)
		if result != test.expected {
			t.Errorf("IsPrefix(%q, %q) = %v; want %v", test.s, test.prefix, result, test.expected)
		}
	}
}

func TestIsSuffix(t *testing.T) {
	tests := []struct {
		s, suffix string
		expected  bool
	}{
		{"hello world", "world", true},
		{"hello world", "hello", false},
	}

	for _, test := range tests {
		result := strix.IsSuffix(test.s, test.suffix)
		if result != test.expected {
			t.Errorf("IsSuffix(%q, %q) = %v; want %v", test.s, test.suffix, result, test.expected)
		}
	}
}

func TestIsSubstringCaseInsensitive(t *testing.T) {
	tests := []struct {
		s, substr string
		expected  bool
	}{
		{"hello world", "WORLD", true},
		{"hello world", "EARTH", false},
	}

	for _, test := range tests {
		result := strix.IsSubstringCaseInsensitive(test.s, test.substr)
		if result != test.expected {
			t.Errorf("IsSubstringCaseInsensitive(%q, %q) = %v; want %v", test.s, test.substr, result, test.expected)
		}
	}
}

func TestIsPrefixCaseInsensitive(t *testing.T) {
	tests := []struct {
		s, prefix string
		expected  bool
	}{
		{"hello world", "HELLO", true},
		{"hello world", "WORLD", false},
	}

	for _, test := range tests {
		result := strix.IsPrefixCaseInsensitive(test.s, test.prefix)
		if result != test.expected {
			t.Errorf("IsPrefixCaseInsensitive(%q, %q) = %v; want %v", test.s, test.prefix, result, test.expected)
		}
	}
}

func TestIsSuffixCaseInsensitive(t *testing.T) {
	tests := []struct {
		s, suffix string
		expected  bool
	}{
		{"hello world", "WORLD", true},
		{"hello world", "HELLO", false},
	}

	for _, test := range tests {
		result := strix.IsSuffixCaseInsensitive(test.s, test.suffix)
		if result != test.expected {
			t.Errorf("IsSuffixCaseInsensitive(%q, %q) = %v; want %v", test.s, test.suffix, result, test.expected)
		}
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs     []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}

	for _, test := range tests {
		result := strix.LongestCommonPrefix(test.strs...)
		if result != test.expected {
			t.Errorf("LongestCommonPrefix(%v) = %v; want %v", test.strs, result, test.expected)
		}
	}
}

func TestLongestCommonSuffix(t *testing.T) {
	tests := []struct {
		strs     []string
		expected string
	}{
		{[]string{"flower", "tower", "shower"}, "ower"},
		{[]string{"dog", "racecar", "car"}, ""},
	}

	for _, test := range tests {
		result := strix.LongestCommonSuffix(test.strs...)
		if result != test.expected {
			t.Errorf("LongestCommonSuffix(%v) = %v; want %v", test.strs, result, test.expected)
		}
	}
}

func TestLongestCommonSubstring(t *testing.T) {
	tests := []struct {
		strs     []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}

	for _, test := range tests {
		result := strix.LongestCommonSubstring(test.strs...)
		if result != test.expected {
			t.Errorf("LongestCommonSubstring(%v) = %v; want %v", test.strs, result, test.expected)
		}
	}
}

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s, substr string
		expected  bool
	}{
		{"hello world", "hlo", true},
		{"hello world", "hld", true},
	}

	for _, test := range tests {
		result := strix.IsSubsequence(test.s, test.substr)
		if result != test.expected {
			t.Errorf("IsSubsequence(%q, %q) = %v; want %v", test.s, test.substr, result, test.expected)
		}
	}
}

func TestContainsPattern(t *testing.T) {
	tests := []struct {
		s, pattern string
		expected   bool
	}{
		{"hello world", "hlo", true},
		{"hello world", "hld", true},
	}

	for _, test := range tests {
		result := strix.ContainsPattern(test.s, test.pattern)
		if result != test.expected {
			t.Errorf("ContainsPattern(%q, %q) = %v; want %v", test.s, test.pattern, result, test.expected)
		}
	}
}

func TestAreFrequencyEqual(t *testing.T) {
	tests := []struct {
		s1, s2   string
		expected bool
	}{
		{"hello", "lleho", true},
		{"hello", "world", false},
	}

	for _, test := range tests {
		result := strix.AreFrequencyEqual(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("AreFrequencyEqual(%q, %q) = %v; want %v", test.s1, test.s2, result, test.expected)
		}
	}
}
