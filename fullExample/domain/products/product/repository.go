package product

import "github.com/farseer-go/collections"

// Repository 产品仓储
type Repository interface {
	// ToEntity 查看产品详细信息
	ToEntity(productId int) DomainObject
	// ToList 获取产品列表
	ToList(pageIndex int, pageSize int) collections.PageList[DomainObject]
}
