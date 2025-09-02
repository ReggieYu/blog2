# blog demo 介绍
该demo项目是基于Go1.23版本开发的简易博客系统

# 技术栈
- Go1.23
- web矿机：Gin
- ORM框架：GORM
- 认证：JWT
- 数据库：SQLite（默认）/ Mysql
- 密码加密：bcrypt

# 安装与运行
```bash
git clone https://github.com/ReggieYu/blog2.git
cd yourrepo
go run main.go
```

# 功能列表
- 用户注册与登录
- 文章发布与管理
- 评论功能

# 项目目录结构
```
blog2/
├── controllers/
├── config/
├── database/
├── dto/
├── middleware/
├── responses/
├── model/
├── router/
├── static/
├── main.go
└── README.md
```

# 快速开始
## 1.安装依赖
go mod tidy

## 2.配置环境（可选）
创建.env文件来覆盖默认配置：
```
# 数据库配置 - 默认使用SQLite
DB_DRIVER=sqlite
SQLITE_PATH=blog.db

# 服务器配置
PORT=8080

# JWT配置
JWT_SECRET=your_secret_key
JWT_TTL_HOURS=24

# 环境配置
GIN_MODE=debug
```
**注意**：如果不创建.env文件，应用将使用以下默认配置:
* 数据库：SQLite
* 数据库文件：blog.db(当前项目根目录下)
* 端口：8080

## 3.运行应用
```
go run main.go
```
或者构建之后运行可执行文件：
```
go buil
./blog
```

## 4.访问API
服务在启动之后，默认可以通过localhost:8080来访问

# SQLite数据库说明
## 默认配置
- 数据库类型：SQLite
- 数据库文件：blog.db(自动创建)
- 位置：项目根目录
- 迁移：自动创建表结构

## SQLite优势
- 无需安装数据库服务器
- 无需配置，直接使用
- 单文件数据库，便于部署
- 支持事务和关系查询
- 适合开发中小型应用

## 数据库文件
启动应用后，会在项目根目录下自动创建blog.db文件，包含：
- users表 -用户信息表
- posts表 -博客文章表
- comments表 -文章评论表

# API端点
## 认证
### 注册
请求
```
curl --location 'localhost:8080/api/v1/auth/register?username=test1&email=test1%40email.com&password=123456' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "test1",
    "email": "test1@email.com",
    "password": "123456"
}'
```

成功响应结果
```
{
    "message": "registered"
}
```

### 登录
请求
```
curl --location 'localhost:8080/api/v1/auth/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "test1",
    "password": "123456"
}'
```

