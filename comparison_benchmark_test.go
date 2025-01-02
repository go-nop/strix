package strix_test

import (
	"strings"
	"testing"

	"github.com/go-nop/strix"
)

// //////////////////////
// Compare Benchmark
// //////////////////////
func BenchmarkCompare(b *testing.B) {
	s1 := "hello"
	s2 := "world"
	for i := 0; i < b.N; i++ {
		strix.Compare(s1, s2)
	}
}

func BenchmarkStringsCompare(b *testing.B) {
	s1 := "hello"
	s2 := "world"
	for i := 0; i < b.N; i++ {
		strings.Compare(s1, s2)
	}
}

// //////////////////////
// Contains Benchmark
// //////////////////////

func BenchmarkContains(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	substr := "fox"
	for i := 0; i < b.N; i++ {
		strix.Contains(s, substr)
	}
}

func BenchmarkStringsContains(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	substr := "fox"
	for i := 0; i < b.N; i++ {
		strings.Contains(s, substr)
	}
}

// //////////////////////
// ContainsAny Benchmark
// //////////////////////

func BenchmarkContainsAny(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	substrs := []string{"fox", "dog"}
	for i := 0; i < b.N; i++ {
		strix.ContainsAny(s, substrs...)
	}
}

func BenchmarkStringsContainsAny(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	substrs := "foxdog"
	for i := 0; i < b.N; i++ {
		strings.ContainsAny(s, substrs)
	}
}

// //////////////////////
// ContainsFunc Benchmark
// //////////////////////

func BenchmarkContainsFunc(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	f := func(r rune) bool { return r == 'o' }
	for i := 0; i < b.N; i++ {
		strix.ContainsFunc(s, f)
	}
}

func BenchmarkStringsContainsFunc(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	f := func(r rune) bool { return r == 'o' }
	for i := 0; i < b.N; i++ {
		strings.ContainsFunc(s, f)
	}
}

// //////////////////////
// ContainsRune Benchmark
// //////////////////////

func BenchmarkContainsRune(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	r := rune('o')
	for i := 0; i < b.N; i++ {
		strix.ContainsRune(s, r)
	}
}

func BenchmarkStringsContainsRune(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	r := rune('o')
	for i := 0; i < b.N; i++ {
		strings.ContainsRune(s, r)
	}
}

// //////////////////////
// ContainsAll Benchmark
// //////////////////////

func BenchmarkEqualFold(b *testing.B) {
	s1 := "hello"
	s2 := "HELLO"
	for i := 0; i < b.N; i++ {
		strix.EqualFold(s1, s2)
	}
}

func BenchmarkStringsEqualFold(b *testing.B) {
	s1 := "hello"
	s2 := "HELLO"
	for i := 0; i < b.N; i++ {
		strings.EqualFold(s1, s2)
	}
}
