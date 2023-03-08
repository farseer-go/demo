package procate

import (
	"github.com/farseer-go/collections"
)

// Repository 商品分类仓储
type Repository interface {
	// ToEntity 获取商品分类
	ToEntity(cateId int) DomainObject
	// ToList 获取商品分类列表
	ToList() collections.List[DomainObject]
	// Count 获取商品分类数量
	Count() int
	// Add 添加商品分类
	Add(product DomainObject)
}
