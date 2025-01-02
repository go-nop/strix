package strix_test

import (
	"testing"

	"github.com/go-nop/strix"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		input    string
		sep      string
		expected []string
	}{
		{"the quick brown fox", " ", []string{"the", "quick", "brown", "fox"}},
		{"hello,world", ",", []string{"hello", "world"}},
		{"no separator", "-", []string{"no separator"}},
	}

	for _, test := range tests {
		result := strix.Split(test.input, test.sep)
		if len(result) != len(test.expected) {
			t.Errorf("Split(%q, %q) = %v; want %v", test.input, test.sep, result, test.expected)
			continue
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("Split(%q, %q) = %v; want %v", test.input, test.sep, result, test.expected)
				break
			}
		}
	}
}

func TestTrimSpace(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"  hello  ", "hello"},
		{"no spaces", "no spaces"},
		{"  leading space", "leading space"},
		{"trailing space  ", "trailing space"},
	}

	for _, test := range tests {
		result := strix.TrimSpace(test.input)
		if result != test.expected {
			t.Errorf("TrimSpace(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestReplaceAll(t *testing.T) {
	tests := []struct {
		input    string
		old      string
		new      string
		expected string
	}{
		{"the quick brown fox", "the", "a", "a quick brown fox"},
		{"hello world", "world", "gopher", "hello gopher"},
		{"no match", "x", "y", "no match"},
	}

	for _, test := range tests {
		result := strix.ReplaceAll(test.input, test.old, test.new)
		if result != test.expected {
			t.Errorf("ReplaceAll(%q, %q, %q) = %v; want %v", test.input, test.old, test.new, result, test.expected)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"", ""},
	}

	for _, test := range tests {
		result := strix.Reverse(test.input)
		if result != test.expected {
			t.Errorf("Reverse(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestToTitleCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "Hello World"},
		{"the quick brown fox", "The Quick Brown Fox"},
		{"", ""},
	}

	for _, test := range tests {
		result := strix.ToTitleCase(test.input)
		if result != test.expected {
			t.Errorf("ToTitleCase(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestToUpperCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "HELLO"},
		{"world", "WORLD"},
		{"", ""},
	}

	for _, test := range tests {
		result := strix.ToUpperCase(test.input)
		if result != test.expected {
			t.Errorf("ToUpperCase(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestToLowerCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HELLO", "hello"},
		{"WORLD", "world"},
		{"", ""},
	}

	for _, test := range tests {
		result := strix.ToLowerCase(test.input)
		if result != test.expected {
			t.Errorf("ToLowerCase(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HelloWorld", "hello_world"},
		{"snakeCase", "snake_case"},
		{"", ""},
	}

	for _, test := range tests {
		result := strix.SnakeCase(test.input)
		if result != test.expected {
			t.Errorf("SnakeCase(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "helloWorld"},
		{"camel case", "camelCase"},
		{"", ""},
	}

	for _, test := range tests {
		result := strix.CamelCase(test.input)
		if result != test.expected {
			t.Errorf("CamelCase(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestPascalCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "HelloWorld"},
		{"pascal case", "PascalCase"},
		{"", ""},
	}

	for _, test := range tests {
		result := strix.PascalCase(test.input)
		if result != test.expected {
			t.Errorf("PascalCase(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestKebabCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HelloWorld", "hello-world"},
		{"kebabCase", "kebab-case"},
		{"", ""},
	}

	for _, test := range tests {
		result := strix.KebabCase(test.input)
		if result != test.expected {
			t.Errorf("KebabCase(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestToogleCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HelloWorld", "hELLOwORLD"},
		{"toggleCase", "TOGGLEcASE"},
		{"", ""},
	}

	for _, test := range tests {
		result := strix.ToogleCase(test.input)
		if result != test.expected {
			t.Errorf("ToogleCase(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestCapitalizeFirst(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "Hello"},
		{"world", "World"},
		{"", ""},
	}

	for _, test := range tests {
		result := strix.CapitalizeFirst(test.input)
		if result != test.expected {
			t.Errorf("CapitalizeFirst(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		hasError bool
	}{
		{"123", 123, false},
		{"-123", -123, false},
		{"abc", 0, true},
	}

	for _, test := range tests {
		result, err := strix.ToInt(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("ToInt(%q) error = %v; want error = %v", test.input, err, test.hasError)
		}
		if result != test.expected {
			t.Errorf("ToInt(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
