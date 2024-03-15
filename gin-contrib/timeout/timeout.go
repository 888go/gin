package timeout

import (
	"time"
	
	"github.com/888go/gin"
)

var bufPool *BufferPool

const (
	defaultTimeout = 5 * time.Second
)

// New 包装一个处理器，如果达到超时时间，则中止处理器的执行流程
func New(opts ...Option) gin.HandlerFunc {
	t := &Timeout{
		timeout:  defaultTimeout,
		handler:  nil,
		response: defaultResponse,
	}

// 遍历每个选项
	for _, opt := range opts {
		if opt == nil {
			panic("timeout Option not be nil")
		}

// 调用选项，传入已实例化的
		opt(t)
	}

	if t.timeout <= 0 {
		return t.handler
	}

	bufPool = &BufferPool{}

	return func(c *gin.Context) {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		w := c.Writer
		buffer := bufPool.Get()
		tw := NewWriter(w, buffer)
		c.Writer = tw
		buffer.Reset()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			t.handler(c)
			finish <- struct{}{}
		}()

		select {
		case p := <-panicChan:
			tw.FreeBuffer()
			c.Writer = w
			panic(p)

		case <-finish:
			c.Next()
			tw.mu.Lock()
			defer tw.mu.Unlock()
			dst := tw.ResponseWriter.Header()
			for k, vv := range tw.Header() {
				dst[k] = vv
			}

			if _, err := tw.ResponseWriter.Write(buffer.Bytes()); err != nil {
				panic(err)
			}
			tw.FreeBuffer()
			bufPool.Put(buffer)

		case <-time.After(t.timeout):
			c.Abort()
			tw.mu.Lock()
			defer tw.mu.Unlock()
			tw.timeout = true
			tw.FreeBuffer()
			bufPool.Put(buffer)

			c.Writer = w
			t.response(c)
			c.Writer = tw
		}
	}
}
