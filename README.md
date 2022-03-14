# tp-template
tp-micro架构模板，包含基础服务以及网关/认证/常用插件等，通过该模板可以快速搭建一套自己的微服务系统

- (tp-micro version (v2/v2.1.0 branch))[https://github.com/xiaoenai/tp-micro/tree/v2.1.0]

- 目录结构

```
- accessToken 用户请求token
- gateway 网关基础工具
- plugin 插件集合
    - innder_auth 内部授权插件
- tpGateway 网关服务
- account 内部服务
- user 外部服务
```

- 使用

```
1. 启动网关
    cd tpGateway
    go run *.go
2. 启动内部服务
    cd account
    micro run
3. 启动对外api服务
    cd user
    micro run

数据库：test
sql:
CREATE TABLE `gray_match` (
  `uri` varchar(190) NOT NULL COMMENT 'URI',
  `regexp` longtext COMMENT 'regular expression to match UID',
  `updated_at` int(11) NOT NULL COMMENT 'updated time',
  `created_at` int(11) NOT NULL COMMENT 'created time',
  PRIMARY KEY (`uri`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='gray rule';

数据库: tp-user
sql:
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `access_token` varchar(48) NOT NULL DEFAULT '' COMMENT '用户token',
  `updated_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `deleted_ts` bigint(20) NOT NULL DEFAULT '0' COMMENT '删除时间(0表示未删除)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`name`),
  KEY `idx_access_token` (`access_token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';

curl 'http://127.0.0.1:5000/user/v1/user/add?system=ios&app_ver=1.0.0&auth_name_=tp' \
-H 'Content-Type: application/json' \
--data-binary '{"name": "tp"}'

curl 'http://127.0.0.1:5000/user/v1/user/get_by_id?system=ios&app_ver=1.0.0&auth_name_=tp&id=1' \
-H 'Content-Type: application/json'

curl 'http://127.0.0.1:5000/user/v1/user/get?system=ios&app_ver=1.0.0&auth_name_=tp&access_token_=1' \
-H 'Content-Type: application/json'

```