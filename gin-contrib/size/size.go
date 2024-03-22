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
// 当请求的大小为0时，底层的io.Reader在遇到EOF（文件结束）时可能不会返回(0, io.EOF)，因此改为读取1个字节。关于在请求0字节时Read方法的返回值，io.Reader的文档有些模糊不清。实际上，{bytes,strings}.Reader也处理得不正确（即使在EOF时，它也会返回(0, nil)）。
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
// 如果我们之前剩余0字节可读（但尚未遇到EOF）
// 而在这里获取到一个字节，则表示我们超过了限制。
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

// RequestSizeLimiter 返回一个中间件，用于限制请求的大小
// 当请求超过限制时，会发生以下情况：
// * 将错误添加到上下文中
// * 设置 "Connection: close" 头部
// * 向客户端发送 413 错误（http.StatusRequestEntityTooLarge，请求实体过大）
// * 中断当前上下文
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
