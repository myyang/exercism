package circular

const testVersion = 3

// NewBuffer return Buffer type
func NewBuffer(n int) *Buffer {
	return &Buffer{wptr: 0, rptr: 0, arr: make([]byte, n)}
}

// Buffer defines structure
type Buffer struct {
	wptr, rptr int
	arr        []byte
}

// BufferError defines naive buffer error
type BufferError struct{}

// Error naive message
func (berr BufferError) Error() string {
	return "Buffer error"
}

// WriteByte into buffer
func (b *Buffer) WriteByte(c byte) error {
	if b.arr[b.wptr] != 0 {
		return BufferError{}
	}
	b.arr[b.wptr] = c
	b.wptr = (b.wptr + 1) % len(b.arr)
	return nil
}

// ReadByte from buffer
func (b *Buffer) ReadByte() (byte, error) {
	if b.arr[b.rptr] == 0 {
		return 0, BufferError{}
	}
	x := b.arr[b.rptr]
	b.arr[b.rptr] = 0
	b.rptr = (b.rptr + 1) % len(b.arr)
	return x, nil
}

// Reset buffer
func (b *Buffer) Reset() {
	b.arr = make([]byte, len(b.arr))
	b.rptr, b.wptr = 0, 0
}

// Overwrite buufer
func (b *Buffer) Overwrite(c byte) {
	err := b.WriteByte(c)
	if err != nil {
		b.arr[b.rptr] = c
		b.rptr = (b.rptr + 1) % len(b.arr)
	}
}
