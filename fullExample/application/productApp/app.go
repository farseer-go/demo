package productApp

import (
	"fullExample/domain/product"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/mapper"
)

// ToEntity 查看产品详细信息
// repository通过container自动注入实现进来
func ToEntity(productId int, repository product.Repository) DTO {
	do := repository.ToEntity(productId)
	dto := mapper.Single[DTO](do)
	return dto
}

func ToList(cateId, pageSize, pageIndex int, repository product.Repository) collections.PageList[DTO] {
	if pageIndex == 0 {
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	lstDO := repository.ToList(cateId, pageSize, pageIndex)
	var lstDTO collections.PageList[DTO]
	lstDO.MapToPageList(&lstDTO)
	return lstDTO
}
