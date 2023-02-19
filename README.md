web_app 是一个用gin实现的 后端论坛论坛源码 采用mvc的架构

数据库文件 [bluebell.sql](bluebell.sql)

配置文件 [config.yaml](config.yaml)

主入口文件 [main.go](main.go)

```
.
├── Dockerfile
├── README.md
├── api
│   ├── docs_models.go
│   └── v1
│       ├── comment.go
│       ├── community.go
│       ├── post.go
│       ├── post_image.go
│       ├── user.go
│       └── vote.go
├── bluebell.sql
├── config
│   ├── captcha.go
│   ├── casbin.go
│   ├── config.go
│   ├── email.go
│   ├── excel.go
│   ├── gorm.go
│   ├── jwt.go
│   ├── minio.go
│   ├── oss.go
│   ├── redis.go
│   ├── snowflake.go
│   ├── system.go
│   └── zap.go
├── config.yaml
├── core
│   ├── jwt.go
│   ├── minio.go
│   ├── mysql.go
│   ├── redis.go
│   ├── snowflake.go
│   ├── viper.go
│   └── zap.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── global
│   ├── code.go
│   ├── error_code.go
│   └── global.go
├── go.mod
├── go.sum
├── latest_log -> log/2022-09-03.log
├── log
│   └── 2022-09-03.log
├── main.go
├── middleware
│   ├── cors.go
│   ├── error.go
│   ├── jwt.go
│   ├── logger.go
│   └── sentinel.go
├── model
│   ├── comment.go
│   ├── community.go
│   ├── post.go
│   ├── redis
│   │   ├── keys.go
│   │   ├── post.go
│   │   └── vote.go
│   ├── request
│   │   ├── comment.go
│   │   ├── comment_rel.go
│   │   ├── common.go
│   │   ├── like.go
│   │   ├── post.go
│   │   ├── user.go
│   │   └── vote.go
│   ├── response
│   │   ├── comment.go
│   │   ├── comment_rel.go
│   │   ├── community.go
│   │   ├── post.go
│   │   └── user.go
│   └── user.go
├── router
│   ├── router_base.go
│   ├── router_comment.go
│   ├── router_community.go
│   ├── router_post.go
│   └── router_user.go
├── service
│   ├── comment.go
│   ├── community.go
│   ├── post.go
│   ├── user.go
│   └── vote.go
├── utils
│   ├── constant.go
│   ├── directory.go
│   ├── md5.go
│   ├── minio.go
│   ├── request.go
│   ├── response.go
│   └── rotatelogs_unix.go
├── validator
│   ├── login_validator.go
│   └── validator.go
└── web_app

16 directories, 86 files


```