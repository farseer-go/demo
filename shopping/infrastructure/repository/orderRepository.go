package repository

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/mapper"
	"shopping/domain/order"
	"shopping/infrastructure/repository/context"
	"shopping/infrastructure/repository/model"
)

// InitOrder 注册商品仓储 ioc Order.Repository
func InitOrder() {
	container.Register(func() order.Repository {
		return &OrderRepository{}
	})
}

// OrderRepository 订单仓储
type OrderRepository struct{}

func (p *OrderRepository) ToEntity(OrderId int64) order.DomainObject {
	po := context.MysqlContextIns.Order.Where("id", OrderId).ToEntity()
	// po 转 do
	return mapper.Single[order.DomainObject](&po)
}

func (p *OrderRepository) ToPageList(pageSize, pageIndex int) collections.PageList[order.DomainObject] {
	// 从数据库读数据
	lstOrder := context.MysqlContextIns.Order.Desc("create_at").ToPageList(pageSize, pageIndex)

	// po 转 do
	var lst collections.PageList[order.DomainObject]
	lstOrder.MapToPageList(&lst)
	return lst
}

func (p *OrderRepository) Count() int64 {
	return context.MysqlContextIns.Order.Count()
}

func (p *OrderRepository) Add(Order order.DomainObject) {
	po := mapper.Single[model.OrderPO](&Order)

	//mysqlContext := context.MysqlContextIns
	//mysqlContext.Begin()
	//_ = mysqlContext.Order.Insert(&po)
	//mysqlContext.Commit()

	//context.MysqlContextIns.Transaction(func() {
	_ = context.MysqlContextIns.Order.Insert(&po)
	//})
}
