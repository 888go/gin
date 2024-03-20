## # 如何生成RSA私钥和数字证书

1. 安装OpenSSL

请访问https://github.com/openssl/openssl获取pkg并进行安装。

2. 生成RSA私钥

```sh
$ mkdir testdata
$ openssl genrsa -out ./testdata/server.key 2048
```

3. 生成数字证书

```sh
$ openssl req -new -x509 -key ./testdata/server.key -out ./testdata/server.pem -days 365
```

翻译：

1. 安装 OpenSSL

请前往 https://github.com/openssl/openssl 下载并安装 OpenSSL。

2. 创建RSA私钥

```shell
$ mkdir testdata
$ openssl 生成rsa密钥 -out ./testdata/server.key 2048
```

3. 创建数字证书

```shell
$ openssl 申请 -new -x509 -使用密钥 ./testdata/server.key -输出 ./testdata/server.pem -有效期 365天
```
