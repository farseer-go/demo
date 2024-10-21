package repository

import (
	"shopping/domain/product"
	"shopping/infrastructure/repository/context"

	"github.com/farseer-go/collections"
	"github.com/farseer-go/data"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/mapper"
)

// InitProduct 注册商品仓储 ioc product.Repository
func InitProduct() {
	container.Register(func() product.Repository {
		return &ProductRepository{}
	})
}

// ProductRepository 产品仓储
// @inject product.Repository
type ProductRepository struct {
	data.IRepository[product.DomainObject]
}

func (p *ProductRepository) ToPageListByCateId(cateId, pageSize, pageIndex int) collections.PageList[product.DomainObject] {
	// 需要筛选商品分类ID
	lstProduct := context.MysqlContext.Product.
		Select("Id", "Caption", "ImgSrc", "Price").
		WhereIgnoreLessZero("cate_id = ?", cateId).
		ToPageList(pageSize, pageIndex)

	// po 转 do
	return mapper.ToPageList[product.DomainObject](lstProduct)
}
