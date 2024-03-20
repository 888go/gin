package validator

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	
	"github.com/888go/gin"
)

// TODO: 支持更多摘要算法

// ErrInvalidDigest：当body的sha256值与提交的摘要不匹配时，返回该错误
var ErrInvalidDigest = &gin.Error{
	Err:  errors.New("sha256 of body is not match with digest"),
	Type: gin.ErrorTypePublic,
}

// DigestValidator 检查请求头中的摘要是否与主体内容匹配
type DigestValidator struct{}

// NewDigestValidator 返回新的DigestValidator的指针
func NewDigestValidator() *DigestValidator {
	return &DigestValidator{}
}

// Validate在检查摘要与主体内容匹配时返回错误
func (v *DigestValidator) Validate(r *http.Request) error {
	headerDigest := r.Header.Get("digest")
	digest, err := calculateDigest(r)
	if err != nil {
		return err
	}
	if digest != headerDigest {
		return ErrInvalidDigest
	}
	return nil
}

func calculateDigest(r *http.Request) (string, error) {
	if r.ContentLength == 0 {
		return "", nil
	}
	// TODO: 使用缓冲区读取正文以防止占用过多内存
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	h := sha256.New()
	h.Write(body)
	if err != nil {
		return "", err
	}
	digest := fmt.Sprintf("SHA-256=%s", base64.StdEncoding.EncodeToString(h.Sum(nil)))
	return digest, nil
}
