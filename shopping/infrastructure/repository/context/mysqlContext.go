package context

import (
	"github.com/farseer-go/data"
	"github.com/farseer-go/data/repository"
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
	Order data.TableSet[model.OrderPO] `data:"name=farseer_go_order"`
	// 定义数据库表 商品分类 映射TableSet
	ProCate data.TableSet[model.ProCatePO] `data:"name=farseer_go_procate"`
	// 定义数据库表 产品 映射TableSet
	Product data.TableSet[model.ProductPO] `data:"name=farseer_go_product"`
}

// InitMysqlContext 初始化上下文
func InitMysqlContext() {
	MysqlContextIns = data.NewContext[MysqlContext]("default", true)
	// 注册通用的仓储接口
	repository.RegisterRepository[order.DomainObject](MysqlContextIns.Order)
	repository.RegisterRepository[procate.DomainObject](MysqlContextIns.ProCate)
	repository.RegisterRepository[product.DomainObject](MysqlContextIns.Product)
}
