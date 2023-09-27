package context

import (
	"github.com/farseer-go/data"
	"github.com/farseer-go/fs/core"
	"shopping/domain/order"
	"shopping/domain/procate"
	"shopping/domain/product"
	"shopping/infrastructure/repository/model"
)

// MysqlContext 初始化数据库上下文
var MysqlContext *mysqlContext

// mysqlContext 数据库上下文
type mysqlContext struct {
	// 手动使用事务时必须定义
	core.ITransaction
	// 获取原生ORM框架（不使用TableSet或DomainSet）
	data.IInternalContext
	// 定义数据库表 订单 映射TableSet
	Order data.DomainSet[model.OrderPO, order.DomainObject] `data:"name=farseer_go_order;migrate"`
	// 定义数据库表 商品分类 映射TableSet
	ProCate data.DomainSet[model.ProCatePO, procate.DomainObject] `data:"name=farseer_go_procate;migrate"`
	// 定义数据库表 产品 映射TableSet
	Product data.DomainSet[model.ProductPO, product.DomainObject] `data:"name=farseer_go_product;migrate"`
}

// InitMysqlContext 初始化上下文
func InitMysqlContext() {
	MysqlContext = data.NewContext[mysqlContext]("default")
}
