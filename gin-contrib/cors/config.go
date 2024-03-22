package cors

import (
	"net/http"
	"strings"
	
	"github.com/888go/gin"
)

type cors struct {
	allowAllOrigins           bool
	allowCredentials          bool
	allowOriginFunc           func(string) bool
	allowOrigins              []string
	normalHeaders             http.Header
	preflightHeaders          http.Header
	wildcardOrigins           [][]string
	optionsResponseStatusCode int
}

var (
	DefaultSchemas = []string{
		"http://",
		"https://",
	}
	ExtensionSchemas = []string{
		"chrome-extension://",
		"safari-extension://",
		"moz-extension://",
		"ms-browser-extension://",
	}
	FileSchemas = []string{
		"file://",
	}
	WebSocketSchemas = []string{
		"ws://",
		"wss://",
	}
)

func newCors(config Config) *cors {
	if err := config.Validate(); err != nil {
		panic(err.Error())
	}

	for _, origin := range config.AllowOrigins {
		if origin == "*" {
			config.AllowAllOrigins = true
		}
	}

	if config.OptionsResponseStatusCode == 0 {
		config.OptionsResponseStatusCode = http.StatusNoContent
	}

	return &cors{
		allowOriginFunc:           config.AllowOriginFunc,
		allowAllOrigins:           config.AllowAllOrigins,
		allowCredentials:          config.AllowCredentials,
		allowOrigins:              normalize(config.AllowOrigins),
		normalHeaders:             generateNormalHeaders(config),
		preflightHeaders:          generatePreflightHeaders(config),
		wildcardOrigins:           config.parseWildcardRules(),
		optionsResponseStatusCode: config.OptionsResponseStatusCode,
	}
}

func (cors *cors) applyCors(c *gin类.Context) {
	origin := c.X请求.Header.Get("Origin")
	if len(origin) == 0 {
		// request is not a CORS request
		return
	}
	host := c.X请求.Host

	if origin == "http://"+host || origin == "https://"+host {
		// request is not a CORS request but have origin header.
		// for example, use fetch api
		return
	}

	if !cors.validateOrigin(origin) {
		c.X停止并带状态码(http.StatusForbidden)
		return
	}

	if c.X请求.Method == "OPTIONS" {
		cors.handlePreflight(c)
		defer c.X停止并带状态码(cors.optionsResponseStatusCode)
	} else {
		cors.handleNormal(c)
	}

	if !cors.allowAllOrigins {
		c.X设置响应协议头值("Access-Control-Allow-Origin", origin)
	}
}

func (cors *cors) validateWildcardOrigin(origin string) bool {
	for _, w := range cors.wildcardOrigins {
		if w[0] == "*" && strings.HasSuffix(origin, w[1]) {
			return true
		}
		if w[1] == "*" && strings.HasPrefix(origin, w[0]) {
			return true
		}
		if strings.HasPrefix(origin, w[0]) && strings.HasSuffix(origin, w[1]) {
			return true
		}
	}

	return false
}

func (cors *cors) validateOrigin(origin string) bool {
	if cors.allowAllOrigins {
		return true
	}
	for _, value := range cors.allowOrigins {
		if value == origin {
			return true
		}
	}
	if len(cors.wildcardOrigins) > 0 && cors.validateWildcardOrigin(origin) {
		return true
	}
	if cors.allowOriginFunc != nil {
		return cors.allowOriginFunc(origin)
	}
	return false
}

func (cors *cors) handlePreflight(c *gin类.Context) {
	header := c.Writer.Header()
	for key, value := range cors.preflightHeaders {
		header[key] = value
	}
}

func (cors *cors) handleNormal(c *gin类.Context) {
	header := c.Writer.Header()
	for key, value := range cors.normalHeaders {
		header[key] = value
	}
}
