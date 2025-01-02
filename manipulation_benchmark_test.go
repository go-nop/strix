package strix_test

import (
	"strings"
	"testing"

	"github.com/go-nop/strix"
)

////////////////////////
// Split Benchmark
////////////////////////

func BenchmarkSplit(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	sep := " "

	for i := 0; i < b.N; i++ {
		strix.Split(s, sep)
	}
}

func BenchmarkStringsSplit(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	sep := " "

	for i := 0; i < b.N; i++ {
		strings.Split(s, sep)
	}
}

////////////////////////
// TrimSpace Benchmark
////////////////////////

func BenchmarkTrimSpace(b *testing.B) {
	s := " the quick brown fox jumps over the lazy dog "

	for i := 0; i < b.N; i++ {
		strix.TrimSpace(s)
	}
}

func BenchmarkStringsTrimSpace(b *testing.B) {
	s := " the quick brown fox jumps over the lazy dog "

	for i := 0; i < b.N; i++ {
		strings.TrimSpace(s)
	}
}

////////////////////////
// ReplaceAll Benchmark
////////////////////////

func BenchmarkReplaceAll(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	old := "the"
	new := "a"

	for i := 0; i < b.N; i++ {
		strix.ReplaceAll(s, old, new)
	}
}

func BenchmarkStringsReplaceAll(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	old := "the"
	new := "a"

	for i := 0; i < b.N; i++ {
		strings.ReplaceAll(s, old, new)
	}
}

////////////////////////
// Reverse Benchmark
////////////////////////

func BenchmarkReverse(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"

	for i := 0; i < b.N; i++ {
		strix.Reverse(s)
	}
}

func BenchmarkStringsReverse(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"

	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		for i := len(s) - 1; i >= 0; i-- {
			sb.WriteByte(s[i])
		}
	}
}

////////////////////////
// ToTitleCase Benchmark
////////////////////////

func BenchmarkToTitleCase(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"

	for i := 0; i < b.N; i++ {
		strix.ToTitleCase(s)
	}
}

func BenchmarkStringsToTitle(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"

	for i := 0; i < b.N; i++ {
		strings.ToTitle(s)
	}
}

////////////////////////
// ToUpperCase Benchmark
////////////////////////

func BenchmarkToUpperCase(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"

	for i := 0; i < b.N; i++ {
		strix.ToUpperCase(s)
	}
}

func BenchmarkStringsToUpper(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"

	for i := 0; i < b.N; i++ {
		strings.ToUpper(s)
	}
}

////////////////////////
// ToLowerCase Benchmark
////////////////////////

func BenchmarkToLowerCase(b *testing.B) {
	s := "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"

	for i := 0; i < b.N; i++ {
		strix.ToLowerCase(s)
	}
}

func BenchmarkStringsToLower(b *testing.B) {
	s := "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"

	for i := 0; i < b.N; i++ {
		strings.ToLower(s)
	}
}
