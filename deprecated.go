// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin类

import (
	"log"
	
	"github.com/888go/gin/binding"
)

// BindWith binds the passed struct pointer using the specified binding engine.
// See the binding package.
func (c *Context) X弃用BindWith(obj any, b binding.Binding) error {
	log.Println(`BindWith(\"any, binding.Binding\") error is going to
	be deprecated, please check issue #662 and either use MustBindWith() if you
	want HTTP 400 to be automatically returned if any error occur, or use
	ShouldBindWith() if you need to manage the error.`)
	return c.X取参数到指针并按类型PANI(obj, b)
}
