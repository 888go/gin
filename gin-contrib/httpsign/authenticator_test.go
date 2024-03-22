package httpsign

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	
	"github.com/888go/gin/gin-contrib/httpsign/crypto"
	"github.com/888go/gin/gin-contrib/httpsign/validator"
	
	"github.com/stretchr/testify/require"
	
	"github.com/888go/gin"
	"github.com/888go/gin/render"
	"github.com/stretchr/testify/assert"
)

const (
	readID                 = KeyID("read")
	writeID                = KeyID("write")
	invalidKeyID           = KeyID("invalid key")
	invaldAlgo             = "invalidAlgo"
	requestNilBodySig      = "ewYjBILGshEmTDDMWLeBc9kQfIscSKxmFLnUBU/eXQCb0hrY1jh7U5SH41JmYowuA4p6+YPLcB9z/ay7OvG/Sg=="
	requestBodyDigest      = "SHA-256=uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek="
	requestBodyFalseDigest = "SHA-256=fakeDigest="
	requestBodySig         = "s8MEyer3dSpSsnL0+mQvUYgKm2S4AEX+hsvKmeNI7wgtLFplbCZtt8YOcySZrCyYbOJdPF1NASDHfupSuekecg=="
	requestHost            = "kyber.network"
	requestHostSig         = "+qpk6uAlILo/1YV1ZDK2suU46fbaRi5guOyg4b6aS4nWqLi9u57V6mVwQNh0s6OpfrVZwAYaWHCmQFCgJiZ6yg=="
	algoHmacSha512         = "hmac-sha512"
)

var (
	hmacsha512 = &crypto.HmacSha512{}
	secrets    = Secrets{
		readID: &Secret{
			Key:       "1234",
			Algorithm: hmacsha512,
		},
		writeID: &Secret{
			Key:       "5678",
			Algorithm: hmacsha512,
		},
	}
	requiredHeaders = []string{"(request-target)", "date", "digest"}
	submitHeader    = []string{"(request-target)", "date", "digest"}
	submitHeader2   = []string{"(request-target)", "date", "digest", "host"}
	requestTime     = time.Date(2018, time.October, 22, 0o7, 0o0, 0o7, 0o0, time.UTC)
)

func runTest(secretKeys Secrets, headers []string, v []validator.Validator, req *http.Request) *gin类.Context {
	gin类.X设置运行模式(gin类.X常量_运行模式_测试)
	auth := NewAuthenticator(secretKeys, WithRequiredHeaders(headers), WithValidator(v...))
	c, _ := gin类.CreateTestContext(httptest.NewRecorder())
	c.X请求 = req
	auth.Authenticated()(c)
	return c
}

func generateSignature(keyID KeyID, algorithm string, headers []string, signature string) string {
	return fmt.Sprintf(
		"Signature keyId=\"%s\",algorithm=\"%s\",headers=\"%s\",signature=\"%s\"",
		keyID, algorithm, strings.Join(headers, " "), signature,
	)
}

func TestAuthenticatedHeaderNoSignature(t *testing.T) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	require.NoError(t, err)
	c := runTest(secrets, requiredHeaders, nil, req)
	assert.Equal(t, http.StatusUnauthorized, c.Writer.Status())
	assert.Equal(t, ErrNoSignature, c.X错误s[0])
}

func TestAuthenticatedHeaderInvalidSignature(t *testing.T) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	require.NoError(t, err)
	req.Header.Set(authorizationHeader, "hello")
	c := runTest(secrets, requiredHeaders, nil, req)
	assert.Equal(t, http.StatusUnauthorized, c.Writer.Status())
	assert.Equal(t, ErrInvalidAuthorizationHeader, c.X错误s[0])
}

func TestAuthenticatedHeaderWrongKey(t *testing.T) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	require.NoError(t, err)
	sigHeader := generateSignature(invalidKeyID, algoHmacSha512, submitHeader, requestNilBodySig)
	req.Header.Set(authorizationHeader, sigHeader)
	req.Header.Set("Date", time.Now().UTC().Format(http.TimeFormat))
	c := runTest(secrets, requiredHeaders, nil, req)
	assert.Equal(t, http.StatusUnauthorized, c.Writer.Status())
	assert.Equal(t, ErrInvalidKeyID, c.X错误s[0])
}

func TestAuthenticateDateNotAccept(t *testing.T) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	require.NoError(t, err)
	sigHeader := generateSignature(readID, algoHmacSha512, submitHeader, requestNilBodySig)
	req.Header.Set(authorizationHeader, sigHeader)
	req.Header.Set("Date", time.Date(1990, time.October, 20, 0, 0, 0, 0, time.UTC).Format(http.TimeFormat))
	c := runTest(secrets, requiredHeaders, nil, req)
	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
	assert.Equal(t, validator.ErrDateNotInRange, c.X错误s[0])
}

func TestAuthenticateInvalidRequiredHeader(t *testing.T) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	require.NoError(t, err)
	invalidRequiredHeaders := []string{"date"}
	sigHeader := generateSignature(readID, algoHmacSha512, invalidRequiredHeaders, requestNilBodySig)
	req.Header.Set(authorizationHeader, sigHeader)

	req.Header.Set("Date", time.Now().UTC().Format(http.TimeFormat))

	c := runTest(secrets, requiredHeaders, nil, req)
	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
	assert.Equal(t, ErrHeaderNotEnough, c.X错误s[0])
}

