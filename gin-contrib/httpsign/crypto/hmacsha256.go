package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
)

const algoHmacSha256 = "hmac-sha256"

// HmacSha256：使用hmac和sha256的签名算法
type HmacSha256 struct{}

// Sign 返回使用秘密字符串对输入msg进行签名的结果
func (h *HmacSha256) Sign(msg string, secret string) ([]byte, error) {
	mac := hmac.New(sha256.New, []byte(secret))
	if _, err := mac.Write([]byte(msg)); err != nil {
		return nil, err
	}
	return mac.Sum(nil), nil
}

// Name 返回算法名称
func (h *HmacSha256) Name() string {
	return algoHmacSha256
}
