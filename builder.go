package strix

import (
	"errors"
	"sync"
)

// Builder is a struct that represents a string Builder.
type Builder struct {
	buffer []byte
}

var builderPool = sync.Pool{
	New: func() interface{} {
		return &Builder{
			buffer: make([]byte, 0, 64), // Initial capacity of 64 bytes
		}
	},
}

// GetBuilder is a function that gets a Builder from the pool.
func GetBuilder() *Builder {
	return builderPool.Get().(*Builder)
}

// PutBuilder is a function that puts the Builder back into the pool.
func PutBuilder(b *Builder) {
	b.Reset() // Reset buffer
	builderPool.Put(b)
}

// Append menambahkan string ke buffer.
func (b *Builder) Append(s string) {
	b.buffer = append(b.buffer, s...)
}

// AppendByte is a function that appends a byte to the buffer.
func (b *Builder) AppendByte(c byte) {
	b.buffer = append(b.buffer, c)
}

// Delete is a function that deletes a substring from the buffer.
func (b *Builder) Delete(start, end int) {
	b.buffer = append(b.buffer[:start], b.buffer[end:]...)
}

// String is a function that returns the string representation of the buffer.
func (b *Builder) String() string {
	return string(b.buffer)
}

// Reset is a function that resets the buffer.
func (b *Builder) Reset() {
	b.buffer = b.buffer[:0]
}

// Len is a function that returns the length of the buffer.
func (b *Builder) Len() int {
	return len(b.buffer)
}

// Cap is a function that returns the capacity of the buffer.
func (b *Builder) Cap() int {
	return cap(b.buffer)
}

// Grow is a function that increases the capacity of the buffer.
func (b *Builder) Grow(n int) error {
	if n < 0 {
		return errors.New("grow: negative size not allowed")
	}
	if n > cap(b.buffer) {
		newBuffer := make([]byte, len(b.buffer), cap(b.buffer)+n)
		copy(newBuffer, b.buffer)
		b.buffer = newBuffer
	}
	return nil
}
