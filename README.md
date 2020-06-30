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

curl 'http://127.0.0.1:5000/user/v1/user/add?system=ios&app_ver=1.0.0&auth_name_=tp' \
-H 'Content-Type: application/json' \
--data-binary '{"name": "tp"}'

curl 'http://127.0.0.1:5000/user/v1/user/get_by_id?system=ios&app_ver=1.0.0&auth_name_=tp&id=1' \
-H 'Content-Type: application/json'

curl 'http://127.0.0.1:5000/user/v1/user/get?system=ios&app_ver=1.0.0&auth_name_=tp&access_token_=1' \
-H 'Content-Type: application/json'

```