// 版权所有2018 Gin Core Team
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package binding

type uriBinding struct{}

func (uriBinding) Name() string {
	return "uri"
}

func (uriBinding) BindUri(m map[string][]string, obj any) error {
	if err := mapURI(obj, m); err != nil {
		return err
	}
	return validate(obj)
}
