# 完整示例
完整展示farseer-go框架的使用案例

模拟一个小型电商平台，用到的技术：
* ddd：使用领域驱动设计
* ioc：使用ioc/container，做解耦、注入、依赖倒置
* webapi：api服务，并使用动态api技术
* data：数据库操作
* redis：redis操作
* eventBus：事件驱动

包含：
* 首页：按商品分类，浏览商品
* 商品：商品详情，下单
* 订单：下订单后，通过事件驱动来通知商品减库存、保存订单数据
  ![img.png](https://farseer-go.github.io/doc/images/shopping_1.png)
  ![img.png](https://farseer-go.github.io/doc/images/shopping_2.png)

## 分层介绍
各目录的作用：
- /application/：提供各个逻辑层（领域层、上下文）的编排及事务处理
  - /job/
    - createOrderEvent.go：订阅下单事件：创建订单
    - noticeWarehouseOrderEvent.go：订阅下单事件：通知仓库
    - productReplenishmentJob.go：模拟供应商补货
  - /orderApp/
    - app.go：订单应用服务
  - /procateApp/
    - app.go：商品分类应用服务
  - /productApp/
    - app.go：商品应用服务
    - dto.go：dto
  - module.go：应用层模块
- /domain/：领域层，行为数据、主要逻辑核心
  - /order/
    - domainObject.go：订单领域层
    - repository.go：订单仓储接口
  - /procate/
    - domainObject.go：商品分类领域层
    - repository.go：商品分类仓储接口
  - /product/
    - cateEO.go：商品分类实体类
    - domainObject.go：商品领域层
    - repository.go：商品仓储接口
  - /stock/
    - repository.go：库存仓储接口
  - module.go：领域层模块
- /infrastructure/：基础设施层，对各资源库（数据库、Redis）的访问以及第三方中间件的直接依赖。
  - /repository/
    - /model/
      - orderPO.go：订单PO
      - proCatePO.go：商品分类PO
      - productPO.go：商品PO
  - orderRepository.go：订单仓储实现
  - proCateRepository.go：商品分类仓储实现
  - productRepository.go：商品仓储实现
  - stockRepository.go：库存仓储实现
  - module.go：基础设施层模块
- /interfaces/：接口层，主要用于暴露API服务或Web页面
  - module.go：接口层模块
- /wwwroot/：静态页面
- startupModule.go：启动模块
- farseer.yaml：配置文件

![](https://farseer-go.github.io/doc/images/farseer-go.png)

## 数据库表
### ProductPO 商品
表名：`farseer_go_product`

> 技术点：实现了数据库的插入、查找、分页技术

### ProCatePO 商品分类
表名：`farseer_go_procate`

> 技术点：与Redis缓存自动同步，redis没数据时自动从数据库读取并缓存到redis

### orderPO 订单信息
表名：`farseer_go_order`

> 技术点：利用事件驱动，解耦下单时，商品、订单的耦合

## Redis缓存
### 库存
key：`farseer_go_stock`

> 保存商品的库存，这里不考虑库存并发等问题，只演示如何使用Redis模块

### 商品分类
key：`procate`

> 数据库表`farseer_go_procate`与Redis同步

## 运行
需自行配置数据库、Redis环境：到farseer.yaml配置文件中，配置`Database.default`、`Redis.default`