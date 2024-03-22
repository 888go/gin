package httpsign

import (
	"github.com/888go/gin/gin-contrib/httpsign/crypto"
)

// KeyID define type
type KeyID string

// Secret 定义密钥及其使用的算法
type Secret struct {
	Key       string
	Algorithm crypto.Crypto
}

// 密钥映射表，其中keyID为键，secret为值
type Secrets map[KeyID]*Secret
