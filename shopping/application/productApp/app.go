// @area api/1.0
package productApp

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/mapper"
	"shopping/domain/product"
	"shopping/domain/stock"
)

// ToEntity 查看商品详细信息
// repository通过container自动注入实现进来
// @get {area}/product/info
// @filter jwt auth
// @message 查询成功
func ToEntity(productId int64, productRepository product.Repository, stockRepository stock.Repository) DTO { //
	do := productRepository.ToEntity(productId)
	dto := mapper.Single[DTO](do)
	// 获取库存
	dto.Stock = stockRepository.Get(productId)
	return dto
}

// List 获取商品列表
// productRepository：商品仓储，webapi自动注入实例
// stockRepository：库存仓储，webapi自动注入实例
// webapi注入请参考：https://farseer-go.gitee.io/#/web/webapi/container
// @get {area}/product/{action}
// @filter jwt auth
// @message 查询成功
func List(cateId, pageSize, pageIndex int, productRepository product.Repository, stockRepository stock.Repository) collections.PageList[DTO] {
	if pageIndex == 0 {
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	// 从仓储接口获取数据
	lstDO := productRepository.ToPageListByCateId(cateId, pageSize, pageIndex)

	// 转成PageList
	var lstDTO collections.PageList[DTO]
	lstDO.MapToPageList(&lstDTO)

	// 这里为了省事，直接读出所有商品的库存
	stocks := stockRepository.GetAll()
	lstDTO.List.Foreach(func(item *DTO) {
		item.Stock = stocks[item.Id]
	})
	return lstDTO
}

// Buy 购买商品
// productRepository：商品仓储，webapi自动注入实例
// stockRepository：库存仓储，webapi自动注入实例
// webapi注入请参考：https://farseer-go.gitee.io/#/web/webapi/container
// @get {area}/product/{action}
// @filter jwt auth
// @di transaction default
// @di buyOrderEvent buyOrder
// @message 下单成功
func Buy(productId int64, productRepository product.Repository, stockRepository stock.Repository, transaction core.ITransaction, buyOrderEvent core.IEvent) {
	// 减库存，剩余库存>0 ，扣减成功
	stockVal := stockRepository.Set(productId, -1)
	if stockVal > -1 {
		// 把商品信息查出来
		productDO := productRepository.ToEntity(productId)
		// 开启数据库事务
		transaction.Transaction(func() {
			// 发布下单事件
			_ = buyOrderEvent.Publish(&productDO)
		})
	}
}
