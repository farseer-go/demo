package repository

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/data"
	"github.com/farseer-go/fs/container"
	"shopping/domain/product"
	"shopping/infrastructure/repository/context"
)

// InitProduct 注册商品仓储 ioc product.Repository
func InitProduct() {
	container.Register(func() product.Repository {
		return &ProductRepository{}
	})
}

type ProductRepository struct {
	data.IRepository[product.DomainObject]
}

func (p *ProductRepository) ToPageListByCateId(cateId, pageSize, pageIndex int) collections.PageList[product.DomainObject] {
	// 需要筛选商品分类ID
	lstProduct := context.MysqlContextIns.Product.
		Select("Id", "Caption", "ImgSrc", "Price").
		WhereIgnoreLessZero("cate_id = ?", cateId).
		ToPageList(pageSize, pageIndex)

	// po 转 do
	var lst collections.PageList[product.DomainObject]
	lstProduct.MapToPageList(&lst)
	return lst
}
