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
## 文章（不需要鉴权访问）
## 文章（需要鉴权访问）
## 评论（需要鉴权认证）
## 用户信息（需要认证）



