package strix_test

import (
	"testing"

	"github.com/go-nop/strix"
)

func TestCharFrequency(t *testing.T) {
	tests := []struct {
		input    string
		expected map[rune]int
	}{
		{"hello", map[rune]int{'h': 1, 'e': 1, 'l': 2, 'o': 1}},
		{"world", map[rune]int{'w': 1, 'o': 1, 'r': 1, 'l': 1, 'd': 1}},
		{"", map[rune]int{}},
	}

	for _, test := range tests {
		result := strix.CharFrequency(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("CharFrequency(%q) = %v; want %v", test.input, result, test.expected)
			continue
		}
		for k, v := range result {
			if v != test.expected[k] {
				t.Errorf("CharFrequency(%q) = %v; want %v", test.input, result, test.expected)
				break
			}
		}
	}
}

func TestCharFrequencyCaseInsensitive(t *testing.T) {
	tests := []struct {
		input    string
		expected map[rune]int
	}{
		{"Hello", map[rune]int{'h': 1, 'e': 1, 'l': 2, 'o': 1}},
		{"WORLD", map[rune]int{'w': 1, 'o': 1, 'r': 1, 'l': 1, 'd': 1}},
		{"", map[rune]int{}},
	}

	for _, test := range tests {
		result := strix.CharFrequencyCaseInsensitive(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("CharFrequencyCaseInsensitive(%q) = %v; want %v", test.input, result, test.expected)
			continue
		}
		for k, v := range result {
			if v != test.expected[k] {
				t.Errorf("CharFrequencyCaseInsensitive(%q) = %v; want %v", test.input, result, test.expected)
				break
			}
		}
	}
}

func TestCharFrequencyAlpha(t *testing.T) {
	tests := []struct {
		input    string
		expected map[rune]int
	}{
		{"hello123", map[rune]int{'h': 1, 'e': 1, 'l': 2, 'o': 1}},
		{"world!", map[rune]int{'w': 1, 'o': 1, 'r': 1, 'l': 1, 'd': 1}},
		{"123", map[rune]int{}},
	}

	for _, test := range tests {
		result := strix.CharFrequencyAlpha(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("CharFrequencyAlpha(%q) = %v; want %v", test.input, result, test.expected)
			continue
		}
		for k, v := range result {
			if v != test.expected[k] {
				t.Errorf("CharFrequencyAlpha(%q) = %v; want %v", test.input, result, test.expected)
				break
			}
		}
	}
}

func TestCharFrequencyNumeric(t *testing.T) {
	tests := []struct {
		input    string
		expected map[rune]int
	}{
		{"hello123", map[rune]int{'1': 1, '2': 1, '3': 1}},
		{"world456", map[rune]int{'4': 1, '5': 1, '6': 1}},
		{"abc", map[rune]int{}},
	}

	for _, test := range tests {
		result := strix.CharFrequencyNumeric(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("CharFrequencyNumeric(%q) = %v; want %v", test.input, result, test.expected)
			continue
		}
		for k, v := range result {
			if v != test.expected[k] {
				t.Errorf("CharFrequencyNumeric(%q) = %v; want %v", test.input, result, test.expected)
				break
			}
		}
	}
}

func TestHasPrefix(t *testing.T) {
	tests := []struct {
		input    string
		prefix   string
		expected bool
	}{
		{"hello", "he", true},
		{"world", "wo", true},
		{"hello", "wo", false},
	}

	for _, test := range tests {
		result := strix.HasPrefix(test.input, test.prefix)
		if result != test.expected {
			t.Errorf("HasPrefix(%q, %q) = %v; want %v", test.input, test.prefix, result, test.expected)
		}
	}
}

func TestHasSuffix(t *testing.T) {
	tests := []struct {
		input    string
		suffix   string
		expected bool
	}{
		{"hello", "lo", true},
		{"world", "ld", true},
		{"hello", "ld", false},
	}

	for _, test := range tests {
		result := strix.HasSuffix(test.input, test.suffix)
		if result != test.expected {
			t.Errorf("HasSuffix(%q, %q) = %v; want %v", test.input, test.suffix, result, test.expected)
		}
	}
}

func TestIndex(t *testing.T) {
	tests := []struct {
		input    string
		substr   string
		expected int
	}{
		{"hello", "ll", 2},
		{"world", "or", 1},
		{"hello", "wo", -1},
	}

	for _, test := range tests {
		result := strix.Index(test.input, test.substr)
		if result != test.expected {
			t.Errorf("Index(%q, %q) = %v; want %v", test.input, test.substr, result, test.expected)
		}
	}
}

