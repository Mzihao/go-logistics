# 物流聚合服务

## 描述
一个聚合多个物流公司的，集成物流下单，查询的服务。
已有物流商:
1. [zeek](https://www.zeek.one): 斑马到家（HK）
2. [Maple Logistics](http://www.25431010.tw/Search.php)：枫叶物流（TW）
3. [sf-express](https://www.sf-express.com/): 顺丰(未完善)
4. [kuaidi100](https://www.kuaidi100.com):快递100（国内聚合）

## Version
go 1.17

## 组件
1. [Gin](https://github.com/gin-gonic/gin): 轻量级Web框架，自称路由速度是golang最快的 
2. [GORM](https://gorm.io/index.html): ORM工具。本项目需要配合Mysql使用 
3. [Swagger](https://github.com/swaggo/swag): Swagger是一个规范和完整的框架，用于生成、描述、调用和可视化 RESTful 风格的 Web 服务。
4. [Gin-Cors](https://github.com/gin-contrib/cors): Gin框架提供的跨域中间件。
5. [jwt-go](https://github.com/dgrijalva/jwt-go): 用户鉴权数字签名。
6. [logrus](https://github.com/sirupsen/logrus): 日志服务。
7. [viper](https://github.com/spf13/viper): 配置文件解析。

## 功能模块
1. 用户管理模块
2. 第三方物流查询、下单
3. 自提物流模块
4. ...

## Swagger
#### 下载Swag for Go
```shell
go get -u github.com/swaggo/swag/cmd/swag
```

#### 生成所需的文件和文件夹

```shell
swag init
swag init --parseDependency -- // 可加载外部包
```

## 运行
```shell
go run main.go
```

## 预览
[http://127.0.0.1:9090](http://127.0.0.1:9090)

[http://127.0.0.1:9090/swagger/index.html ](http://127.0.0.1:9090/swagger/index.html )

## 演示
#### 测试页面
![demo1](https://github.com/Mzihao/go-logistics/blob/master/demo/demo.png)
![demo2](https://github.com/Mzihao/go-logistics/blob/master/demo/demo2.png)
