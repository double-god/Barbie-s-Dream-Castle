# Go 用户管理 API (practice_usermanagement)

这是一个使用 Go 语言 (Gin + GORM) 构建的后端 API 项目，用于实现用户的注册、登录和信息管理。

## ✨ 核心功能

* **用户认证**:
    * `POST /api/register`: 注册新用户 (密码使用 bcrypt 加密)。
    * `POST /api/login`: 用户登录，成功后返回 JWT。
* **用户管理 (需JWT鉴权)**:
    * `GET /api/user/:uid`: 查询指定 ID 用户的公开信息。
    * `PUT /api/user`: 更新当前登录用户自己的信息。
    * `GET /api/user/search`: 根据用户名模糊搜索用户列表 (支持分页)。

## 🚀 技术栈

* **Go**: 编程语言
* **Gin**: Web 框架
* **GORM**: ORM 库
* **MySQL**: 关系型数据库
* **JWT (jwt-go)**: Token 认证
* **godotenv**: 环境变量管理

