
<原文开始>
// Config is a struct for specifying configuration options for the secure.
<原文结束>

# <翻译开始>
// Config is a struct for specifying configuration options for the secure.
# <翻译结束>


<原文开始>
	// AllowedHosts is a list of fully qualified domain names that are allowed.
	//Default is empty list, which allows any and all host names.
<原文结束>

# <翻译开始>
	// AllowedHosts is a list of fully qualified domain names that are allowed.
	//Default is empty list, which allows any and all host names.
# <翻译结束>


<原文开始>
	// If SSLRedirect is set to true, then only allow https requests.
	// Default is false.
<原文结束>

# <翻译开始>
	// If SSLRedirect is set to true, then only allow https requests.
	// Default is false.
# <翻译结束>


<原文开始>
	// If SSLTemporaryRedirect is true, the a 302 will be used while redirecting.
	// Default is false (301).
<原文结束>

# <翻译开始>
	// If SSLTemporaryRedirect is true, the a 302 will be used while redirecting.
	// Default is false (301).
# <翻译结束>


<原文开始>
	// SSLHost is the host name that is used to redirect http requests to https.
	// Default is "", which indicates to use the same host.
<原文结束>

# <翻译开始>
	// SSLHost is the host name that is used to redirect http requests to https.
	// Default is "", which indicates to use the same host.
# <翻译结束>


<原文开始>
	// STSSeconds is the max-age of the Strict-Transport-Security header.
	// Default is 0, which would NOT include the header.
<原文结束>

# <翻译开始>
	// STSSeconds is the max-age of the Strict-Transport-Security header.
	// Default is 0, which would NOT include the header.
# <翻译结束>


<原文开始>
	// If STSIncludeSubdomains is set to true, the `includeSubdomains` will
	// be appended to the Strict-Transport-Security header. Default is false.
<原文结束>

# <翻译开始>
	// If STSIncludeSubdomains is set to true, the `includeSubdomains` will
	// be appended to the Strict-Transport-Security header. Default is false.
# <翻译结束>


<原文开始>
	// If FrameDeny is set to true, adds the X-Frame-Options header with
	// the value of `DENY`. Default is false.
<原文结束>

# <翻译开始>
	// If FrameDeny is set to true, adds the X-Frame-Options header with
	// the value of `DENY`. Default is false.
# <翻译结束>


<原文开始>
	// CustomFrameOptionsValue allows the X-Frame-Options header value
	// to be set with a custom value. This overrides the FrameDeny option.
<原文结束>

# <翻译开始>
	// CustomFrameOptionsValue allows the X-Frame-Options header value
	// to be set with a custom value. This overrides the FrameDeny option.
# <翻译结束>


<原文开始>
	// If ContentTypeNosniff is true, adds the X-Content-Type-Options header
	// with the value `nosniff`. Default is false.
<原文结束>

# <翻译开始>
	// If ContentTypeNosniff is true, adds the X-Content-Type-Options header
	// with the value `nosniff`. Default is false.
# <翻译结束>


<原文开始>
	// If BrowserXssFilter is true, adds the X-XSS-Protection header with
	// the value `1; mode=block`. Default is false.
<原文结束>

# <翻译开始>
	// If BrowserXssFilter is true, adds the X-XSS-Protection header with
	// the value `1; mode=block`. Default is false.
# <翻译结束>


<原文开始>
	// ContentSecurityPolicy allows the Content-Security-Policy header value
	// to be set with a custom value. Default is "".
<原文结束>

# <翻译开始>
	// ContentSecurityPolicy allows the Content-Security-Policy header value
	// to be set with a custom value. Default is "".
# <翻译结束>


<原文开始>
	// HTTP header "Referrer-Policy" governs which referrer information, sent in the Referrer header, should be included with requests made.
<原文结束>

# <翻译开始>
	// HTTP header "Referrer-Policy" governs which referrer information, sent in the Referrer header, should be included with requests made.
# <翻译结束>


<原文开始>
	// When true, the whole security policy applied by the middleware is disabled completely.
<原文结束>

# <翻译开始>
	// When true, the whole security policy applied by the middleware is disabled completely.
# <翻译结束>


<原文开始>
	// Handlers for when an error occurs (ie bad host).
<原文结束>

# <翻译开始>
	// Handlers for when an error occurs (ie bad host).
# <翻译结束>


<原文开始>
	// Feature Policy is a new header that allows a site to control which features and APIs can be used in the browser.
<原文结束>

# <翻译开始>
	// Feature Policy is a new header that allows a site to control which features and APIs can be used in the browser.
# <翻译结束>


<原文开始>
	// If DontRedirectIPV4Hostnames is true, requests to hostnames that are IPV4
	// addresses aren't redirected. This is to allow load balancer health checks
	// to succeed.
<原文结束>

# <翻译开始>
	// If DontRedirectIPV4Hostnames is true, requests to hostnames that are IPV4
	// addresses aren't redirected. This is to allow load balancer health checks
	// to succeed.
# <翻译结束>


<原文开始>
	// If the request is insecure, treat it as secure if any of the headers in this dict are set to their corresponding value
	// This is useful when your app is running behind a secure proxy that forwards requests to your app over http (such as on Heroku).
<原文结束>

# <翻译开始>
	// If the request is insecure, treat it as secure if any of the headers in this dict are set to their corresponding value
	// This is useful when your app is running behind a secure proxy that forwards requests to your app over http (such as on Heroku).
# <翻译结束>


<原文开始>
// DefaultConfig returns a Configuration with strict security settings.
// ```
//		SSLRedirect:           true
//		IsDevelopment:         false
//		STSSeconds:            315360000
//		STSIncludeSubdomains:  true
//		FrameDeny:             true
//		ContentTypeNosniff:    true
//		BrowserXssFilter:      true
//		ContentSecurityPolicy: "default-src 'self'"
//		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
// ```
<原文结束>

# <翻译开始>
// DefaultConfig returns a Configuration with strict security settings.
// ```
//		SSLRedirect:           true
//		IsDevelopment:         false
//		STSSeconds:            315360000
//		STSIncludeSubdomains:  true
//		FrameDeny:             true
//		ContentTypeNosniff:    true
//		BrowserXssFilter:      true
//		ContentSecurityPolicy: "default-src 'self'"
//		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
// ```
# <翻译结束>


<原文开始>
// New creates an instance of the secure middleware using the specified configuration.
// router.Use(secure.N)
<原文结束>

# <翻译开始>
// New creates an instance of the secure middleware using the specified configuration.
// router.Use(secure.N)
# <翻译结束>

