# 全栈用户管理系统 AI 助手指南

此文档帮助 AI 代码助手理解项目架构和开发约定。

## 项目概述

这是一个基于 Gin 框架的用户管理系统，采用经典的 MVC 架构模式。项目使用 Go 1.25.1，主要依赖包括：
- Gin (Web 框架)
- GORM (ORM)
- JWT (认证)
- bcrypt (密码加密)

## 目录结构和职责

```
practice_usermanagement/
├── config/       # 配置管理
├── controller/   # HTTP 请求处理和业务逻辑协调
├── middlewares/  # 中间件（认证等）
├── model/        # 数据模型定义
├── pkg/         
│   ├── database/ # 数据库连接管理
│   ├── e/       # 错误消息定义
│   └── util/    # 通用工具（JWT、响应处理）
└── routes/      # 路由定义
```

## 关键设计模式

1. **数据模型**: 
   - 模型定义在 `model/` 目录，使用 GORM 标签
   - 示例：参考 `model/user.go` 中的 `User` 结构体定义

2. **控制器模式**:
   - 控制器位于 `controller/` 目录
   - 负责：请求解析 -> 业务逻辑调用 -> 响应构造
   - 使用 `util.Response` 统一响应格式

3. **中间件链**:
   - JWT 认证中间件在 `middlewares/auth.go`
   - 在路由注册时配置

## 开发工作流

1. **数据库设置**:
   ```bash
   # 数据库配置在 .env 文件中定义
   DBUSER=root
   DBPASSWORD=your_password
   DBHOST=localhost
   DBPORT=3306
   DBNAME=user_management
   ```

2. **API 文档**:
   - 使用 Swagger 注解记录 API
   - 参考 `controller/user.go` 中的注解示例

## 常见模式

1. **错误处理**:
   - 使用 `pkg/e/msg.go` 中定义的错误码和消息
   - 通过 `util.Response` 返回统一格式响应

2. **认证流程**:
   - JWT 令牌生成和验证在 `pkg/util/jwt.go`
   - 受保护路由需要 `middlewares.AuthMiddleware()`

3. **数据验证**:
   - 使用 Gin 的 binding 标签进行请求验证
   - 在控制器中使用 `ShouldBindJSON` 验证请求

## 集成要点

1. **数据库连接**:
   - 在 `main.go` 启动时通过 `database.InitDB()` 初始化
   - 使用 GORM 处理数据库操作

2. **配置管理**:
   - 环境变量通过 `config/db.go` 加载
   - 使用 godotenv 处理 .env 文件

_注：此文档反映当前代码库的实际实践，如有更新请相应修改。_