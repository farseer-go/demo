package product

import "github.com/farseer-go/collections"

// Repository 产品仓储
type Repository interface {
	// ToEntity 查看产品详细信息
	ToEntity(productId int64) DomainObject
	// ToPageList 获取产品列表
	ToPageList(cateId, pageSize, pageIndex int) collections.PageList[DomainObject]
	// ToList 获取所有产品
	ToList() collections.List[DomainObject]
	// Count 获取数量
	Count() int64
	// Add 添加产品
	Add(product DomainObject)
}
