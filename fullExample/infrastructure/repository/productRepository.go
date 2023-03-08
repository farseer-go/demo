package repository

import (
	"fullExample/domain/product"
	"fullExample/infrastructure/repository/model"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/data"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/mapper"
)

// InitProduct 注册产品仓储 ioc product.Repository
func InitProduct() {
	container.RegisterTransient(func() product.Repository {
		// 初始化数据库上下文
		// default = farseer.yaml > Database.default
		// true = autoCreateTable
		return data.NewContext[ProductRepository]("default", true)
	})
}

type ProductRepository struct {
	// 定义数据库表映射TableSet
	Product data.TableSet[model.ProductPO] `data:"name=farseer_go_product"`
}

func (p *ProductRepository) ToEntity(productId int64) product.DomainObject {
	po := p.Product.Where("id", productId).ToEntity()
	// po 转 do
	return mapper.Single[product.DomainObject](&po)
}

func (p *ProductRepository) ToPageList(cateId, pageSize, pageIndex int) collections.PageList[product.DomainObject] {
	ts := p.Product.Select("Id", "Caption", "ImgSrc", "Price")
	// 需要筛选产品分类ID
	if cateId > 0 {
		ts.Where("cate_id = ?", cateId)
	}
	// 从数据库读数据
	lstProduct := ts.ToPageList(pageSize, pageIndex)

	// po 转 do
	var lst collections.PageList[product.DomainObject]
	lstProduct.MapToPageList(&lst)
	return lst
}

func (p *ProductRepository) ToList() collections.List[product.DomainObject] {
	// 从数据库读数据
	lstProduct := p.Product.ToList()

	// po 转 do
	return mapper.ToList[product.DomainObject](lstProduct)
}

func (p *ProductRepository) Count() int64 {
	return p.Product.Count()
}

func (p *ProductRepository) Add(product product.DomainObject) {
	po := mapper.Single[model.ProductPO](&product)
	_ = p.Product.Insert(&po)
}