func TestAuthenticateInvalidAlgo(t *testing.T) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	require.NoError(t, err)
	sigHeader := generateSignature(readID, invaldAlgo, submitHeader, requestNilBodySig)
	req.Header.Set(authorizationHeader, sigHeader)
	req.Header.Set("Date", time.Now().UTC().Format(http.TimeFormat))

	c := runTest(secrets, requiredHeaders, nil, req)
	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
	assert.Equal(t, ErrIncorrectAlgorithm, c.X错误s[0])
}

func TestInvalidSign(t *testing.T) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	require.NoError(t, err)
	sigHeader := generateSignature(readID, algoHmacSha512, submitHeader, requestNilBodySig)
	req.Header.Set(authorizationHeader, sigHeader)
	req.Header.Set("Date", time.Now().UTC().Format(http.TimeFormat))

	c := runTest(secrets, requiredHeaders, nil, req)
	assert.Equal(t, http.StatusUnauthorized, c.Writer.Status())
	assert.Equal(t, ErrInvalidSign, c.X错误s[0])
}

// mock interface always return true
type dateAlwaysValid struct{}

func (v *dateAlwaysValid) Validate(r *http.Request) error { return nil }

var mockValidator = []validator.Validator{
	&dateAlwaysValid{},
	validator.NewDigestValidator(),
}

func httpTestGet(c *gin类.Context) {
	c.X输出JSON(http.StatusOK,
		gin类.H{
			"success": true,
		})
}

func httpTestPost(c *gin类.Context) {
	body, err := c.X取流数据()
	if err != nil {
		c.X停止并带状态码(http.StatusInternalServerError)
	}
	c.Render底层方法(http.StatusOK, render.Data{Data: body})
}

func TestHttpInvalidRequest(t *testing.T) {
	gin类.X设置运行模式(gin类.X常量_运行模式_测试)

	r := gin类.X创建默认对象()
	auth := NewAuthenticator(secrets, WithValidator(mockValidator...))
	r.X中间件(auth.Authenticated())
	r.X绑定GET("/", httpTestGet)

	req, err := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	require.NoError(t, err)
	sigHeader := generateSignature(readID, algoHmacSha512, submitHeader, requestBodySig)
	req.Header.Set(authorizationHeader, sigHeader)
	req.Header.Set("Date", requestTime.Format(http.TimeFormat))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.NotEqual(t, http.StatusOK, w.Code)
}

func TestHttpInvalidDigest(t *testing.T) {
	gin类.X设置运行模式(gin类.X常量_运行模式_测试)

	r := gin类.X创建默认对象()
	auth := NewAuthenticator(secrets, WithValidator(mockValidator...))
	r.X中间件(auth.Authenticated())
	r.X绑定POST("/", httpTestPost)

	req, err := http.NewRequestWithContext(context.Background(), "POST", "/", strings.NewReader(sampleBodyContent))
	require.NoError(t, err)
	sigHeader := generateSignature(readID, algoHmacSha512, submitHeader, requestBodySig)
	req.Header.Set(authorizationHeader, sigHeader)
	req.Header.Set("Date", requestTime.Format(http.TimeFormat))
	req.Header.Set("Digest", requestBodyFalseDigest)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHttpValidRequest(t *testing.T) {
	gin类.X设置运行模式(gin类.X常量_运行模式_测试)

	r := gin类.X创建默认对象()
	auth := NewAuthenticator(secrets, WithValidator(mockValidator...))
	r.X中间件(auth.Authenticated())
	r.X绑定GET("/", httpTestGet)

	req, err := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	require.NoError(t, err)
	sigHeader := generateSignature(readID, algoHmacSha512, submitHeader, requestNilBodySig)
	req.Header.Set(authorizationHeader, sigHeader)
	req.Header.Set("Date", requestTime.Format(http.TimeFormat))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHttpValidRequestBody(t *testing.T) {
	gin类.X设置运行模式(gin类.X常量_运行模式_测试)

	r := gin类.X创建默认对象()
	auth := NewAuthenticator(secrets, WithValidator(mockValidator...))
	r.X中间件(auth.Authenticated())
	r.X绑定POST("/", httpTestPost)

	req, err := http.NewRequestWithContext(context.Background(), "POST", "/", strings.NewReader(sampleBodyContent))
	require.NoError(t, err)
	sigHeader := generateSignature(readID, algoHmacSha512, submitHeader, requestBodySig)
	req.Header.Set(authorizationHeader, sigHeader)
	req.Header.Set("Date", requestTime.Format(http.TimeFormat))
	req.Header.Set("Digest", requestBodyDigest)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	body, err := ioutil.ReadAll(w.Result().Body)
	assert.NoError(t, err)
	assert.Equal(t, body, []byte(sampleBodyContent))
}

func TestHttpValidRequestHost(t *testing.T) {
	gin类.X设置运行模式(gin类.X常量_运行模式_测试)

	r := gin类.X创建默认对象()
	auth := NewAuthenticator(secrets, WithValidator(mockValidator...))
	r.X中间件(auth.Authenticated())
	r.X绑定POST("/", httpTestPost)

	requestURL := fmt.Sprintf("http://%s/", requestHost)
	req, err := http.NewRequestWithContext(context.Background(), "POST", requestURL, strings.NewReader(sampleBodyContent))
	assert.NoError(t, err)
	sigHeader := generateSignature(readID, algoHmacSha512, submitHeader2, requestHostSig)
	req.Header.Set(authorizationHeader, sigHeader)
	req.Header.Set("Date", requestTime.Format(http.TimeFormat))
	req.Header.Set("Digest", requestBodyDigest)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	body, err := ioutil.ReadAll(w.Result().Body)
	assert.NoError(t, err)
	assert.Equal(t, body, []byte(sampleBodyContent))
}
