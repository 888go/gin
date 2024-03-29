package crypto

import (
	"crypto/hmac"
	"crypto/sha512"
)

const algoHmacSha512 = "hmac-sha512"

// HmacSha512 使用hmac和sha512的签名算法
type HmacSha512 struct{}

// Sign 返回使用秘密字符串对输入msg进行签名的结果
func (h *HmacSha512) Sign(msg string, secret string) ([]byte, error) {
	mac := hmac.New(sha512.New, []byte(secret))
	if _, err := mac.Write([]byte(msg)); err != nil {
		return nil, err
	}
	return mac.Sum(nil), nil
}

// Name 返回算法的名称
func (h *HmacSha512) Name() string {
	return algoHmacSha512
}
