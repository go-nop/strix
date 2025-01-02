package strix

// Clone is a function that creates a copy of a string.
func Clone(s string) string {
	return s
}

// Count is a function that counts the number of non-overlapping instances of a substring in a string.
func Count(s, substr string) int {
	if len(substr) == 0 {
		return 0
	}
	var count int
	for i := 0; i <= len(s)-len(substr); {
		if s[i:i+len(substr)] == substr {
			count++
			i += len(substr)
		} else {
			i++
		}
	}
	return count
}

// Fields is a function that splits a string into a slice of fields.
func Fields(s string) []string {
	var result []string
	start := -1
	wordCount := 0

	// Menghitung jumlah kata
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			if start != -1 {
				wordCount++
				start = -1
			}
		} else {
			if start == -1 {
				start = i
			}
		}
	}
	if start != -1 {
		wordCount++
	}

	result = make([]string, 0, wordCount)

	start = -1
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			if start != -1 {
				result = append(result, s[start:i])
				start = -1
			}
		} else {
			if start == -1 {
				start = i
			}
		}
	}
	if start != -1 {
		result = append(result, s[start:])
	}

	return result
}

// FieldsFunc is a function that splits a string into a slice of fields using a custom function.
func FieldsFunc(s string, f func(byte) bool) []string {
	var result []string
	var start, end int
	for start < len(s) {
		for start < len(s) && f(byte(s[start])) {
			start++
		}
		end = start
		for end < len(s) && !f(byte(s[end])) {
			end++
		}
		if start < end {
			result = append(result, s[start:end])
		}
		start = end
	}
	return result
}

// Repeat is a function that repeats a string n times.
func Repeat(s string, n int) string {
	totalLength := len(s) * n

	sb := GetBuilder()
	defer PutBuilder(sb)

	_ = sb.Grow(totalLength)

	for i := 0; i < n; i++ {
		sb.Append(s)
	}

	return sb.String()
}
