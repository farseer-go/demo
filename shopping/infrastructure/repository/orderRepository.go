package repository

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/data/repository"
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
type OrderRepository struct {
	repository.IRepository[order.DomainObject]
}

func (p *OrderRepository) ToPageList(pageSize, pageIndex int) collections.PageList[order.DomainObject] {
	// 从数据库读数据
	lstOrder := context.MysqlContextIns.Order.Desc("create_at").ToPageList(pageSize, pageIndex)

	// po 转 do
	var lst collections.PageList[order.DomainObject]
	lstOrder.MapToPageList(&lst)
	return lst
}
