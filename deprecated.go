// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package gin
import (
	"log"
	
	"e.coding.net/gogit/go/gin/binding"
	)
// BindWith使用指定的绑定引擎绑定传递的结构指针
// 参见绑定包
func (c *Context) BindWith(obj any, b binding.Binding) error {
	log.Println(`BindWith(\"any, binding.Binding\") error is going to
	be deprecated, please check issue #662 and either use MustBindWith() if you
	want HTTP 400 to be automatically returned if any error occur, or use
	ShouldBindWith() if you need to manage the error.`)
	return c.MustBindWith(obj, b)
}
