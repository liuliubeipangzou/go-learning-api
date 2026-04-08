# Go Learning API

这是一个给 Java 后端开发者练手的 Go 入门项目。

目标：

- 使用 Go 标准库搭一个最小可运行后端
- 理解 Go 常见分层：`handler -> service -> repository`
- 熟悉结构体、接口、错误处理、JSON 编解码、HTTP 服务
- 对比 Java/Spring Boot 的常见概念
- 让你可以从 Java 后端开发视角，快速建立 Go Web 项目的工程感

## 项目结构

```text
go-learning-api
├── cmd/server/main.go            // 启动入口，类似 Java 的 main 方法
├── internal/app/app.go           // 应用装配
├── internal/config/config.go     // 配置读取
├── internal/handler/user_handler.go
├── internal/middleware           // 日志与 panic 恢复中间件
├── internal/model/user.go
├── internal/repository/user_repository.go
├── internal/response/json.go
├── internal/router/router.go
├── internal/service/user_service.go
├── Dockerfile
├── Makefile
└── .env.example
```

## 你会学到什么

### 1. Go 和 Java 的常见对应关系

- `struct` 类似 Java 的类，但通常只放数据，不强调继承
- `interface` 更轻量，通常由“使用方”定义
- `error` 是显式返回值，不靠异常控制流程
- `net/http` 类似一个更底层的 Web 框架基础设施
- `internal` 目录表示包私有边界，外部模块不能直接导入

### 2. 项目分层

- `handler`：接收 HTTP 请求，做参数解析和响应
- `service`：放业务逻辑
- `repository`：封装数据访问，这里先用内存实现
- `model`：领域对象

## 已实现接口

### 健康检查

```bash
GET /health
```

返回：

```json
{
  "status": "ok"
}
```

### 查询用户列表

```bash
GET /api/v1/users
```

### 查询单个用户

```bash
GET /api/v1/users/{id}
```

### 创建用户

```bash
POST /api/v1/users
Content-Type: application/json
```

请求体：

```json
{
  "name": "Alice",
  "email": "alice@example.com"
}
```

### 更新用户

```bash
PUT /api/v1/users/{id}
Content-Type: application/json
```

### 删除用户

```bash
DELETE /api/v1/users/{id}
```

## 本地安装 Go

建议安装 Go `1.22+`，这个项目当前 `go.mod` 使用的是 `1.22.0`。

### macOS + Homebrew

```bash
brew install go
```

安装完成后检查：

```bash
go version
go env GOPATH GOROOT
```

如果你是第一次接触 Go，建议把下面这行加到 `~/.zshrc`：

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

然后执行：

```bash
source ~/.zshrc
```

## 本地运行

```bash
cd go-learning-api
cp .env.example .env
go run ./cmd/server
```

程序启动时会自动读取项目根目录下的 `.env`。

默认监听：

```text
http://localhost:8080
```

也可以自定义端口：

```bash
APP_PORT=9090 go run ./cmd/server
```

也可以用 `Makefile`：

```bash
make run
make test
make build
make fmt
```

## Docker 运行

如果你暂时不想先装 Go，也可以先用 Docker 跑起来：

```bash
docker build -t go-learning-api .
docker run --rm -p 8080:8080 go-learning-api
```

## 建议学习顺序

1. 先看 `cmd/server/main.go`，理解程序如何启动
2. 再看 `internal/router/router.go`，理解路由如何注册
3. 然后从 `handler -> service -> repository` 顺着调用链读
4. 最后自己扩展一个接口，比如“删除用户”或“更新用户”

现在这个仓库已经自带了“更新用户”和“删除用户”，你可以接着练下面这些更像真实后端需求的内容。

## 推荐练习

1. 给 `User` 增加 `Age` 字段，并补齐校验与测试
2. 把当前内存仓储替换成 MySQL
3. 为用户列表增加分页参数
4. 引入统一错误码和业务错误类型
5. 增加请求日志 trace id
6. 把配置改造成 `.env + envconfig` 或 `viper`

## 测试

安装 Go 后可以直接执行：

```bash
go test ./...
```

当前测试覆盖了：

- service 层的创建、更新、删除、邮箱校验
- handler 层的创建、更新不存在用户、删除用户

## 后续可升级方向

- 引入 `chi` 或 `gin` 做路由
- 引入 `sqlx` 或 `gorm` 操作数据库
- 增加配置文件、依赖注入、分环境配置
- 增加 JWT 登录、权限控制、数据库迁移

## Java 开发者迁移建议

- 先把 `struct` 理解为“轻量 DTO + 领域对象”
- 先接受 Go 的错误返回值风格，不要急着套异常思维
- 不要一开始就找 Spring Boot 对标框架，先把标准库写顺
- 重点体会接口、组合、包边界，而不是继承
- 如果你会 Java，这个项目最值得你练的是“少框架下如何把代码组织清楚”
