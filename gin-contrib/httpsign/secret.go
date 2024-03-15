package httpsign

import (
	"github.com/888go/gin/gin-contrib/httpsign/crypto"
)

// KeyID 定义类型
type KeyID string

// Secret 定义密钥及其使用的算法
type Secret struct {
	Key       string
	Algorithm crypto.Crypto
}

// Secrets：使用keyID作为键、秘密信息作为值的映射（字典）
type Secrets map[KeyID]*Secret
