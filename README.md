# A Cup of Coffee ☕

基于 [go-zero](https://go-zero.dev/) 框架的后端项目工程化模板。

## 📁 项目结构

```
acupofcoffee/
├── api/                        # HTTP API 服务
│   ├── etc/                    # 配置文件
│   │   └── config.yaml
│   ├── internal/               # 内部代码
│   │   ├── config/             # 配置结构
│   │   ├── handler/            # HTTP 处理器
│   │   ├── logic/              # 业务逻辑
│   │   ├── middleware/         # 中间件
│   │   ├── svc/                # 服务上下文
│   │   └── types/              # 类型定义
│   └── main.go                 # 入口文件
├── common/                     # 公共模块
│   ├── errorx/                 # 错误处理
│   ├── response/               # 统一响应
│   └── utils/                  # 工具函数
├── model/                      # 数据模型
├── deploy/                     # 部署配置
│   ├── docker/                 # Docker 配置
│   └── k8s/                    # Kubernetes 配置
├── go.mod
├── Makefile
└── README.md
```

## 🚀 快速开始

### 环境要求

- Go 1.21+
- MySQL 8.0+
- Redis 7.0+
- Docker & Docker Compose (可选)

### 本地开发

1. **克隆项目**
```bash
git clone <repository-url>
cd acupofcoffee
```

2. **安装依赖**
```bash
make deps
```

3. **配置数据库**

修改 `api/etc/config.yaml` 中的数据库配置：
```yaml
MySQL:
  DataSource: root:password@tcp(localhost:3306)/acupofcoffee?charset=utf8mb4&parseTime=True&loc=Local
```

4. **启动服务**
```bash
make run
```

服务将在 `http://localhost:8080` 启动。

### Docker 部署

使用 Docker Compose 一键启动所有服务：

```bash
make docker-up
```

停止服务：
```bash
make docker-down
```

## 📖 API 文档

### 健康检查
```
GET /api/v1/health
```

### 用户认证

**注册**
```
POST /api/v1/auth/register
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123",
  "email": "test@example.com",
  "nickname": "Test User"
}
```

**登录**
```
POST /api/v1/auth/login
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123"
}
```

响应：
```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "accessToken": "eyJhbGciOiJIUzI1NiIs...",
    "accessExpire": 1700000000,
    "refreshAfter": 1699956800
  }
}
```

### 用户信息

**获取用户信息** (需要认证)
```
GET /api/v1/user/info
Authorization: Bearer <token>
```

**更新用户信息** (需要认证)
```
PUT /api/v1/user/info
Authorization: Bearer <token>
Content-Type: application/json

{
  "nickname": "New Nickname",
  "avatar": "https://example.com/avatar.jpg"
}
```

## 🛠 常用命令

| 命令 | 描述 |
|------|------|
| `make deps` | 安装依赖 |
| `make run` | 开发模式运行 |
| `make build` | 构建可执行文件 |
| `make test` | 运行测试 |
| `make fmt` | 格式化代码 |
| `make lint` | 代码检查 |
| `make docker` | 构建 Docker 镜像 |
| `make docker-up` | Docker Compose 启动 |
| `make docker-down` | Docker Compose 停止 |
| `make help` | 查看所有命令 |

## 📦 技术栈

- **框架**: [go-zero](https://go-zero.dev/) - 高性能微服务框架
- **ORM**: [GORM](https://gorm.io/) - Go 语言 ORM 库
- **数据库**: MySQL 8.0
- **缓存**: Redis 7.0
- **认证**: JWT (JSON Web Token)
- **日志**: go-zero logx
- **配置**: YAML

## 🏗 项目特性

- ✅ 清晰的分层架构
- ✅ 统一的错误处理
- ✅ 统一的响应格式
- ✅ JWT 认证中间件
- ✅ CORS 跨域支持
- ✅ 请求日志中间件
- ✅ Docker 容器化支持
- ✅ Kubernetes 部署配置
- ✅ 数据库自动迁移
- ✅ 常用工具函数

## 📝 扩展指南

### 添加新的 API

1. 在 `api/internal/types/` 中定义请求/响应结构
2. 在 `api/internal/logic/` 中编写业务逻辑
3. 在 `api/internal/handler/` 中创建处理器
4. 在 `api/internal/handler/routes.go` 中注册路由

### 添加新的数据模型

1. 在 `model/` 目录下创建新的模型文件
2. 在 `api/internal/svc/servicecontext.go` 中添加自动迁移

### 添加新的中间件

1. 在 `api/internal/middleware/` 中创建中间件
2. 在 `api/internal/handler/routes.go` 中使用中间件

## 📄 License

MIT License

