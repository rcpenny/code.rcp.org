package rc_strings

import "unsafe"

// A Builder is used to efficiently build a string using Write methods.
// It minimizes memory copying. The zero value is ready to use.
// Do not copy a non-zero Builder.
type Builder struct {
	addr *Builder // of receiver, to detect copies by value
	buf  []byte
}

func (b *Builder) copyCheck() {
	if b.addr == nil {
		b.addr = b
	} else if b.addr != b {
		panic("rc_strings: illegal use of non-zero Builder copied by value")
	}
}

// String returns the accumulated string.
func (b *Builder) String() string {
	// TODO Daniel
	return *(*string)(unsafe.Pointer(&b.buf))
}

// Len returns the number of accumulated bytes; b.Len() == len(b.String()).
func (b *Builder) Len() int {
	return len(b.buf)
}

// Cap returns the capacity of the builder's underlying byte slice. It is the
// total space allocated for the string being built and includes any bytes
// already written.
func (b *Builder) Cap() int {
	return cap(b.buf)
}

// Reset resets the Builder to be empty.
func (b *Builder) Reset() {
	b.addr = nil
	b.buf = nil
}

// grow copies the bufer to a new, larger buffer so that there are at least n
// bytes of capacity beyond len(b.buf).
func (b *Builder) grow(n int) {
	buf := make([]byte, len(b.buf), 2*cap(b.buf)+n)
	copy(buf, b.buf)
	b.buf = buf
}

// Grow grows b's capacity, if necessary, to guarantee space for
// another n bytes. After Grow(n), at least n bytes can be written to b
// without another allocation. If n is negative,Grow panics.
func (b *Builder) Grow(n int) {
	b.copyCheck()
	if n < 0 {
		panic("rc_strings.Builder.Grow: negative count")
	}
	if cap(b.buf)-len(b.buf) < n {
		b.grow(n)
	}
}

// Write appends the contents of p to b's buffer
// Write always returns len(p), nil.
func (b *Builder) Write(p []byte) (int, error) {
	b.copyCheck()
	b.buf = append(b.buf, p...)
	return len(p), nil
}

// WriteByte appends the byte c to b's buffer.
// The returned error is always nil.
func (b *Builder) WriteByte(c byte) error {
	b.copyCheck()
	b.buf = append(b.buf, c)
	return nil
}

// WriteString appends the content of s to b's buffer.
// It returns the length of s and a nil error
func (b *Builder) WriteString(s string) (int, error) {
	b.copyCheck()
	b.buf = append(b.buf, s...)
	return len(s), nil
}
