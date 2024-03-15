// 版权所有2017马努·马丁内斯-阿尔梅达
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

//go:build appengine

package gin

func init() {
	defaultPlatform = PlatformGoogleAppEngine
}
