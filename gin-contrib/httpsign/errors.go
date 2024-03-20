package httpsign

import (
	"errors"
	
	"github.com/888go/gin"
)

func newPublicError(msg string) *gin.Error {
	return &gin.Error{
		Err:  errors.New(msg),
		Type: gin.ErrorTypePublic,
	}
}

var (
	// ErrInvalidAuthorizationHeader 当获取到格式无效的Authorization头时返回错误
	ErrInvalidAuthorizationHeader = newPublicError("Authorization header format is incorrect")
	// ErrInvalidKeyID 当Header中的KeyID未提供时返回错误
	ErrInvalidKeyID = newPublicError("Invalid keyId")
	// ErrDateNotFound 当头部中未找到日期时返回的错误
	ErrDateNotFound = newPublicError("There is no Date on Headers")
	// ErrIncorrectAlgorithm 当头文件中的算法与密钥不匹配时产生的错误
	ErrIncorrectAlgorithm = newPublicError("Algorithm does not match")
	// ErrHeaderNotEnough：当需求头信息没有出现在头部字段时返回的错误
	ErrHeaderNotEnough = newPublicError("Header field is not match requirement")
	// ErrNoSignature 当在头部未找到签名时返回的错误
	ErrNoSignature = newPublicError("No Signature header found in request")
	// ErrInvalidSign 当签名字符串不匹配时产生的错误
	ErrInvalidSign = newPublicError("Invalid sign")
	// ErrMissingKeyID 当keyId未在header中时产生的错误
	ErrMissingKeyID = newPublicError("keyId must be on header")
	// ErrMissingSignature 当签名不在头部时返回错误
	ErrMissingSignature = newPublicError("signature must be on header")

	// ErrUnterminatedParameter 当无法解析值时产生的错误
	ErrUnterminatedParameter = newPublicError("Unterminated parameter")
	// ErrMisingDoubleQuote 当字符 = 后面缺少双引号时的错误
	ErrMisingDoubleQuote = newPublicError(`Missing " after = character`)
	// ErrMisingEqualCharacter 当在 " 或 , 字符前缺少等号（=）时的错误
	ErrMisingEqualCharacter = newPublicError(`Missing = character =`)
)
