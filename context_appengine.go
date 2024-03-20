// 版权所有 ? 2017 Manu Martinez-Almeida。保留所有权利。
// 本源代码的使用受 MIT 风格许可证协议约束，
// 该协议可在 LICENSE 文件中查阅。

//go:build appengine

package gin

func init() {
	defaultPlatform = PlatformGoogleAppEngine
}
