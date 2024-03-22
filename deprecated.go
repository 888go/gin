// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package gin类

import (
	"log"
	
	"github.com/888go/gin/binding"
)

// BindWith 使用指定的绑定引擎绑定传入的结构体指针。
// 请参阅binding包以了解更多信息。
func (c *Context) X弃用BindWith(obj any, b binding.Binding) error {
	log.Println(`BindWith(\"any, binding.Binding\") error is going to
	be deprecated, please check issue #662 and either use MustBindWith() if you
	want HTTP 400 to be automatically returned if any error occur, or use
	ShouldBindWith() if you need to manage the error.`)
	return c.X取参数到指针并按类型PANI(obj, b)
}
