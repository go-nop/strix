package strix

// IsNumeric is a function that checks if a string is numeric.
func IsNumeric(s string) bool {
	for _, r := range s {
		if (r < '0' || r > '9') && r != '.' {
			return false
		}
	}
	return true
}

// IsAlpha is a function that checks if a string is alphabetic.
func IsAlpha(s string) bool {
	for _, r := range s {
		if r < 'A' || (r > 'Z' && r < 'a') || r > 'z' {
			return false
		}
	}
	return true
}

// IsAlphaNumeric is a function that checks if a string is alphanumeric.
func IsAlphaNumeric(s string) bool {
	for _, r := range s {
		if (r < 'A' || (r > 'Z' && r < 'a') || r > 'z') && (r < '0' || r > '9') {
			return false
		}
	}
	return true
}

// IsAlphaDash is a function that checks if a string is alphabetic with dashes.
func IsAlphaDash(s string) bool {
	for _, r := range s {
		if (r < 'A' || (r > 'Z' && r < 'a') || r > 'z') && (r < '0' || r > '9') && r != '-' {
			return false
		}
	}
	return true
}

// IsAlphaSpace is a function that checks if a string is alphabetic with spaces.
func IsAlphaSpace(s string) bool {
	for _, r := range s {
		if (r < 'A' || (r > 'Z' && r < 'a') || r > 'z') && r != ' ' {
			return false
		}
	}
	return true
}

// IsAlphaNumericDash is a function that checks if a string is alphanumeric with dashes.
func IsAlphaNumericDash(s string) bool {
	for _, r := range s {
		if (r < 'A' || (r > 'Z' && r < 'a') || r > 'z') && (r < '0' || r > '9') && r != '-' {
			return false
		}
	}
	return true
}

// IsAlphaNumericSpace is a function that checks if a string is alphanumeric with spaces.
func IsAlphaNumericSpace(s string) bool {
	for _, r := range s {
		if (r < 'A' || (r > 'Z' && r < 'a') || r > 'z') && (r < '0' || r > '9') && r != ' ' {
			return false
		}
	}
	return true
}

// IsAlphaNumericDashSpace is a function that checks if a string is alphanumeric with dashes and spaces.
func IsAlphaNumericDashSpace(s string) bool {
	for _, r := range s {
		if (r < 'A' || (r > 'Z' && r < 'a') || r > 'z') && (r < '0' || r > '9') && r != '-' && r != ' ' {
			return false
		}
	}
	return true
}

// IsEmail is a function that checks if a string is an email.
func IsEmail(s string) bool {
	at := false
	for i, r := range s {
		if r == '@' {
			if at {
				return false
			}
			at = true
			if i == 0 {
				return false
			}
		} else if r == '.' {
			if i == 0 || i == len(s)-1 || s[i-1] == '.' || s[i-1] == '@' {
				return false
			}
		} else if (r < 'A' || (r > 'Z' && r < 'a') || r > 'z') && (r < '0' || r > '9') && r != '-' && r != '_' {
			return false
		}
	}
	return at
}

// IsURL is a function that checks if a string is a URL.
func IsURL(s string) bool {
	if len(s) < 8 {
		return false
	}
	if s[:7] != "http://" && s[:8] != "https://" {
		return false
	}
	for _, r := range s[7:] {
		if (r < 'A' || (r > 'Z' && r < 'a') || r > 'z') && (r < '0' || r > '9') && r != '.' && r != '-' && r != '_' && r != '/' {
			return false
		}
	}
	return true
}

// IsIPv4 is a function that checks if a string is an IPv4 address.
func IsIPv4(s string) bool {
	parts := Split(s, ".")
	if len(parts) != 4 {
		return false
	}
	for _, part := range parts {
		if len(part) == 0 || len(part) > 3 {
			return false
		}
		for _, r := range part {
			if r < '0' || r > '9' {
				return false
			}
		}
		num, err := toInt(part)
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}
	return true
}

// IsIPv6 is a function that checks if a string is an IPv6 address.
func IsIPv6(s string) bool {
	if len(s) < 2 {
		return false
	}
	for i, r := range s {
		if r == ':' {
			if i == 0 || i == len(s)-1 {
				return false
			}
		} else if (r < 'A' || (r > 'F' && r < 'a') || r > 'f') && (r < '0' || r > '9') && r != ':' {
			return false
		}
	}
	return true
}

// IsHexColor is a function that checks if a string is a hex color.
func IsHexColor(s string) bool {
	if len(s) != 7 {
		return false
	}
	if s[0] != '#' {
		return false
	}
	for i, r := range s[1:] {
		if (r < 'A' || (r > 'F' && r < 'a') || r > 'f') && (r < '0' || r > '9') {
			return false
		}
		if i == 6 {
			return false
		}
	}
	return true
}

// IsRGBColor is a function that checks if a string is an RGB color.
func IsRGBColor(s string) bool {
	if len(s) < 10 {
		return false
	}
	if s[:4] != "rgb(" || s[len(s)-1] != ')' {
		return false
	}
	values := s[4 : len(s)-1]
	parts := Split(values, ",")
	if len(parts) != 3 {
		return false
	}
	for _, part := range parts {
		part = TrimSpace(part)
		num, err := ToInt(part)
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}
	return true
}

