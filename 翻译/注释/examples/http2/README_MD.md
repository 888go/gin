
<原文开始>
How to generate RSA private key and digital certificate

1. Install Openssl

Please visit https://github.com/openssl/openssl to get pkg and install.

2. Generate RSA private key

```sh
$ mkdir testdata
$ openssl genrsa -out ./testdata/server.key 2048
```

3. Generate digital certificate

```sh
$ openssl req -new -x509 -key ./testdata/server.key -out ./testdata/server.pem -days 365
```

<原文结束>

# <翻译开始>
# 如何生成RSA私钥和数字证书

1. 安装OpenSSL

请访问 https://github.com/openssl/openssl 获取安装包并进行安装。

2. 生成RSA私钥

```sh
$ mkdir testdata
$ openssl genrsa -out ./testdata/server.key 2048
```

3. 生成数字证书

```sh
$ openssl req -new -x509 -key ./testdata/server.key -out ./testdata/server.pem -days 365
```

# <翻译结束>

