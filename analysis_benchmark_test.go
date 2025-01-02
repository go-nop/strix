package strix_test

import (
	"strings"
	"testing"

	"github.com/go-nop/strix"
)

// //////////////////////
// HasPrefix Benchmark
// //////////////////////
func BenchmarkHasPrefix(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	prefix := "the"
	for i := 0; i < b.N; i++ {
		strix.HasPrefix(s, prefix)
	}
}

func BenchmarkStringsHasPrefix(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	prefix := "the"
	for i := 0; i < b.N; i++ {
		strings.HasPrefix(s, prefix)
	}
}

////////////////////////
// HasSuffix Benchmark
////////////////////////

func BenchmarkHasSuffix(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	suffix := "dog"
	for i := 0; i < b.N; i++ {
		strix.HasSuffix(s, suffix)
	}
}

func BenchmarkStringsHasSuffix(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	suffix := "dog"
	for i := 0; i < b.N; i++ {
		strings.HasSuffix(s, suffix)
	}
}

////////////////////////
// Index Benchmark
////////////////////////

func BenchmarkIndex(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	substr := "fox"
	for i := 0; i < b.N; i++ {
		strix.Index(s, substr)
	}
}

func BenchmarkStringsIndex(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	substr := "fox"
	for i := 0; i < b.N; i++ {
		strings.Index(s, substr)
	}
}

////////////////////////
// IndexAny Benchmark
////////////////////////

func BenchmarkIndexAny(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	chars := "aeiou"
	for i := 0; i < b.N; i++ {
		strix.IndexAny(s, chars)
	}
}

func BenchmarkStringsIndexAny(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	chars := "aeiou"
	for i := 0; i < b.N; i++ {
		strings.IndexAny(s, chars)
	}
}

////////////////////////
// IndexByte Benchmark
////////////////////////

func BenchmarkIndexByte(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	c := byte('o')
	for i := 0; i < b.N; i++ {
		strix.IndexByte(s, c)
	}
}

func BenchmarkStringsIndexByte(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	c := byte('o')
	for i := 0; i < b.N; i++ {
		strings.IndexByte(s, c)
	}
}

////////////////////////
// IndexFunc Benchmark
////////////////////////

func BenchmarkIndexFunc(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	f := func(r rune) bool { return r == 'o' }
	for i := 0; i < b.N; i++ {
		strix.IndexFunc(s, f)
	}
}

func BenchmarkStringsIndexFunc(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	f := func(r rune) bool { return r == 'o' }
	for i := 0; i < b.N; i++ {
		strings.IndexFunc(s, f)
	}
}

////////////////////////
// IndexRune Benchmark
////////////////////////

func BenchmarkIndexRune(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	r := rune('o')
	for i := 0; i < b.N; i++ {
		strix.IndexRune(s, r)
	}
}

func BenchmarkStringsIndexRune(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	r := rune('o')
	for i := 0; i < b.N; i++ {
		strings.IndexRune(s, r)
	}
}

////////////////////////
// LastIndex Benchmark
////////////////////////

func BenchmarkLastIndex(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	substr := "o"
	for i := 0; i < b.N; i++ {
		strix.LastIndex(s, substr)
	}
}

func BenchmarkStringsLastIndex(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	substr := "o"
	for i := 0; i < b.N; i++ {
		strings.LastIndex(s, substr)
	}
}

////////////////////////
// LastIndexAny Benchmark
////////////////////////

func BenchmarkLastIndexAny(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	chars := "aeiou"
	for i := 0; i < b.N; i++ {
		strix.LastIndexAny(s, chars)
	}
}

func BenchmarkStringsLastIndexAny(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	chars := "aeiou"
	for i := 0; i < b.N; i++ {
		strings.LastIndexAny(s, chars)
	}
}

////////////////////////
// LastIndexByte Benchmark
////////////////////////

func BenchmarkLastIndexByte(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	c := byte('o')
	for i := 0; i < b.N; i++ {
		strix.LastIndexByte(s, c)
	}
}

func BenchmarkStringsLastIndexByte(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	c := byte('o')
	for i := 0; i < b.N; i++ {
		strings.LastIndexByte(s, c)
	}
}

////////////////////////
// LastIndexFunc Benchmark
////////////////////////

func BenchmarkLastIndexFunc(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	f := func(r rune) bool { return r == 'o' }
	for i := 0; i < b.N; i++ {
		strix.LastIndexFunc(s, f)
	}
}

func BenchmarkStringsLastIndexFunc(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	f := func(r rune) bool { return r == 'o' }
	for i := 0; i < b.N; i++ {
		strings.LastIndexFunc(s, f)
	}
}
