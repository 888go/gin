package timeout

import (
	"bytes"
	"sync"
)

// BufferPool 表示一个缓冲区池。
type BufferPool struct {
	pool sync.Pool
}

// Get 从缓冲池返回一个缓冲区。
// 如果缓冲池为空，则创建并返回一个新的缓冲区。
func (p *BufferPool) Get() *bytes.Buffer {
	buf := p.pool.Get()
	if buf == nil {
		return &bytes.Buffer{}
	}
	return buf.(*bytes.Buffer)
}

// Put 将缓冲区放回池中。
func (p *BufferPool) Put(buf *bytes.Buffer) {
	p.pool.Put(buf)
}
