// +build !debug

package dlog

// Ln is a build-time enabled println
func (b *Buffered) Ln(v ...interface{}) {}

// P is a build-time enabled print
func (b *Buffered) P(v ...interface{}) {}

// F is a build-time enabled printf
func (b *Buffered) F(f string, v ...interface{}) {}

// D dumps a value
func (b *Buffered) D(v ...interface{}) {}

func (b *Buffered) prepare() {}

func (b *Buffered) String() string {
	return ""
}
