package repository

import (
	"fullExample/domain/products/product"
	"github.com/farseer-go/cache"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/container"
)

func RegisterProductRepository() {
	container.Register(func() product.Repository {
		return &productRepository{}
	})
}

type productRepository struct {
	DB cache.ICacheManage[product.DomainObject] `inject:"product"`
}

func (p *productRepository) ToEntity(productId int) product.DomainObject {
	item, _ := p.DB.GetItem(productId)
	return item
}

func (p *productRepository) ToList(pageIndex int, pageSize int) collections.PageList[product.DomainObject] {
	return p.DB.Get().ToPageList(pageSize, pageIndex)
}
