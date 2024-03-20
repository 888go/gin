
<原文开始>
// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// TODO
// func (w *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
// func (w *responseWriter) CloseNotify() <-chan bool {
// func (w *responseWriter) Flush() {
<原文结束>

# <翻译开始>
// TODO：待办事项（需要实现或处理）
// func (w *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
//   // 函数功能：Hijack方法，获取原始的网络连接、读写缓冲器和可能发生的错误
//   
// func (w *responseWriter) CloseNotify() <-chan bool {
//   // 函数功能：CloseNotify方法，返回一个只读通道，当客户端连接关闭时，该通道会接收到一个布尔值true通知
//   
// func (w *responseWriter) Flush() {
//   // 函数功能：Flush方法，用于立即将响应数据刷新到客户端，通常用于在HTTP流式传输中强制发送已缓存的数据
# <翻译结束>


<原文开始>
// status must be 200 although we tried to change it
<原文结束>

# <翻译开始>
// 状态码必须为200，尽管我们尝试改变它
# <翻译结束>


<原文开始>
// mockPusherResponseWriter is an http.ResponseWriter that implements http.Pusher.
<原文结束>

# <翻译开始>
// mockPusherResponseWriter 是一个实现了 http.Pusher 接口的 http.ResponseWriter。
# <翻译结束>


<原文开始>
// nonPusherResponseWriter is an http.ResponseWriter that does not implement http.Pusher.
<原文结束>

# <翻译开始>
// nonPusherResponseWriter 是一个 http.ResponseWriter，但它并不实现 http.Pusher 接口。
# <翻译结束>

