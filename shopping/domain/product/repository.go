package product

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/data/repository"
)

// Repository 商品仓储
type Repository interface {
	// IRepository 通用的仓储接口
	repository.IRepository[DomainObject]
	// ToPageListByCateId 获取商品列表
	ToPageListByCateId(cateId, pageSize, pageIndex int) collections.PageList[DomainObject]
}
