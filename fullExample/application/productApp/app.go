package productApp

import (
	"fullExample/domain/product"
	"fullExample/domain/stock"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/mapper"
)

// ToEntity 查看产品详细信息
// repository通过container自动注入实现进来
func ToEntity(productId int64, repository product.Repository) DTO {
	do := repository.ToEntity(productId)
	dto := mapper.Single[DTO](do)
	return dto
}

// ToList 获取产品列表
// productRepository：产品仓储，webapi自动注入实例
// stockRepository：库存仓储，webapi自动注入实例
// webapi注入请参考：https://farseer-go.gitee.io/#/web/webapi/container
func ToList(cateId, pageSize, pageIndex int, productRepository product.Repository, stockRepository stock.Repository) collections.PageList[DTO] {
	if pageIndex == 0 {
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	// 从仓储接口获取数据
	lstDO := productRepository.ToPageList(cateId, pageSize, pageIndex)

	// 转成PageList
	var lstDTO collections.PageList[DTO]
	lstDO.MapToPageList(&lstDTO)

	// 这里为了省事，直接读出所有商品的库存
	stocks := stockRepository.GetAll()
	for i := 0; i < lstDTO.List.Count(); i++ {
		item := lstDTO.List.Index(i)
		item.Stock = stocks[item.Id]
		lstDTO.List.Set(i, item)
	}
	return lstDTO
}
