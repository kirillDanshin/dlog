// +build debug

package dlog

import "fmt"

// Ln is a build-time enabled println
func (b *Buffered) Ln(v ...interface{}) {
	b.prepare()
	b.Lock()
	fmt.Fprintln(b.bb, v...)
	b.Unlock()
}

// P is a build-time enabled print
func (b *Buffered) P(v ...interface{}) {
	b.prepare()
	b.Lock()
	fmt.Fprint(b.bb, v...)
	b.Unlock()
}

// F is a build-time enabled printf
func (b *Buffered) F(f string, v ...interface{}) {
	b.prepare()
	b.Lock()
	fmt.Fprintf(b.bb, f, v...)
	b.Unlock()
}

// D dumps a value
func (b *Buffered) D(v ...interface{}) {
	b.prepare()
	b.Lock()
	for _, v := range v {
		fmt.Fprintf(b.bb, "%+#v", v)
	}
	b.Unlock()
}

func (b *Buffered) prepare() {
	if b.bb == nil {
		b = NewBuffered()
	}
}