响应
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6InRlc3QxIiwic3ViIjoiMSIsImV4cCI6MTc1Njg5NzQ1OSwiaWF0IjoxNzU2ODExMDU5fQ.kklHGbo_EJ4pGbsYrvyOccYdNsG2RdL0quTpiXOiHVE"
}
```
## 文章（不需要鉴权访问）
### 查询文章列表
请求
```
curl --location 'localhost:8080/api/v1/posts'
```

响应
```
[
    {
        "ID": 1,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "title": "test",
        "content": "testContent",
        "user_id": 1,
        "author": {
            "ID": 1,
            "CreatedAt": "2025-09-02T17:38:06.7042482+08:00",
            "UpdatedAt": "2025-09-02T17:38:06.7042482+08:00",
            "DeletedAt": null,
            "username": "test1",
            "email": "test1@email.com"
        }
    }
]
```

### 根据文章id请求文章
请求
```
curl --location 'localhost:8080/api/v1/posts/2'
```
响应
```
{
    "ID": 2,
    "CreatedAt": "2025-09-02T19:29:36.7973033+08:00",
    "UpdatedAt": "2025-09-02T19:31:30.9758146+08:00",
    "DeletedAt": null,
    "title": "testPost2",
    "content": "testContent2",
    "user_id": 1,
    "author": {
        "ID": 1,
        "CreatedAt": "2025-09-02T17:38:06.7042482+08:00",
        "UpdatedAt": "2025-09-02T17:38:06.7042482+08:00",
        "DeletedAt": null,
        "username": "test1",
        "email": "test1@email.com"
    }
}
```
## 文章（需要鉴权访问）
### 查询我的文章
请求
```
curl --location 'localhost:8080/api/v1/posts' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6InRlc3QxIiwic3ViIjoiMSIsImV4cCI6MTc1Njg5NzQ1OSwiaWF0IjoxNzU2ODExMDU5fQ.kklHGbo_EJ4pGbsYrvyOccYdNsG2RdL0quTpiXOiHVE' \
--header 'Content-Type: application/json' \
--data '{
    "title": "testPost",
    "content": "testContent"
}'
```

响应
```
{
    "ID": 2,
    "CreatedAt": "2025-09-02T19:29:36.7973033+08:00",
    "UpdatedAt": "2025-09-02T19:29:36.7973033+08:00",
    "DeletedAt": null,
    "title": "testPost",
    "content": "testContent",
    "user_id": 1,
    "author": {
        "ID": 1,
        "CreatedAt": "2025-09-02T17:38:06.7042482+08:00",
        "UpdatedAt": "2025-09-02T17:38:06.7042482+08:00",
        "DeletedAt": null,
        "username": "test1",
        "email": "test1@email.com"
    }
}
```

### 更新文章
请求
```
curl --location --request PUT 'localhost:8080/api/v1/posts/2' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6InRlc3QxIiwic3ViIjoiMSIsImV4cCI6MTc1Njg5NzQ1OSwiaWF0IjoxNzU2ODExMDU5fQ.kklHGbo_EJ4pGbsYrvyOccYdNsG2RdL0quTpiXOiHVE' \
--header 'Content-Type: application/json' \
--data '{
    "title": "testPost2",
    "content": "testContent2"
}'
```

响应
```
{
    "ID": 2,
    "CreatedAt": "2025-09-02T19:29:36.7973033+08:00",
    "UpdatedAt": "2025-09-02T19:31:30.9758146+08:00",
    "DeletedAt": null,
    "title": "testPost2",
    "content": "testContent2",
    "user_id": 1,
    "author": {
        "ID": 1,
        "CreatedAt": "2025-09-02T17:38:06.7042482+08:00",
        "UpdatedAt": "2025-09-02T17:38:06.7042482+08:00",
        "DeletedAt": null,
        "username": "test1",
        "email": "test1@email.com"
    }
}
```

## 评论（需要鉴权认证）
### 创建评论
请求
```
curl --location 'localhost:8080/api/v1/posts/2/comments' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6InRlc3QxIiwic3ViIjoiMSIsImV4cCI6MTc1Njg5NzQ1OSwiaWF0IjoxNzU2ODExMDU5fQ.kklHGbo_EJ4pGbsYrvyOccYdNsG2RdL0quTpiXOiHVE' \
--header 'Content-Type: application/json' \
--data '{
    "content": "comment3"
}'
```

响应
```
{
    "ID": 2,
    "CreatedAt": "2025-09-02T19:33:55.3129899+08:00",
    "UpdatedAt": "2025-09-02T19:33:55.3129899+08:00",
    "DeletedAt": null,
    "content": "comment3",
    "user_id": 1,
    "author": {
        "ID": 1,
        "CreatedAt": "2025-09-02T17:38:06.7042482+08:00",
        "UpdatedAt": "2025-09-02T17:38:06.7042482+08:00",
        "DeletedAt": null,
        "username": "test1",
        "email": "test1@email.com"
    },
    "post_id": 2
}
```

### 查询文章及评论
请求
```
curl --location 'localhost:8080/api/v1/posts/2/comments'
```

响应
```
[
    {
        "ID": 1,
        "CreatedAt": "2025-09-02T19:33:38.5395997+08:00",
        "UpdatedAt": "2025-09-02T19:33:38.5395997+08:00",
        "DeletedAt": null,
        "content": "comment2",
        "user_id": 1,
        "author": {
            "ID": 1,
            "CreatedAt": "2025-09-02T17:38:06.7042482+08:00",
            "UpdatedAt": "2025-09-02T17:38:06.7042482+08:00",
            "DeletedAt": null,
            "username": "test1",
            "email": "test1@email.com"
        },
        "post_id": 2
    },
    {
        "ID": 2,
        "CreatedAt": "2025-09-02T19:33:55.3129899+08:00",
        "UpdatedAt": "2025-09-02T19:33:55.3129899+08:00",
        "DeletedAt": null,
        "content": "comment3",
        "user_id": 1,
        "author": {
            "ID": 1,
            "CreatedAt": "2025-09-02T17:38:06.7042482+08:00",
            "UpdatedAt": "2025-09-02T17:38:06.7042482+08:00",
            "DeletedAt": null,
            "username": "test1",
            "email": "test1@email.com"
        },
        "post_id": 2
    }
]
```

## 用户信息（需要认证）
请求
```
curl --location 'localhost:8080/api/v1/me' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6InRlc3QxIiwic3ViIjoiMSIsImV4cCI6MTc1Njg5NzQ1OSwiaWF0IjoxNzU2ODExMDU5fQ.kklHGbo_EJ4pGbsYrvyOccYdNsG2RdL0quTpiXOiHVE'
```

响应
```
{
    "ID": 1,
    "CreatedAt": "2025-09-02T17:38:06.7042482+08:00",
    "UpdatedAt": "2025-09-02T17:38:06.7042482+08:00",
    "DeletedAt": null,
    "username": "test1",
    "email": "test1@email.com"
}
```


