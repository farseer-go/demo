package procate

import (
	"github.com/farseer-go/collections"
)

// Repository 产品分类仓储
type Repository interface {
	// ToEntity 获取产品分类
	ToEntity(cateId int) DomainObject
	// ToList 获取产品分类列表
	ToList() collections.List[DomainObject]
	// Count 获取产品分类数量
	Count() int
	// Add 添加产品分类
	Add(product DomainObject)
}
