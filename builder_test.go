package strix_test

import (
	"testing"

	"github.com/go-nop/strix"
)

func TestBuilderAppend(t *testing.T) {
	b := strix.GetBuilder()
	defer strix.PutBuilder(b)

	b.Append("hello")
	b.Append(" ")
	b.Append("world")

	expected := "hello world"
	if result := b.String(); result != expected {
		t.Errorf("Builder.Append() = %v; want %v", result, expected)
	}
}

func TestBuilderAppendByte(t *testing.T) {
	b := strix.GetBuilder()
	defer strix.PutBuilder(b)

	b.AppendByte('h')
	b.AppendByte('e')
	b.AppendByte('l')
	b.AppendByte('l')
	b.AppendByte('o')

	expected := "hello"
	if result := b.String(); result != expected {
		t.Errorf("Builder.AppendByte() = %v; want %v", result, expected)
	}
}

func TestBuilderDelete(t *testing.T) {
	b := strix.GetBuilder()
	defer strix.PutBuilder(b)

	b.Append("hello world")
	b.Delete(5, 6)

	expected := "helloworld"
	if result := b.String(); result != expected {
		t.Errorf("Builder.Delete() = %v; want %v", result, expected)
	}
}

func TestBuilderReset(t *testing.T) {
	b := strix.GetBuilder()
	defer strix.PutBuilder(b)

	b.Append("hello world")
	b.Reset()

	expected := ""
	if result := b.String(); result != expected {
		t.Errorf("Builder.Reset() = %v; want %v", result, expected)
	}
}

func TestBuilderLen(t *testing.T) {
	b := strix.GetBuilder()
	defer strix.PutBuilder(b)

	b.Append("hello world")

	expected := 11
	if result := b.Len(); result != expected {
		t.Errorf("Builder.Len() = %v; want %v", result, expected)
	}
}

func TestBuilderCap(t *testing.T) {
	b := strix.GetBuilder()
	defer strix.PutBuilder(b)

	expected := 64
	if result := b.Cap(); result != expected {
		t.Errorf("Builder.Cap() = %v; want %v", result, expected)
	}
}

func TestBuilderGrow(t *testing.T) {
	b := strix.GetBuilder()
	defer strix.PutBuilder(b)

	err := b.Grow(128)
	if err != nil {
		t.Errorf("Builder.Grow() error = %v; want nil", err)
	}

	expected := 192
	if result := b.Cap(); result != expected {
		t.Errorf("Builder.Cap() after Grow() = %v; want %v", result, expected)
	}
}
