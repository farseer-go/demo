# 完整示例

在这个示例中，我们将使用一个Api服务来展示farseer-go的框架在实际应用中是如何使用的。

## 分层介绍

本示例中，使用的是DDD的方式。各目录的作用：

- interfaces：接口层，主要用于暴露API服务或Web页面
- application：提供各个逻辑层（领域层、上下文）的编排及事务处理
- domain：领域层，行为数据、主要逻辑核心
- infrastructure：基础设施层，对各资源库（数据库、Redis）的访问以及第三方中间件的直接依赖。

![](https://farseer-go.gitee.io/images/farseer-go.png)

## 故事

模拟提供一个小型电商功能。提供：

- 产品目录
- 订单处理
- 库存管理
- 物流通知

本示例主要用来展示如何通过farseer-go来实现应用程序的。上方提到的功能点都是以演示为主，不会考虑实际的因素。

模拟的行为：

1. 用户浏览产品目录
2. 对感兴趣的产品下单购买。
3. 检查库存，库存满足时，将生成订单、并扣除库存数
4. 仓库管理人员对订单发货
5. 模拟推送物流通知。
6. 用户查看订单、查看物流信息

> 为方便用户运行，这里我们不依赖数据库、Redis。使用本地缓存组件来模拟数据库。

## 目录

```shell
|____go.mod
|____go.sum
|____README.md
|____startupModule.go
|____application
| |____module.go
|____infrastructure
| |____module.go
|____domain
| |____module.go
|____main.go
|____interfaces
| |____module.go
```

## 说明

- 2022-12-23：实现了第一接口：查看产品详情
- 2022-12-21：提交框架，定义好各个分层