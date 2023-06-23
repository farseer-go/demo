package context

import (
	"github.com/farseer-go/data"
	"shopping/domain/order"
	"shopping/domain/procate"
	"shopping/domain/product"
	"shopping/infrastructure/repository/model"
)

// MysqlContextIns 初始化数据库上下文
var MysqlContextIns *MysqlContext

// MysqlContext 数据库上下文
type MysqlContext struct {
	// 手动使用事务时必须定义
	//core.ITransaction
	// 定义数据库表 订单 映射TableSet
	Order data.DomainSet[model.OrderPO, order.DomainObject] `data:"name=farseer_go_order"`
	// 定义数据库表 商品分类 映射TableSet
	ProCate data.DomainSet[model.ProCatePO, procate.DomainObject] `data:"name=farseer_go_procate"`
	// 定义数据库表 产品 映射TableSet
	Product data.DomainSet[model.ProductPO, product.DomainObject] `data:"name=farseer_go_product"`
}

// InitMysqlContext 初始化上下文
func InitMysqlContext() {
	MysqlContextIns = data.NewContext[MysqlContext]("default", true)
}
