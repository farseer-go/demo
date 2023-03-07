# demo
完整展示farseer-go框架的使用案例

模拟一个小型电商平台，用到的技术：
* ddd：使用ddd分层
* ioc：使用ioc/container，做解耦、注入、依赖倒置
* webapi：api服务，并使用动态api技术
* data：数据库操作
* redis：redis操作
* eventBus：事件驱动

包含：
* 首页：按产品分类，浏览商品
* 商品：商品详情，下单
* 订单：下订单后，通过事件驱动来通知产品减库存、保存订单数据

## 数据库表
### ProductPO 产品
表名：`farseer_go_product`

> 技术点：实现了数据库的插入、查找、分页技术

### ProCatePO 产品分类
表名：`farseer_go_procate`

> 技术点：与Redis缓存自动同步，redis没数据时自动从数据库读取并缓存到redis