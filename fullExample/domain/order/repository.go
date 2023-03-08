package order

import (
	"github.com/farseer-go/collections"
)

// Repository 订单仓储
type Repository interface {
	// ToEntity 查看订单信息
	ToEntity(OrderId int64) DomainObject
	// ToPageList 订单列表
	ToPageList(pageSize, pageIndex int) collections.PageList[DomainObject]
	// Count 订单数量
	Count() int64
	// Add 添加订单
	Add(Order DomainObject)
}
