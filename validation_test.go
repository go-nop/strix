package strix_test

import (
	"testing"

	"github.com/go-nop/strix"
)

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"123", true},
		{"123.45", true},
		{"abc", false},
	}

	for _, test := range tests {
		if result := strix.IsNumeric(test.input); result != test.expected {
			t.Errorf("IsNumeric(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsAlpha(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc", true},
		{"ABC", true},
		{"abc123", false},
	}

	for _, test := range tests {
		if result := strix.IsAlpha(test.input); result != test.expected {
			t.Errorf("IsAlpha(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc123", true},
		{"abc", true},
		{"123", true},
		{"abc-123", false},
	}

	for _, test := range tests {
		if result := strix.IsAlphaNumeric(test.input); result != test.expected {
			t.Errorf("IsAlphaNumeric(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsAlphaDash(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc-123", true},
		{"abc", true},
		{"123", true},
		{"abc_123", false},
	}

	for _, test := range tests {
		if result := strix.IsAlphaDash(test.input); result != test.expected {
			t.Errorf("IsAlphaDash(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsAlphaSpace(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc def", true},
		{"abc", true},
		{"abc123", false},
	}

	for _, test := range tests {
		if result := strix.IsAlphaSpace(test.input); result != test.expected {
			t.Errorf("IsAlphaSpace(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsAlphaNumericDash(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc-123", true},
		{"abc123", true},
		{"abc", true},
		{"abc_123", false},
	}

	for _, test := range tests {
		if result := strix.IsAlphaNumericDash(test.input); result != test.expected {
			t.Errorf("IsAlphaNumericDash(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsAlphaNumericSpace(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc 123", true},
		{"abc123", true},
		{"abc", true},
		{"abc_123", false},
	}

	for _, test := range tests {
		if result := strix.IsAlphaNumericSpace(test.input); result != test.expected {
			t.Errorf("IsAlphaNumericSpace(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsAlphaNumericDashSpace(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc-123 def", true},
		{"abc123", true},
		{"abc", true},
		{"abc_123", false},
	}

	for _, test := range tests {
		if result := strix.IsAlphaNumericDashSpace(test.input); result != test.expected {
			t.Errorf("IsAlphaNumericDashSpace(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsEmail(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"test@example.com", true},
		{"test.example.com", false},
		{"test@.com", false},
	}

	for _, test := range tests {
		if result := strix.IsEmail(test.input); result != test.expected {
			t.Errorf("IsEmail(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsURL(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"http://example.com", true},
		{"https://example.com", true},
		{"ftp://example.com", false},
	}

	for _, test := range tests {
		if result := strix.IsURL(test.input); result != test.expected {
			t.Errorf("IsURL(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsIPv4(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"192.168.0.1", true},
		{"255.255.255.255", true},
		{"256.256.256.256", false},
	}

	for _, test := range tests {
		if result := strix.IsIPv4(test.input); result != test.expected {
			t.Errorf("IsIPv4(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsIPv6(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", true},
		{"2001:db8:85a3::8a2e:370:7334", true},
		{"2001:db8:85a3:0:0:8a2e:370g:7334", false},
	}

	for _, test := range tests {
		if result := strix.IsIPv6(test.input); result != test.expected {
			t.Errorf("IsIPv6(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsHexColor(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"#FFFFFF", true},
		{"#000000", true},
		{"#GGGGGG", false},
	}

	for _, test := range tests {
		if result := strix.IsHexColor(test.input); result != test.expected {
			t.Errorf("IsHexColor(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsRGBColor(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"rgb(255,255,255)", true},
		{"rgb(0,0,0)", true},
		{"rgb(256,256,256)", false},
	}

	for _, test := range tests {
		if result := strix.IsRGBColor(test.input); result != test.expected {
			t.Errorf("IsRGBColor(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsHSLColor(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"hsl(360,100%,50%)", true},
		{"hsl(0,0%,0%)", true},
		{"hsl(361,100%,50%)", false},
	}

	for _, test := range tests {
		if result := strix.IsHSLColor(test.input); result != test.expected {
			t.Errorf("IsHSLColor(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsHexadecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"1A2B3C", true},
		{"123456", true},
		{"1G2H3I", false},
	}

	for _, test := range tests {
		if result := strix.IsHexadecimal(test.input); result != test.expected {
			t.Errorf("IsHexadecimal(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsBase64(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"SGVsbG8gd29ybGQ=", true},
		{"SGVsbG8gd29ybGQ", true},
		{"SGVsbG8gd29ybGQ!", false},
	}

	for _, test := range tests {
		if result := strix.IsBase64(test.input); result != test.expected {
			t.Errorf("IsBase64(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsUUID(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"123e4567-e89b-12d3-a456-426614174000", true},
		{"123e4567-e89b-12d3-a456-42661417400", false},
		{"123e4567-e89b-12d3-a456-4266141740000", false},
	}

	for _, test := range tests {
		if result := strix.IsUUID(test.input); result != test.expected {
			t.Errorf("IsUUID(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"not empty", false},
	}

	for _, test := range tests {
		if result := strix.IsEmpty(test.input); result != test.expected {
			t.Errorf("IsEmpty(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsNotEmpty(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", false},
		{"not empty", true},
	}

	for _, test := range tests {
		if result := strix.IsNotEmpty(test.input); result != test.expected {
			t.Errorf("IsNotEmpty(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsBlank(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"   ", true},
		{"not blank", false},
	}

	for _, test := range tests {
		if result := strix.IsBlank(test.input); result != test.expected {
			t.Errorf("IsBlank(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsNotBlank(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"   ", false},
		{"not blank", true},
	}

	for _, test := range tests {
		if result := strix.IsNotBlank(test.input); result != test.expected {
			t.Errorf("IsNotBlank(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsASCII(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"ASCII", true},
		{"ASCII with space", true},
		{"ASCII with non-ASCII: ñ", false},
	}

	for _, test := range tests {
		if result := strix.IsASCII(test.input); result != test.expected {
			t.Errorf("IsASCII(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsPrintable(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Printable", true},
		{"Printable with space", true},
		{"Non-printable: \x01", false},
	}

	for _, test := range tests {
		if result := strix.IsPrintable(test.input); result != test.expected {
			t.Errorf("IsPrintable(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsWhitespace(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{" \t\n\v\f\r", true},
		{"Whitespace", false},
	}

	for _, test := range tests {
		if result := strix.IsWhitespace(test.input); result != test.expected {
			t.Errorf("IsWhitespace(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsDigit(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"123", true},
		{"123a", false},
	}

	for _, test := range tests {
		if result := strix.IsDigit(test.input); result != test.expected {
			t.Errorf("IsDigit(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsLetter(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc", true},
		{"abc123", false},
	}

	for _, test := range tests {
		if result := strix.IsLetter(test.input); result != test.expected {
			t.Errorf("IsLetter(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsLower(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc", true},
		{"ABC", false},
	}

	for _, test := range tests {
		if result := strix.IsLower(test.input); result != test.expected {
			t.Errorf("IsLower(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsUpper(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"ABC", true},
		{"abc", false},
	}

	for _, test := range tests {
		if result := strix.IsUpper(test.input); result != test.expected {
			t.Errorf("IsUpper(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsTitle(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Title", true},
		{"title", false},
	}

	for _, test := range tests {
		if result := strix.IsTitle(test.input); result != test.expected {
			t.Errorf("IsTitle(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsPascal(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"PascalCase", true},
		{"pascalCase", false},
	}

	for _, test := range tests {
		if result := strix.IsPascal(test.input); result != test.expected {
			t.Errorf("IsPascal(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsCamel(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"camelCase", true},
		{"CamelCase", false},
	}

	for _, test := range tests {
		if result := strix.IsCamel(test.input); result != test.expected {
			t.Errorf("IsCamel(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsSnake(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"snake_case", true},
		{"SnakeCase", false},
	}

	for _, test := range tests {
		if result := strix.IsSnake(test.input); result != test.expected {
			t.Errorf("IsSnake(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsKebab(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"kebab-case", true},
		{"KebabCase", false},
	}

	for _, test := range tests {
		if result := strix.IsKebab(test.input); result != test.expected {
			t.Errorf("IsKebab(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsSpace(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{" ", true},
		{"not space", false},
	}

	for _, test := range tests {
		if result := strix.IsSpace(test.input); result != test.expected {
			t.Errorf("IsSpace(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"madam", true},
		{"hello", false},
	}

	for _, test := range tests {
		if result := strix.IsPalindrome(test.input); result != test.expected {
			t.Errorf("IsPalindrome(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
