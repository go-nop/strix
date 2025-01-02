package strix

// Cut is a function that removes a substring from a string.
func Cut(s, substr string) string {
	return ReplaceAll(s, substr, "")
}

// CutPrefix is a function that removes a prefix from a string.
func CutPrefix(s, prefix string) string {
	if len(s) < len(prefix) {
		return s
	}
	if s[:len(prefix)] == prefix {
		return s[len(prefix):]
	}
	return s
}

// CutSuffix is a function that removes a suffix from a string.
func CutSuffix(s, suffix string) string {
	if len(s) < len(suffix) {
		return s
	}
	if s[len(s)-len(suffix):] == suffix {
		return s[:len(s)-len(suffix)]
	}
	return s
}

// Split is a function that splits a string into a slice of substrings.
func Split(s, sep string) []string {
	if sep == "" {
		return []string{s}
	}
	if len(sep) == 1 {
		return splitByte(s, sep[0])
	}
	return splitOptimized(s, sep)
}

// TrimSpace is a function that removes leading and trailing whitespaces from a string.
func TrimSpace(s string) string {
	var start, end int
	for start < len(s) && s[start] == ' ' {
		start++
	}
	for end = len(s) - 1; end >= 0 && s[end] == ' '; end-- {
	}
	return s[start : end+1]
}

// Replace is a function that replaces the first occurrence of a substring with another substring.
func Replace(s, old, new string) string {
	if old == "" {
		return s
	}

	i := Index(s, old)
	if i == -1 {
		return s
	}

	sb := GetBuilder()
	defer PutBuilder(sb)

	sb.Append(s[:i])
	sb.Append(new)
	sb.Append(s[i+len(old):])

	return sb.String()
}

// ReplaceAll is a function that replaces all occurrences of a substring with another substring.
func ReplaceAll(s, old, new string) string {
	if old == "" {
		return s
	}

	sb := GetBuilder()
	defer PutBuilder(sb)

	idx := 0
	for {
		i := Index(s[idx:], old)
		if i == -1 {
			sb.Append(s[idx:])
			break
		}
		sb.Append(s[idx : idx+i])
		sb.Append(new)
		idx += i + len(old)
	}

	return sb.String()
}

// Reverse is a function that reverses a string.
func Reverse(s string) string {
	sb := GetBuilder()
	defer PutBuilder(sb)

	for i := len(s) - 1; i >= 0; i-- {
		sb.AppendByte(s[i])
	}

	return sb.String()
}

// ToTitleCase is a function that converts a string to title case.
func ToTitleCase(s string) string {
	sb := GetBuilder()
	defer PutBuilder(sb)

	shouldCapitalize := true
	for _, r := range s {
		if shouldCapitalize {
			sb.AppendByte(byte(toUpper(r)))
			shouldCapitalize = false
		} else {
			sb.AppendByte(byte(toLower(r)))
		}

		if r == ' ' {
			shouldCapitalize = true
		}
	}

	return sb.String()
}

// ToUpperCase is a function that converts a string to uppercase.
func ToUpperCase(s string) string {
	sb := GetBuilder()
	defer PutBuilder(sb)

	for _, r := range s {
		sb.AppendByte(byte(toUpper(r)))

	}

	return sb.String()
}

// ToLowerCase is a function that converts a string to lowercase.
func ToLowerCase(s string) string {
	sb := GetBuilder()
	defer PutBuilder(sb)

	for _, r := range s {
		sb.AppendByte(byte(toLower(r)))
	}

	return sb.String()
}

// SnakeCase is a function that converts a string to snake case.
func SnakeCase(s string) string {
	sb := GetBuilder()
	defer PutBuilder(sb)

	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			sb.AppendByte('_')
		}
		sb.AppendByte(byte(toLower(r)))
	}

	return sb.String()
}

// CamelCase is a function that converts a string to camel case.
func CamelCase(s string) string {
	sb := GetBuilder()
	defer PutBuilder(sb)

	shouldCapitalize := false
	for _, r := range s {
		if r == ' ' {
			shouldCapitalize = true
		} else if shouldCapitalize {
			sb.AppendByte(byte(toUpper(r)))
			shouldCapitalize = false
		} else {
			sb.AppendByte(byte(toLower(r)))
		}
	}

	return sb.String()
}

// PascalCase is a function that converts a string to pascal case.
func PascalCase(s string) string {
	sb := GetBuilder()
	defer PutBuilder(sb)

	shouldCapitalize := true
	for _, r := range s {
		if r == ' ' {
			shouldCapitalize = true
		} else if shouldCapitalize {
			sb.AppendByte(byte(toUpper(r)))
			shouldCapitalize = false
		} else {
			sb.AppendByte(byte(toLower(r)))
		}
	}

	return sb.String()
}

// KebabCase is a function that converts a string to kebab case.
func KebabCase(s string) string {
	sb := GetBuilder()
	defer PutBuilder(sb)

	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			sb.AppendByte('-')
		}
		sb.AppendByte(byte(toLower(r)))
	}

	return sb.String()
}

// ToogleCase is a function that converts a string to toogle case.
func ToogleCase(s string) string {
	sb := GetBuilder()
	defer PutBuilder(sb)

	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			sb.AppendByte(byte(toLower(r)))
		} else if r >= 'a' && r <= 'z' {
			sb.AppendByte(byte(toUpper(r)))
		} else {
			sb.AppendByte(byte(r))
		}
	}

	return sb.String()
}

// CapitalizeFirst is a function that capitalizes the first letter of a string.
func CapitalizeFirst(s string) string {
	sb := GetBuilder()
	defer PutBuilder(sb)

	if len(s) > 0 {
		sb.AppendByte(byte(toUpper(rune(s[0]))))
		sb.Append(s[1:])
	}

	return sb.String()
}

// ToInt is a function that converts a string to int.
func ToInt(s string) (int, error) {
	return toInt(s)
}

// Join is a function that concatenates a slice of strings into a single string with a separator.
func Join(ss []string, sep string) string {
	sb := GetBuilder()
	defer PutBuilder(sb)

	for i, s := range ss {
		sb.Append(s)
		if i < len(ss)-1 {
			sb.Append(sep)
		}
	}

	return sb.String()
}

// Map is a function that applies a function to each rune in a string and returns a new string.
func Map(s string, f func(rune) rune) string {
	sb := GetBuilder()
	defer PutBuilder(sb)

	for _, r := range s {
		sb.AppendByte(byte(f(r)))
	}

	return sb.String()
}
