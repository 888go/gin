package httpsign

import (
	"net/http"
	"strings"
)

const (
	authorizationHeader           = "Authorization"
	authorizationHeaderInitString = "Signature "
	signatureHeader               = "Signature"
	signingKeyID                  = "keyId"
	signingAlgorithm              = "algorithm"
	signingSignature              = "signature"
	signingHeaders                = "headers"
)

// SignatureHeader 包含签名头的基本信息
type SignatureHeader struct {
	keyID     KeyID
	headers   []string
	signature string
	algorithm string
}

// NewSignatureHeader 创建一个新的 SignatureHeader 实例
func NewSignatureHeader(r *http.Request) (*SignatureHeader, error) {
	return parseHTTPRequest(r)
}

func parseHTTPRequest(r *http.Request) (*SignatureHeader, error) {
	s, err := getSignatureString(r)
	if err != nil {
		return nil, err
	}
	return parseSignatureString(s)
}

func parseSignatureString(s string) (*SignatureHeader, error) {
	p := newParser(s)
	results, err := p.parse()
	if err != nil {
		return nil, err
	}
	keyID, ok := results[signingKeyID]
	if !ok {
		return nil, ErrMissingKeyID
	}
	signature, ok := results[signingSignature]
	if !ok {
		return nil, ErrMissingSignature
	}
	headerString, ok := results[signingHeaders]
	var headers []string
	if !ok || len(headerString) == 0 {
		headers = []string{"date"}
	} else {
		headers = strings.Split(headerString, " ")
	}

	algorithm := results[signingAlgorithm]

	return &SignatureHeader{
		keyID:     KeyID(keyID),
		signature: signature,
		headers:   headers,
		algorithm: algorithm,
	}, nil
}

func getSignatureString(r *http.Request) (string, error) {
	s := r.Header.Get(signatureHeader)
	if s != "" {
		return s, nil
	}

	s = r.Header.Get(authorizationHeader)
	if s != "" {
		if strings.Index(s, authorizationHeaderInitString) != 0 {
			return "", ErrInvalidAuthorizationHeader
		}
		return strings.TrimPrefix(s, authorizationHeaderInitString), nil
	}

	return "", ErrNoSignature
}
