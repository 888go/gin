
<原文开始>
// Package ginzap provides log handling using zap package.
// Code structure based on ginrus package.
<原文结束>

# <翻译开始>
// Package ginzap provides log handling using zap package.
// Code structure based on ginrus package.
# <翻译结束>


<原文开始>
// ZapLogger is the minimal logger interface compatible with zap.Logger
<原文结束>

# <翻译开始>
// ZapLogger is the minimal logger interface compatible with zap.Logger
# <翻译结束>


<原文开始>
// Config is config setting for Ginzap
<原文结束>

# <翻译开始>
// Config is config setting for Ginzap
# <翻译结束>


<原文开始>
// Ginzap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
//
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
//
// It receives:
//  1. A time package format string (e.g. time.RFC3339).
//  2. A boolean stating whether to use UTC time zone or local.
<原文结束>

# <翻译开始>
// Ginzap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
//
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
//
// It receives:
//  1. A time package format string (e.g. time.RFC3339).
//  2. A boolean stating whether to use UTC time zone or local.
# <翻译结束>


<原文开始>
// GinzapWithConfig returns a gin.HandlerFunc using configs
<原文结束>

# <翻译开始>
// GinzapWithConfig returns a gin.HandlerFunc using configs
# <翻译结束>


<原文开始>
		// some evil middlewares modify this values
<原文结束>

# <翻译开始>
		// some evil middlewares modify this values
# <翻译结束>


<原文开始>
				// Append error field if this is an erroneous request.
<原文结束>

# <翻译开始>
				// Append error field if this is an erroneous request.
# <翻译结束>


<原文开始>
// RecoveryWithZap returns a gin.HandlerFunc (middleware)
// that recovers from any panics and logs requests using uber-go/zap.
// All errors are logged using zap.Error().
// stack means whether output the stack info.
// The stack info is easy to find where the error occurs but the stack info is too large.
<原文结束>

# <翻译开始>
// RecoveryWithZap returns a gin.HandlerFunc (middleware)
// that recovers from any panics and logs requests using uber-go/zap.
// All errors are logged using zap.Error().
// stack means whether output the stack info.
// The stack info is easy to find where the error occurs but the stack info is too large.
# <翻译结束>


<原文开始>
// CustomRecoveryWithZap returns a gin.HandlerFunc (middleware) with a custom recovery handler
// that recovers from any panics and logs requests using uber-go/zap.
// All errors are logged using zap.Error().
// stack means whether output the stack info.
// The stack info is easy to find where the error occurs but the stack info is too large.
<原文结束>

# <翻译开始>
// CustomRecoveryWithZap returns a gin.HandlerFunc (middleware) with a custom recovery handler
// that recovers from any panics and logs requests using uber-go/zap.
// All errors are logged using zap.Error().
// stack means whether output the stack info.
// The stack info is easy to find where the error occurs but the stack info is too large.
# <翻译结束>


<原文开始>
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
<原文结束>

# <翻译开始>
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
# <翻译结束>


<原文开始>
					// If the connection is dead, we can't write a status to it.
<原文结束>

# <翻译开始>
					// If the connection is dead, we can't write a status to it.
# <翻译结束>

