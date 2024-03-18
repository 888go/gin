package logger

import (
	"io"
	"regexp"
	
	"github.com/888go/gin"
	"github.com/rs/zerolog"
)

// Option 指定了仪器化配置选项。
type Option interface {
	apply(*config)
}

var _ Option = (*optionFunc)(nil)

type optionFunc func(*config)

func (o optionFunc) apply(c *config) {
	o(c)
}

// WithLogger 设置自定义日志器函数

// ff:
// fn:
// zerolog.Logger:
// *gin.Context:

// ff:
// fn:
// zerolog.Logger:
// *gin.Context:

// ff:
// fn:
// zerolog.Logger:
// *gin.Context:

// ff:
// fn:
// zerolog.Logger:
// *gin.Context:

// ff:
// fn:
// zerolog.Logger:
// *gin.Context:

// ff:
// fn:
// zerolog.Logger:
// *gin.Context:

// ff:
// fn:
// zerolog.Logger:
// *gin.Context:
func WithLogger(fn func(*gin.Context, zerolog.Logger) zerolog.Logger) Option {
	return optionFunc(func(c *config) {
		c.logger = fn
	})
}

// WithSkipPathRegexps 通过正则表达式模式添加多个需要跳过的URL路径

// ff:
// regs:

// ff:
// regs:

// ff:
// regs:

// ff:
// regs:

// ff:
// regs:

// ff:
// regs:

// ff:
// regs:
func WithSkipPathRegexps(regs ...*regexp.Regexp) Option {
	return optionFunc(func(c *config) {
		if len(regs) == 0 {
			return
		}

		c.skipPathRegexps = append(c.skipPathRegexps, regs...)
	})
}

// WithUTC 返回一个时间t，其位置设置为UTC。

// ff:
// s:

// ff:
// s:

// ff:
// s:

// ff:
// s:

// ff:
// s:

// ff:
// s:

// ff:
// s:
func WithUTC(s bool) Option {
	return optionFunc(func(c *config) {
		c.utc = s
	})
}

// WithSkipPath 根据特定模式跳过URL路径

// ff:
// s:

// ff:
// s:

// ff:
// s:

// ff:
// s:

// ff:
// s:

// ff:
// s:

// ff:
// s:
func WithSkipPath(s []string) Option {
	return optionFunc(func(c *config) {
		c.skipPath = s
	})
}

// WithWriter 更改默认输出 writer。
// 默认为 gin.DefaultWriter

// ff:
// s:

// ff:
// s:

// ff:
// s:

// ff:
// s:

// ff:
// s:

// ff:
// s:

// ff:
// s:
func WithWriter(s io.Writer) Option {
	return optionFunc(func(c *config) {
		c.output = s
	})
}


// ff:
// lvl:

// ff:
// lvl:

// ff:
// lvl:

// ff:
// lvl:

// ff:
// lvl:

// ff:
// lvl:

// ff:
// lvl:
func WithDefaultLevel(lvl zerolog.Level) Option {
	return optionFunc(func(c *config) {
		c.defaultLevel = lvl
	})
}


// ff:
// lvl:

// ff:
// lvl:

// ff:
// lvl:

// ff:
// lvl:

// ff:
// lvl:

// ff:
// lvl:

// ff:
// lvl:
func WithClientErrorLevel(lvl zerolog.Level) Option {
	return optionFunc(func(c *config) {
		c.clientErrorLevel = lvl
	})
}


// ff:
// lvl:

// ff:
// lvl:

// ff:
// lvl:

// ff:
// lvl:

// ff:
// lvl:

// ff:
// lvl:

// ff:
// lvl:
func WithServerErrorLevel(lvl zerolog.Level) Option {
	return optionFunc(func(c *config) {
		c.serverErrorLevel = lvl
	})
}
