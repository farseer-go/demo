package repository

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/data"
	"github.com/farseer-go/fs/container"
	"shopping/domain/order"
	"shopping/infrastructure/repository/context"
)

// InitOrder 注册商品仓储 ioc Order.Repository
func InitOrder() {
	container.Register(func() order.Repository {
		return &OrderRepository{}
	})
}

// OrderRepository 订单仓储
// @register order.Repository
type OrderRepository struct {
	data.IRepository[order.DomainObject]
}

func (p *OrderRepository) ToPageList(pageSize, pageIndex int) collections.PageList[order.DomainObject] {
	// 从数据库读数据
	lstOrder := context.MysqlContext.Order.Desc("create_at").ToPageList(pageSize, pageIndex)

	// po 转 do
	var lst collections.PageList[order.DomainObject]
	lstOrder.MapToPageList(&lst)
	return lst
}
