package main

// CircularBuffer implements a circular float64 buffer. It is a fixed size,
// and new writes overwrite older data, such that for a buffer
// of size N, for any amount of writes, only the last N values
// are retained.
type CircularBuffer struct {
	data         []float64
	size         int64
	writeCursor  int64
	writeOpCount int64
}

// NewCircularBuffer creates a new buffer of a given size
func NewCircularBuffer(size int64) *CircularBuffer {
	if size <= 0 {
		panic("Size must be positive")
	}

	b := &CircularBuffer{
		size: size,
		data: make([]float64, size),
	}
	return b
}

// Add adds new value to the buffer, overriding old data if necessary.
func (b *CircularBuffer) Add(value float64) error {
	b.data[b.writeCursor] = value
	b.writeCursor = ((b.writeCursor + 1) % b.size)
	b.writeOpCount++
	return nil
}

// Size returns the size of the buffer
func (b *CircularBuffer) Size() int64 {
	return b.size
}

// TotalWriteOpCount provides the total number of values written
func (b *CircularBuffer) TotalWriteOpCount() int64 {
	return b.writeOpCount
}

// Data returns ordered data from buffer, from old to new.
func (b *CircularBuffer) Data() []float64 {
	switch {
	case b.writeOpCount >= b.size && b.writeCursor == 0:
		out := make([]float64, b.size)
		copy(out, b.data)
		return out
	case b.writeOpCount > b.size:
		out := make([]float64, b.size)
		copy(out, b.data[b.writeCursor:])
		copy(out[b.size-b.writeCursor:], b.data[:b.writeCursor])
		return out
	default:
		out := make([]float64, b.writeCursor)
		copy(out, b.data[:b.writeCursor])
		return out
	}
}

// Reset resets the buffer so it has no content.
func (b *CircularBuffer) Reset() {
	b.writeCursor = 0
	b.writeOpCount = 0
}
