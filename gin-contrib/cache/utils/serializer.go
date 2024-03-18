package utils

import (
	"bytes"
	"encoding/gob"
	"reflect"
	"strconv"
)

// Serialize 将传入的值转换为 []byte 表示形式并返回

// ff:
// value:

// ff:
// value:

// ff:
// value:

// ff:
// value:

// ff:
// value:

// ff:
// value:
func Serialize(value interface{}) ([]byte, error) {
	if bytes, ok := value.([]byte); ok {
		return bytes, nil
	}

	switch v := reflect.ValueOf(value); v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []byte(strconv.FormatInt(v.Int(), 10)), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []byte(strconv.FormatUint(v.Uint(), 10)), nil
	}

	var b bytes.Buffer
	encoder := gob.NewEncoder(&b)
	if err := encoder.Encode(value); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// Deserialize 将传入的 []byte 反序列化为传入的 ptr interface{}

// ff:
// err:
// ptr:
// byt:

// ff:
// err:
// ptr:
// byt:

// ff:
// err:
// ptr:
// byt:

// ff:
// err:
// ptr:
// byt:

// ff:
// err:
// ptr:
// byt:

// ff:
// err:
// ptr:
// byt:
func Deserialize(byt []byte, ptr interface{}) (err error) {
	if bytes, ok := ptr.(*[]byte); ok {
		*bytes = byt
		return nil
	}

	if v := reflect.ValueOf(ptr); v.Kind() == reflect.Ptr {
		switch p := v.Elem(); p.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			var i int64
			i, err = strconv.ParseInt(string(byt), 10, 64)
			if err != nil {
				return err
			}

			p.SetInt(i)
			return nil

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			var i uint64
			i, err = strconv.ParseUint(string(byt), 10, 64)
			if err != nil {
				return err
			}

			p.SetUint(i)
			return nil
		}
	}

	b := bytes.NewBuffer(byt)
	decoder := gob.NewDecoder(b)
	if err = decoder.Decode(ptr); err != nil {
		return err
	}
	return nil
}
