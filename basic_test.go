package strix_test

import (
	"testing"

	"github.com/go-nop/strix"
)

func TestClone(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "hello"},
		{"world", "world"},
		{"", ""},
	}

	for _, test := range tests {
		result := strix.Clone(test.input)
		if result != test.expected {
			t.Errorf("Clone(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		input    string
		substr   string
		expected int
	}{
		{"hello world", "o", 2},
		{"hello world", "l", 3},
		{"hello world", "x", 0},
	}

	for _, test := range tests {
		result := strix.Count(test.input, test.substr)
		if result != test.expected {
			t.Errorf("Count(%q, %q) = %v; want %v", test.input, test.substr, result, test.expected)
		}
	}
}

func TestFields(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"hello world", []string{"hello", "world"}},
		{"  hello   world  ", []string{"hello", "world"}},
		{"", []string{}},
	}

	for _, test := range tests {
		result := strix.Fields(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("Fields(%q) = %v; want %v", test.input, result, test.expected)
			continue
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("Fields(%q) = %v; want %v", test.input, result, test.expected)
				break
			}
		}
	}
}

func TestFieldsFunc(t *testing.T) {
	tests := []struct {
		input    string
		f        func(byte) bool
		expected []string
	}{
		{"hello,world", func(r byte) bool { return r == ',' }, []string{"hello", "world"}},
		{"a,b,c", func(r byte) bool { return r == ',' }, []string{"a", "b", "c"}},
		{"", func(r byte) bool { return r == ',' }, []string{}},
	}

	for _, test := range tests {
		result := strix.FieldsFunc(test.input, test.f)
		if len(result) != len(test.expected) {
			t.Errorf("FieldsFunc(%q) = %v; want %v", test.input, result, test.expected)
			continue
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("FieldsFunc(%q) = %v; want %v", test.input, result, test.expected)
				break
			}
		}
	}
}

func TestRepeat(t *testing.T) {
	tests := []struct {
		input    string
		n        int
		expected string
	}{
		{"hello", 3, "hellohellohello"},
		{"world", 2, "worldworld"},
		{"", 5, ""},
	}

	for _, test := range tests {
		result := strix.Repeat(test.input, test.n)
		if result != test.expected {
			t.Errorf("Repeat(%q, %d) = %v; want %v", test.input, test.n, result, test.expected)
		}
	}
}
