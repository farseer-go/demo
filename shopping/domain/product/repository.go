package product

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/data"
)

// Repository 商品仓储
type Repository interface {
	// IRepository 通用的仓储接口
	data.IRepository[DomainObject]
	// ToPageListByCateId 获取商品列表
	ToPageListByCateId(cateId, pageSize, pageIndex int) collections.PageList[DomainObject]
}
