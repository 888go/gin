
<原文开始>
// Config represents all available options for the middleware.
<原文结束>

# <翻译开始>
// Config represents all available options for the middleware.
# <翻译结束>


<原文开始>
	// AllowOrigins is a list of origins a cross-domain request can be executed from.
	// If the special "*" value is present in the list, all origins will be allowed.
	// Default value is []
<原文结束>

# <翻译开始>
	// AllowOrigins is a list of origins a cross-domain request can be executed from.
	// If the special "*" value is present in the list, all origins will be allowed.
	// Default value is []
# <翻译结束>


<原文开始>
	// AllowOriginFunc is a custom function to validate the origin. It takes the origin
	// as an argument and returns true if allowed or false otherwise. If this option is
	// set, the content of AllowOrigins is ignored.
<原文结束>

# <翻译开始>
	// AllowOriginFunc is a custom function to validate the origin. It takes the origin
	// as an argument and returns true if allowed or false otherwise. If this option is
	// set, the content of AllowOrigins is ignored.
# <翻译结束>


<原文开始>
	// AllowMethods is a list of methods the client is allowed to use with
	// cross-domain requests. Default value is simple methods (GET, POST, PUT, PATCH, DELETE, HEAD, and OPTIONS)
<原文结束>

# <翻译开始>
	// AllowMethods is a list of methods the client is allowed to use with
	// cross-domain requests. Default value is simple methods (GET, POST, PUT, PATCH, DELETE, HEAD, and OPTIONS)
# <翻译结束>


<原文开始>
	// AllowPrivateNetwork indicates whether the response should include allow private network header
<原文结束>

# <翻译开始>
	// AllowPrivateNetwork indicates whether the response should include allow private network header
# <翻译结束>


<原文开始>
	// AllowHeaders is list of non simple headers the client is allowed to use with
	// cross-domain requests.
<原文结束>

# <翻译开始>
	// AllowHeaders is list of non simple headers the client is allowed to use with
	// cross-domain requests.
# <翻译结束>


<原文开始>
	// AllowCredentials indicates whether the request can include user credentials like
	// cookies, HTTP authentication or client side SSL certificates.
<原文结束>

# <翻译开始>
	// AllowCredentials indicates whether the request can include user credentials like
	// cookies, HTTP authentication or client side SSL certificates.
# <翻译结束>


<原文开始>
	// ExposeHeaders indicates which headers are safe to expose to the API of a CORS
	// API specification
<原文结束>

# <翻译开始>
	// ExposeHeaders indicates which headers are safe to expose to the API of a CORS
	// API specification
# <翻译结束>


<原文开始>
	// MaxAge indicates how long (with second-precision) the results of a preflight request
	// can be cached
<原文结束>

# <翻译开始>
	// MaxAge indicates how long (with second-precision) the results of a preflight request
	// can be cached
# <翻译结束>


<原文开始>
	// Allows to add origins like http://some-domain/*, https://api.* or http://some.*.subdomain.com
<原文结束>

# <翻译开始>
	// Allows to add origins like http://some-domain/*, https://api.* or http://some.*.subdomain.com
# <翻译结束>


<原文开始>
	// Allows usage of popular browser extensions schemas
<原文结束>

# <翻译开始>
	// Allows usage of popular browser extensions schemas
# <翻译结束>


<原文开始>
	// Allows usage of WebSocket protocol
<原文结束>

# <翻译开始>
	// Allows usage of WebSocket protocol
# <翻译结束>


<原文开始>
	// Allows usage of file:// schema (dangerous!) use it only when you 100% sure it's needed
<原文结束>

# <翻译开始>
	// Allows usage of file:// schema (dangerous!) use it only when you 100% sure it's needed
# <翻译结束>


<原文开始>
	// Allows to pass custom OPTIONS response status code for old browsers / clients
<原文结束>

# <翻译开始>
	// Allows to pass custom OPTIONS response status code for old browsers / clients
# <翻译结束>


<原文开始>
// AddAllowMethods is allowed to add custom methods
<原文结束>

# <翻译开始>
// AddAllowMethods is allowed to add custom methods
# <翻译结束>


<原文开始>
// AddAllowHeaders is allowed to add custom headers
<原文结束>

# <翻译开始>
// AddAllowHeaders is allowed to add custom headers
# <翻译结束>


<原文开始>
// AddExposeHeaders is allowed to add custom expose headers
<原文结束>

# <翻译开始>
// AddExposeHeaders is allowed to add custom expose headers
# <翻译结束>


<原文开始>
// Validate is check configuration of user defined.
<原文结束>

# <翻译开始>
// Validate is check configuration of user defined.
# <翻译结束>


<原文开始>
// DefaultConfig returns a generic default configuration mapped to localhost.
<原文结束>

# <翻译开始>
// DefaultConfig returns a generic default configuration mapped to localhost.
# <翻译结束>


<原文开始>
// Default returns the location middleware with default configuration.
<原文结束>

# <翻译开始>
// Default returns the location middleware with default configuration.
# <翻译结束>


<原文开始>
// New returns the location middleware with user-defined custom configuration.
<原文结束>

# <翻译开始>
// New returns the location middleware with user-defined custom configuration.
# <翻译结束>

