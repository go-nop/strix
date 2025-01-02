package strix_test

import (
	"strings"
	"testing"

	"github.com/go-nop/strix"
)

// //////////////////////
// Clone Benchmark
// //////////////////////
func BenchmarkClone(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	for i := 0; i < b.N; i++ {
		strix.Clone(s)
	}
}

func BenchmarkStringsClone(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	for i := 0; i < b.N; i++ {
		strings.Clone(s)
	}
}

// //////////////////////
// Count Benchmark
// //////////////////////

func BenchmarkCount(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	substr := "o"
	for i := 0; i < b.N; i++ {
		strix.Count(s, substr)
	}
}

func BenchmarkStringsCount(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	substr := "o"
	for i := 0; i < b.N; i++ {
		strings.Count(s, substr)
	}
}

// //////////////////////
// Fields Benchmark
// //////////////////////

func BenchmarkFields(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	for i := 0; i < b.N; i++ {
		strix.Fields(s)
	}
}

func BenchmarkStringsFields(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	for i := 0; i < b.N; i++ {
		strings.Fields(s)
	}
}

// //////////////////////
// FieldsFunc Benchmark
// //////////////////////

func BenchmarkFieldsFunc(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	f := func(r byte) bool { return r == ' ' }
	for i := 0; i < b.N; i++ {
		strix.FieldsFunc(s, f)
	}
}

func BenchmarkStringsFieldsFunc(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	f := func(r rune) bool { return r == ' ' }
	for i := 0; i < b.N; i++ {
		strings.FieldsFunc(s, f)
	}
}

// //////////////////////
// Repeat Benchmark
// //////////////////////

func BenchmarkRepeat(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	n := 10
	for i := 0; i < b.N; i++ {
		strix.Repeat(s, n)
	}
}

func BenchmarkStringsRepeat(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	n := 10
	for i := 0; i < b.N; i++ {
		strings.Repeat(s, n)
	}
}
