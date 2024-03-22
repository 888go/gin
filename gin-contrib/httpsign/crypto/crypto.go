package crypto

// Crypto接口，用于签名算法
type Crypto interface {
	Name() string
	Sign(msg string, secret string) ([]byte, error)
}
