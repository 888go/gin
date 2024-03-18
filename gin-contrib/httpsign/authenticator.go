package httpsign

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	
	"github.com/888go/gin/gin-contrib/httpsign/validator"
	"github.com/888go/gin"
)

const (
	requestTarget = "(request-target)"
	date          = "date"
	digest        = "digest"
	host          = "host"
)

var defaultRequiredHeaders = []string{requestTarget, date, digest}

// Authenticator 是 Gin 框架的身份验证中间件。
type Authenticator struct {
	secrets    Secrets
	validators []validator.Validator
	headers    []string
}

// Option 是 Authenticator 构造函数的选项。
type Option func(*Authenticator)

// WithValidator 配置 Authenticator 以使用自定义验证器。
// 默认的验证器基于时间和摘要。

// ff:
// validators:

// ff:
// validators:

// ff:
// validators:

// ff:
// validators:

// ff:
// validators:

// ff:
// validators:

// ff:
// validators:
func WithValidator(validators ...validator.Validator) Option {
	return func(a *Authenticator) {
		a.validators = validators
	}
}

// WithRequiredHeaders 是一个包含所有必需HTTP头的列表，客户端
// 必须在签名字符串中包含这些头信息，以便请求被认为是有效的。
// 如果未提供，创建的Authenticator实例将使用默认的defaultRequiredHeaders变量。

// ff:
// headers:

// ff:
// headers:

// ff:
// headers:

// ff:
// headers:

// ff:
// headers:

// ff:
// headers:

// ff:
// headers:
func WithRequiredHeaders(headers []string) Option {
	return func(a *Authenticator) {
		a.headers = headers
	}
}

// NewAuthenticator 创建一个全新的 Authenticator 实例，其参数包括
// 允许的权限列表以及请求头中所需的密钥和秘密键。

// ff:
// options:
// secretKeys:

// ff:
// options:
// secretKeys:

// ff:
// options:
// secretKeys:

// ff:
// options:
// secretKeys:

// ff:
// options:
// secretKeys:

// ff:
// options:
// secretKeys:

// ff:
// options:
// secretKeys:
func NewAuthenticator(secretKeys Secrets, options ...Option) *Authenticator {
	a := &Authenticator{secrets: secretKeys}

	for _, fn := range options {
		fn(a)
	}

	if a.validators == nil {
		a.validators = []validator.Validator{
			validator.NewDateValidator(),
			validator.NewDigestValidator(),
		}
	}

	if len(a.headers) == 0 {
		a.headers = defaultRequiredHeaders
	}

	return a
}

// Authenticated 返回一个 gin 中间件，该中间件允许在参数中指定的权限。

// ff:

// ff:

// ff:

// ff:

// ff:

// ff:

// ff:
func (a *Authenticator) Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		sigHeader, err := NewSignatureHeader(c.Request)
		if err != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		for _, v := range a.validators {
			if err := v.Validate(c.Request); err != nil {
				_ = c.AbortWithError(http.StatusBadRequest, err)
				return
			}
		}
		if !a.isValidHeader(sigHeader.headers) {
			_ = c.AbortWithError(http.StatusBadRequest, ErrHeaderNotEnough)
			return
		}

		secret, err := a.getSecret(sigHeader.keyID, sigHeader.algorithm)
		if err != nil {
			if err == ErrInvalidKeyID {
				_ = c.AbortWithError(http.StatusUnauthorized, err)
				return
			}
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		signString := constructSignMessage(c.Request, sigHeader.headers)
		signature, err := secret.Algorithm.Sign(signString, secret.Key)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		signatureBase64 := base64.StdEncoding.EncodeToString(signature)
		if signatureBase64 != sigHeader.signature {
			_ = c.AbortWithError(http.StatusUnauthorized, ErrInvalidSign)
			return
		}
		c.Next()
	}
}

// isValidHeader 检查是否所有服务器要求的头部都在头部列表中
func (a *Authenticator) isValidHeader(headers []string) bool {
	m := len(headers)
	for _, h := range a.headers {
		i := 0
		for i = 0; i < m; i++ {
			if h == headers[i] {
				break
			}
		}
		if i == m {
			return false
		}
	}
	return true
}

func (a *Authenticator) getSecret(keyID KeyID, algorithm string) (*Secret, error) {
	secret, ok := a.secrets[keyID]
	if !ok {
		return nil, ErrInvalidKeyID
	}

	if secret.Algorithm.Name() != algorithm {
		if algorithm != "" {
			return nil, ErrIncorrectAlgorithm
		}
	}
	return secret, nil
}

func constructSignMessage(r *http.Request, headers []string) string {
	var signBuffer bytes.Buffer
	for i, field := range headers {
		var fieldValue string
		switch field {
		case host:
			fieldValue = r.Host
		case requestTarget:
			fieldValue = fmt.Sprintf("%s %s", strings.ToLower(r.Method), r.URL.RequestURI())
		default:
			fieldValue = r.Header.Get(field)
		}
		signString := fmt.Sprintf("%s: %s", field, fieldValue)
		signBuffer.WriteString(signString)
		if i < len(headers)-1 {
			signBuffer.WriteString("\n")
		}
	}
	return signBuffer.String()
}
