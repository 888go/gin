// Gin Core团队版权所有版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

//go:build !nomsgpack

package binding

import (
	"bytes"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/ugorji/go/codec"
)


// ff:
// t:
func TestBindingMsgPack(t *testing.T) {
	test := FooStruct{
		Foo: "bar",
	}

	h := new(codec.MsgpackHandle)
	assert.NotNil(t, h)
	buf := bytes.NewBuffer([]byte{})
	assert.NotNil(t, buf)
	err := codec.NewEncoder(buf, h).Encode(test)
	assert.NoError(t, err)

	data := buf.Bytes()

	testMsgPackBodyBinding(t,
		MsgPack, "msgpack",
		"/", "/",
		string(data), string(data[1:]))
}

func testMsgPackBodyBinding(t *testing.T, b Binding, name, path, badPath, body, badBody string) {
	assert.Equal(t, name, b.Name())

	obj := FooStruct{}
	req := requestWithBody("POST", path, body)
	req.Header.Add("Content-Type", MIMEMSGPACK)
	err := b.Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, "bar", obj.Foo)

	obj = FooStruct{}
	req = requestWithBody("POST", badPath, badBody)
	req.Header.Add("Content-Type", MIMEMSGPACK)
	err = MsgPack.Bind(req, &obj)
	assert.Error(t, err)
}


// ff:
// t:
func TestBindingDefaultMsgPack(t *testing.T) {
	assert.Equal(t, MsgPack, Default("POST", MIMEMSGPACK))
	assert.Equal(t, MsgPack, Default("PUT", MIMEMSGPACK2))
}
