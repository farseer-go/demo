package repository

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/mapper"
	"shopping/domain/product"
	"shopping/infrastructure/repository/context"
	"shopping/infrastructure/repository/model"
)

// InitProduct 注册商品仓储 ioc product.Repository
func InitProduct() {
	container.Register(func() product.Repository {
		return &ProductRepository{}
	})
}

type ProductRepository struct{}

func (p *ProductRepository) ToEntity(productId int64) product.DomainObject {
	po := context.MysqlContextIns.Product.Where("id", productId).ToEntity()
	// po 转 do
	return mapper.Single[product.DomainObject](&po)
}

func (p *ProductRepository) ToPageList(cateId, pageSize, pageIndex int) collections.PageList[product.DomainObject] {
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

func (p *ProductRepository) ToList() collections.List[product.DomainObject] {
	// 从数据库读数据
	lstProduct := context.MysqlContextIns.Product.ToList()
	// po 转 do
	return mapper.ToList[product.DomainObject](lstProduct)
}

func (p *ProductRepository) Count() int64 {
	return context.MysqlContextIns.Product.Count()
}

func (p *ProductRepository) Add(product product.DomainObject) {
	po := mapper.Single[model.ProductPO](&product)
	_ = context.MysqlContextIns.Product.Insert(&po)
}
