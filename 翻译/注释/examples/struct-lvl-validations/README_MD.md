
<原文开始>
Struct level validations

Validations can also be registered at the `struct` level when field level validations
don't make much sense. This can also be used to solve cross-field validation elegantly.
Additionally, it can be combined with tag validations. Struct Level validations run after
the structs tag validations.


<原文结束>

# <翻译开始>
# 结构级别验证

当字段级别验证不太合理时，也可以在 `struct` 级别注册验证。这也可以用于优雅地解决跨字段验证问题。此外，它可以与标签验证相结合。结构级别验证在结构的标签验证之后运行。

# <翻译结束>


<原文开始>
Example requests

```shell

<原文结束>

# <翻译开始>
# 示例请求

```shell

# <翻译结束>


<原文开始>
Validation errors are generated for struct tags as well as at the struct level
$ curl -s -X POST http://localhost:8085/user \
  -H 'content-type: application/json' \
  -d '{}' | jq
{
  "error": "Key: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag\nKey: 'User.FirstName' Error:Field validation for 'FirstName' failed on the 'fnameorlname' tag\nKey: 'User.LastName' Error:Field validation for 'LastName' failed on the 'fnameorlname' tag",
  "message": "User validation failed!"
}


<原文结束>

# <翻译开始>
# 对于结构体标签以及结构体级别，都会生成验证错误。
通过以下命令执行：

```bash
curl -s -X POST http://localhost:8085/user \
  -H 'content-type: application/json' \
  -d '{}' | jq
```

返回结果如下（经过`jq`格式化）：

```json
{
  "error": "Key: 'User.Email' 错误: 对'Email'字段的验证在'required'标签上失败\nKey: 'User.FirstName' 错误: 对'FirstName'字段的验证在'fnameorlname'标签上失败\nKey: 'User.LastName' 错误: 对'LastName'字段的验证在'fnameorlname'标签上失败",
  "message": "用户验证失败！"
}
```

这意味着在尝试创建或更新用户时，由于"Email"字段缺少必填项（required），且"FirstName"和"LastName"字段未能通过名为'fnameorlname'的自定义校验规则，所以整个用户对象验证未通过。

# <翻译结束>


<原文开始>
Validation fails at the struct level because neither first name nor last name are present
$ curl -s -X POST http://localhost:8085/user \
    -H 'content-type: application/json' \
    -d '{"email": "george@vandaley.com"}' | jq
{
  "error": "Key: 'User.FirstName' Error:Field validation for 'FirstName' failed on the 'fnameorlname' tag\nKey: 'User.LastName' Error:Field validation for 'LastName' failed on the 'fnameorlname' tag",
  "message": "User validation failed!"
}


<原文结束>

# <翻译开始>
# 验证在结构体级别失败，因为既没有提供名也没有提供姓。
$ curl -s -X POST http://localhost:8085/user \
    -H 'content-type: application/json' \
    -d '{"email": "george@vandaley.com"}' | jq
{
  "error": "错误：'User.FirstName'，'FirstName'字段的验证在'fnameorlname'标签上失败\n错误：'User.LastName'，'LastName'字段的验证在'fnameorlname'标签上失败",
  "message": "用户验证失败！"
}

# <翻译结束>


<原文开始>
No validation errors when either first name or last name is present
$ curl -X POST http://localhost:8085/user \
    -H 'content-type: application/json' \
    -d '{"fname": "George", "email": "george@vandaley.com"}'
{"message":"User validation successful."}

$ curl -X POST http://localhost:8085/user \
    -H 'content-type: application/json' \
    -d '{"lname": "Contanza", "email": "george@vandaley.com"}'
{"message":"User validation successful."}

$ curl -X POST http://localhost:8085/user \
    -H 'content-type: application/json' \
    -d '{"fname": "George", "lname": "Costanza", "email": "george@vandaley.com"}'
{"message":"User validation successful."}
```

#
<原文结束>

# <翻译开始>
# 当存在名或姓时，没有验证错误

```bash
# 当仅提供名和邮箱时
$ curl -X POST http://localhost:8085/user \
    -H 'content-type: application/json' \
    -d '{"fname": "George", "email": "george@vandaley.com"}'
{"message":"用户验证成功。"}

# 当仅提供姓和邮箱时
$ curl -X POST http://localhost:8085/user \
    -H 'content-type: application/json' \
    -d '{"lname": "Contanza", "email": "george@vandaley.com"}'
{"message":"用户验证成功。"}

# 当同时提供名、姓和邮箱时
$ curl -X POST http://localhost:8085/user \
    -H 'content-type: application/json' \
    -d '{"fname": "George", "lname": "Costanza", "email": "george@vandaley.com"}'
{"message":"用户验证成功。"}
```

```

# <翻译结束>


<原文开始>
Useful links

- [Validator docs](https://pkg.go.dev/github.com/go-playground/validator/v10#Validate.RegisterStructValidation)
- [Struct level example](https://github.com/go-playground/validator/blob/master/_examples/struct-level/main.go)
- [Validator release notes](https://github.com/go-playground/validator/releases/tag/v10.7.0)

<原文结束>

# <翻译开始>
# 有用链接

- [验证器文档](https://pkg.go.dev/github.com/go-playground/validator/v10#Validate.RegisterStructValidation)
- [结构级别示例](https://github.com/go-playground/validator/blob/master/_examples/struct-level/main.go)
- [验证器发行说明](https://github.com/go-playground/validator/releases/tag/v10.7.0)

# <翻译结束>

