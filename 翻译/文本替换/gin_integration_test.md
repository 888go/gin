
<原文开始>
	router := New()

	unixTestSocket := filepath.Join(os.TempDir(), "unix_unit_test")

	defer os.Remove(unixTestSocket)

	go func() {
		router.GET("/example", func(c *Context) { c.String(http.StatusOK, "it worked") })
		assert.NoError(t, router.RunUnix(unixTestSocket))
	}()
	// have to wait for the goroutine to start and run the server
	// otherwise the main thread will complete
	time.Sleep(5 * time.Millisecond)

	c, err := net.Dial("unix", unixTestSocket)
	assert.NoError(t, err)

	fmt.Fprint(c, "GET /example HTTP/1.0\r\n\r\n")
	scanner := bufio.NewScanner(c)
	var response string
	for scanner.Scan() {
		response += scanner.Text()
	}
	assert.Contains(t, response, "HTTP/1.0 200", "should get a 200")
	assert.Contains(t, response, "it worked", "resp body should match")
<原文结束>

# <替换开始>
	//2023-12-09 这在win平台编译不过. 原版gin就是如此.
	//router := New()
	//
	//unixTestSocket := filepath.Join(os.TempDir(), "unix_unit_test")
	//
	//defer os.Remove(unixTestSocket)
	//
	//go func() {
	//	router.GET("/example", func(c *Context) { c.String(http.StatusOK, "it worked") })
	//	assert.NoError(t, router.RunUnix(unixTestSocket))
	//}()
	//// have to wait for the goroutine to start and run the server
	//// otherwise the main thread will complete
	//time.Sleep(5 * time.Millisecond)
	//
	//c, err := net.Dial("unix", unixTestSocket)
	//assert.NoError(t, err)
	//
	//fmt.Fprint(c, "GET /example HTTP/1.0\r\n\r\n")
	//scanner := bufio.NewScanner(c)
	//var response string
	//for scanner.Scan() {
	//	response += scanner.Text()
	//}
	//assert.Contains(t, response, "HTTP/1.0 200", "should get a 200")
	//assert.Contains(t, response, "it worked", "resp body should match")
# <替换结束>


<原文开始>
"path/filepath"
<原文结束>

# <替换开始>
//单元测试"TestUnixSocket"被屏蔽, 引入作废. path/filepath
# <替换结束>