// IsHSLColor is a function that checks if a string is an HSL color.
func IsHSLColor(s string) bool {
	if len(s) < 10 {
		return false
	}
	if s[:4] != "hsl(" || s[len(s)-1] != ')' {
		return false
	}
	values := s[4 : len(s)-1]
	parts := Split(values, ",")
	if len(parts) != 3 {
		return false
	}
	hue, err := ToInt(TrimSpace(parts[0]))
	if err != nil || hue < 0 || hue > 360 {
		return false
	}
	saturation := TrimSpace(parts[1])
	if len(saturation) < 2 || saturation[len(saturation)-1] != '%' {
		return false
	}
	satValue, err := ToInt(saturation[:len(saturation)-1])
	if err != nil || satValue < 0 || satValue > 100 {
		return false
	}
	lightness := TrimSpace(parts[2])
	if len(lightness) < 2 || lightness[len(lightness)-1] != '%' {
		return false
	}
	lightValue, err := ToInt(lightness[:len(lightness)-1])
	if err != nil || lightValue < 0 || lightValue > 100 {
		return false
	}
	return true
}

// IsHexadecimal is a function that checks if a string is hexadecimal.
func IsHexadecimal(s string) bool {
	for _, r := range s {
		if (r < 'A' || (r > 'F' && r < 'a') || r > 'f') && (r < '0' || r > '9') {
			return false
		}
	}
	return true
}

// IsBase64 is a function that checks if a string is base64.
func IsBase64(s string) bool {
	for _, r := range s {
		if (r < 'A' || (r > 'Z' && r < 'a') || r > 'z') && (r < '0' || r > '9') && r != '+' && r != '/' && r != '=' {
			return false
		}
	}
	return true
}

// IsUUID is a function that checks if a string is a UUID.
func IsUUID(s string) bool {
	if len(s) != 36 {
		return false
	}
	for i, r := range s {
		if r == '-' {
			if i != 8 && i != 13 && i != 18 && i != 23 {
				return false
			}
		} else if (r < 'A' || (r > 'F' && r < 'a') || r > 'f') && (r < '0' || r > '9') {
			return false
		}
	}
	return true
}

// IsEmpty is a function that checks if a string is empty.
func IsEmpty(s string) bool {
	return len(s) == 0
}

// IsNotEmpty is a function that checks if a string is not empty.
func IsNotEmpty(s string) bool {
	return len(s) > 0
}

// IsBlank is a function that checks if a string is blank.
func IsBlank(s string) bool {
	for _, r := range s {
		if r != ' ' {
			return false
		}
	}
	return true
}

// IsNotBlank is a function that checks if a string is not blank.
func IsNotBlank(s string) bool {
	for _, r := range s {
		if r != ' ' {
			return true
		}
	}
	return false
}

// IsASCII is a function that checks if a string is ASCII.
func IsASCII(s string) bool {
	for _, r := range s {
		if r > 127 {
			return false
		}
	}
	return true
}

// IsPrintable is a function that checks if a string is printable.
func IsPrintable(s string) bool {
	for _, r := range s {
		if r < 32 || r > 126 {
			return false
		}
	}
	return true
}

// IsWhitespace is a function that checks if a string is whitespace.
func IsWhitespace(s string) bool {
	for _, r := range s {
		if r != ' ' && r != '\t' && r != '\n' && r != '\v' && r != '\f' && r != '\r' {
			return false
		}
	}
	return true
}

// IsDigit is a function that checks if a string is a digit.
func IsDigit(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// IsLetter is a function that checks if a string is a letter.
func IsLetter(s string) bool {
	for _, r := range s {
		if r < 'A' || (r > 'Z' && r < 'a') || r > 'z' {
			return false
		}
	}
	return true
}

// IsLower is a function that checks if a string is lowercase.
func IsLower(s string) bool {
	for _, r := range s {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}

// IsUpper is a function that checks if a string is uppercase.
func IsUpper(s string) bool {
	for _, r := range s {
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	return true
}

// IsTitle is a function that checks if a string is title case.
func IsTitle(s string) bool {
	if len(s) == 0 {
		return false
	}
	words := Split(s, " ")
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		if word[0] < 'A' || word[0] > 'Z' {
			return false
		}
		for _, r := range word[1:] {
			if r < 'a' || r > 'z' {
				return false
			}
		}
	}
	return true
}

// IsPascal is a function that checks if a string is pascal case.
func IsPascal(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] < 'A' || s[0] > 'Z' {
		return false
	}
	for _, r := range s[1:] {
		if r < 'A' || (r > 'Z' && r < 'a') || r > 'z' {
			return false
		}
	}
	return true
}

// IsCamel is a function that checks if a string is camel case.
func IsCamel(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] < 'a' || s[0] > 'z' {
		return false
	}
	for _, r := range s {
		if r < 'A' || (r > 'Z' && r < 'a') || r > 'z' {
			return false
		}
	}
	return true
}

// IsSnake is a function that checks if a string is snake case.
func IsSnake(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] < 'a' || s[0] > 'z' {
		return false
	}
	for _, r := range s {
		if (r < 'A' || (r > 'Z' && r < 'a') || r > 'z') && r != '_' {
			return false
		}
	}
	return true
}

// IsKebab is a function that checks if a string is kebab case.
func IsKebab(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] < 'a' || s[0] > 'z' {
		return false
	}
	for _, r := range s {
		if (r < 'A' || (r > 'Z' && r < 'a') || r > 'z') && r != '-' {
			return false
		}
	}
	return true
}

// IsSpace is a function that checks if a string is space case.
func IsSpace(s string) bool {
	for _, r := range s {
		if r != ' ' {
			return false
		}
	}
	return true
}

// IsPalindrome is a function that checks if a string is a palindrome.
func IsPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
