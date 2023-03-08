package repository

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/data"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/mapper"
	"shopping/domain/order"
	"shopping/infrastructure/repository/model"
)

// InitOrder 注册商品仓储 ioc Order.Repository
func InitOrder() {
	container.RegisterTransient(func() order.Repository {
		// 初始化数据库上下文
		// default = farseer.yaml > Database.default
		// true = autoCreateTable
		return data.NewContext[OrderRepository]("default", true)
	})
}

type OrderRepository struct {
	// 定义数据库表映射TableSet
	Order data.TableSet[model.OrderPO] `data:"name=farseer_go_order"`
}

func (p *OrderRepository) ToEntity(OrderId int64) order.DomainObject {
	po := p.Order.Where("id", OrderId).ToEntity()
	// po 转 do
	return mapper.Single[order.DomainObject](&po)
}

func (p *OrderRepository) ToPageList(pageSize, pageIndex int) collections.PageList[order.DomainObject] {
	// 从数据库读数据
	lstOrder := p.Order.Desc("create_at").ToPageList(pageSize, pageIndex)

	// po 转 do
	var lst collections.PageList[order.DomainObject]
	lstOrder.MapToPageList(&lst)
	return lst
}

func (p *OrderRepository) Count() int64 {
	return p.Order.Count()
}

func (p *OrderRepository) Add(Order order.DomainObject) {
	po := mapper.Single[model.OrderPO](&Order)
	_ = p.Order.Insert(&po)
}
