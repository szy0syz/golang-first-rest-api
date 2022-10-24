# golang-first-rest-api

## Architecture

![image](https://user-images.githubusercontent.com/10555820/197550582-5798d9e5-7063-4a11-8a7f-c5ed55f49238.png)

## bash

```bash
go get github.com/gorilla/mux
go get github.com/sirupsen/logrus
go get github.com/jmoiron/sqlx
go get github.com/lib/pq
go get github.com/asaskevich/govalidator
go get github.com/joho/godotenv
go get github.com/kelseyhightower/envconfig
```

## package

- **mux**
  - 用于Web框架处理路由的库，超轻量
  - 实现了`http.Handler`，天生与原生模块`net/http`结合使用
  - 可以根据请求的主机名、路径、路径前缀、协议、HTTP 首部、查询字符串和 HTTP 方法匹配处理器，自定义匹配逻辑
- `logrus`
  - 一个结构化的日志包
  - 感觉就是把日志定好规则，让我们好检索吧
- `sqlx`
  - 看包的名称就知道一定很性能
  - 看了包的文档还真是sql语句直接上
- `govalidator`
  - 很好，看名字就知道验证字符串、结构体和集合的
  - 毕竟Api服务，肯定需要验证传递的参数
- `envconfig`
  - 和我们node.js一样，管理配置环境变量的

## Config

> 先搞定配置问题

```bash
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=postgres
DATABASE_PASSWORD=password
DATABASE_NAME=rest
SERVER_PORT=80
```

```go
func NewParsedConfig() (Config, error) {
  // 载入环境变量，是否是还应该考虑开发环境和生产环境 🤔🤔
	_ = godotenv.Load(".env")
  // 初始化个空的Config结构体，对了，golang是个强类型，怕不能搞个空对象
	cnf := Config{}
  // 向进程注入环境变量
	err := envconfig.Process("", &cnf)
  // 得熟悉这样golang风格的返回
	return cnf, err
}
```
