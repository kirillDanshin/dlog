package dlog

import (
	"sync"

	"github.com/valyala/bytebufferpool"
)

// Buffered thread-safe dlog
type Buffered struct {
	bb *bytebufferpool.ByteBuffer
	sync.RWMutex
}

// NewBuffered dlog
func NewBuffered() *Buffered {
	return &Buffered{
		bb: bytebufferpool.Get(),
	}
}

// Release the buffer for dlog
func (b *Buffered) Release() {
	if b.bb == nil {
		return
	}
	b.Lock()
	bytebufferpool.Put(b.bb)
	b.Unlock()
}
