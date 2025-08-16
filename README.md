# PasswordGen
PassGen 是一个基于 Go 的轻量级密码生成器

你可以访问示例网站：[PassGen](https://passgen.parrotstudio.xyz/)

## 功能

- 生成强密码
- 支持自定义字符集

## 编译

```bash
./build.ps1
````

## 运行

默认端口8080

```bash
./PassGen-go_v1.0_linux_amd64
```

自定义端口

```bash
./PassGen-go_v1.0_linux_amd64 -port 1234
```

## 访问

打开浏览器访问 `http://localhost:端口`。

## 备注

* 建议通过反向代理（如 Caddy）进行 TLS 加密