func TestIndexAny(t *testing.T) {
	tests := []struct {
		input    string
		chars    string
		expected int
	}{
		{"hello", "aeiou", 1},
		{"world", "aeiou", 1},
		{"hello", "xyz", -1},
	}

	for _, test := range tests {
		result := strix.IndexAny(test.input, test.chars)
		if result != test.expected {
			t.Errorf("IndexAny(%q, %q) = %v; want %v", test.input, test.chars, result, test.expected)
		}
	}
}

func TestIndexByte(t *testing.T) {
	tests := []struct {
		input    string
		c        byte
		expected int
	}{
		{"hello", 'e', 1},
		{"world", 'o', 1},
		{"hello", 'x', -1},
	}

	for _, test := range tests {
		result := strix.IndexByte(test.input, test.c)
		if result != test.expected {
			t.Errorf("IndexByte(%q, %q) = %v; want %v", test.input, test.c, result, test.expected)
		}
	}
}

func TestIndexFunc(t *testing.T) {
	tests := []struct {
		input    string
		f        func(rune) bool
		expected int
	}{
		{"hello", func(r rune) bool { return r == 'e' }, 1},
		{"world", func(r rune) bool { return r == 'o' }, 1},
		{"hello", func(r rune) bool { return r == 'x' }, -1},
	}

	for _, test := range tests {
		result := strix.IndexFunc(test.input, test.f)
		if result != test.expected {
			t.Errorf("IndexFunc(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIndexRune(t *testing.T) {
	tests := []struct {
		input    string
		r        rune
		expected int
	}{
		{"hello", 'e', 1},
		{"world", 'o', 1},
		{"hello", 'x', -1},
	}

	for _, test := range tests {
		result := strix.IndexRune(test.input, test.r)
		if result != test.expected {
			t.Errorf("IndexRune(%q, %q) = %v; want %v", test.input, test.r, result, test.expected)
		}
	}
}

func TestLastIndex(t *testing.T) {
	tests := []struct {
		input    string
		substr   string
		expected int
	}{
		{"hello", "l", 3},
		{"world", "o", 1},
		{"hello", "x", -1},
	}

	for _, test := range tests {
		result := strix.LastIndex(test.input, test.substr)
		if result != test.expected {
			t.Errorf("LastIndex(%q, %q) = %v; want %v", test.input, test.substr, result, test.expected)
		}
	}
}

func TestLastIndexAny(t *testing.T) {
	tests := []struct {
		input    string
		chars    string
		expected int
	}{
		{"hello", "aeiou", 4},
		{"world", "aeiou", 1},
		{"hello", "xyz", -1},
	}

	for _, test := range tests {
		result := strix.LastIndexAny(test.input, test.chars)
		if result != test.expected {
			t.Errorf("LastIndexAny(%q, %q) = %v; want %v", test.input, test.chars, result, test.expected)
		}
	}
}

func TestLastIndexByte(t *testing.T) {
	tests := []struct {
		input    string
		c        byte
		expected int
	}{
		{"hello", 'l', 3},
		{"world", 'o', 1},
		{"hello", 'x', -1},
	}

	for _, test := range tests {
		result := strix.LastIndexByte(test.input, test.c)
		if result != test.expected {
			t.Errorf("LastIndexByte(%q, %q) = %v; want %v", test.input, test.c, result, test.expected)
		}
	}
}

func TestLastIndexFunc(t *testing.T) {
	tests := []struct {
		input    string
		f        func(rune) bool
		expected int
	}{
		{"hello", func(r rune) bool { return r == 'l' }, 3},
		{"world", func(r rune) bool { return r == 'o' }, 1},
		{"hello", func(r rune) bool { return r == 'x' }, -1},
	}

	for _, test := range tests {
		result := strix.LastIndexFunc(test.input, test.f)
		if result != test.expected {
			t.Errorf("LastIndexFunc(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestLastIndexRune(t *testing.T) {
	tests := []struct {
		input    string
		r        rune
		expected int
	}{
		{"hello", 'l', 3},
		{"world", 'o', 1},
		{"hello", 'x', -1},
	}

	for _, test := range tests {
		result := strix.LastIndexRune(test.input, test.r)
		if result != test.expected {
			t.Errorf("LastIndexRune(%q, %q) = %v; want %v", test.input, test.r, result, test.expected)
		}
	}
}
