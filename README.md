# top-go

京东联盟、宙斯开放平台 API 请求 SDK（Go 语言版）。

由于京东联盟仅提供 Java 语言的 SDK 包，宙斯平台的 PHP SDK 也比较老旧，遂手撸此项目。

此 SDK 提供了对 API 的请求封装、摘要签名等功能，使用 SDK 可以轻松完成 API 的调用。

## 接入文档

* [联盟 API 接入文档](https://union.jd.com/helpcenter/13246-13247-46301)
* [宙斯 API 接入文档](https://union.jd.com/helpcenter/13246-13312-57749)
* [宙斯 API 调用方法详解](https://open.jd.com/home/home#/doc/common?listId=890)

## 结构说明

```
├── Gopkg.lock
├── Gopkg.toml
├── LICENSE
├── Makefile
├── README.md
├── config
│   └── config.go # 加载配置
├── config_example.toml # 配置示例
├── jop
│   └── jop.go # 业务逻辑
├── jop-sdk
│   ├── client-new.go # 同 client.go，把系统参数封装为结构体
│   ├── client.go # 摘要签名、基础请求
│   ├── client_test.go
│   ├── jop-api.go # 接口业务参数封装
│   ├── jop-api_test.go
│   └── vars.go # 包内公共变量
├── main.go
└── vendor
    ├── github.com
    └── golang.org
```
