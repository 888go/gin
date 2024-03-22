package limits

import (
	"fmt"
	"io"
	"net/http"
	
	"github.com/888go/gin"
)

type maxBytesReader struct {
	ctx        *gin类.Context
	rdr        io.ReadCloser
	remaining  int64
	wasAborted bool
	sawEOF     bool
}

func (mbr *maxBytesReader) tooLarge() (n int, err error) {
	n, err = 0, fmt.Errorf("HTTP request too large")

	if !mbr.wasAborted {
		mbr.wasAborted = true
		ctx := mbr.ctx
		_ = ctx.X错误(err)
		ctx.X设置响应协议头值("connection", "close")
		ctx.X输出文本(http.StatusRequestEntityTooLarge, "request too large")
		ctx.X停止并带状态码(http.StatusRequestEntityTooLarge)
	}
	return
}

func (mbr *maxBytesReader) Read(p []byte) (n int, err error) {
	toRead := mbr.remaining
	if mbr.remaining == 0 {
		if mbr.sawEOF {
			return mbr.tooLarge()
		}
		// The underlying io.Reader may not return (0, io.EOF)
		// at EOF if the requested size is 0, so read 1 byte
		// instead. The io.Reader docs are a bit ambiguous
		// about the return value of Read when 0 bytes are
		// requested, and {bytes,strings}.Reader gets it wrong
		// too (it returns (0, nil) even at EOF).
		toRead = 1
	}
	if int64(len(p)) > toRead {
		p = p[:toRead]
	}
	n, err = mbr.rdr.Read(p)
	if err == io.EOF {
		mbr.sawEOF = true
	}
	if mbr.remaining == 0 {
		// If we had zero bytes to read remaining (but hadn't seen EOF)
		// and we get a byte here, that means we went over our limit.
		if n > 0 {
			return mbr.tooLarge()
		}
		return 0, err
	}
	mbr.remaining -= int64(n)
	if mbr.remaining < 0 {
		mbr.remaining = 0
	}
	return n, err
}

func (mbr *maxBytesReader) Close() error {
	return mbr.rdr.Close()
}

// RequestSizeLimiter returns a middleware that limits the size of request
// When a request is over the limit, the following will happen:
// * Error will be added to the context
// * Connection: close header will be set
// * Error 413 will be sent to the client (http.StatusRequestEntityTooLarge)
// * Current context will be aborted
func RequestSizeLimiter(limit int64) gin类.HandlerFunc {
	return func(ctx *gin类.Context) {
		ctx.X请求.Body = &maxBytesReader{
			ctx:        ctx,
			rdr:        ctx.X请求.Body,
			remaining:  limit,
			wasAborted: false,
			sawEOF:     false,
		}
		ctx.X中间件继续()
	}
}
