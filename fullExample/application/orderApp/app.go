package orderApp

import (
	"fullExample/domain/order"
	"github.com/farseer-go/collections"
)

// ToList 获取订单列表
// repository：订单仓储，webapi自动注入实例
// webapi注入请参考：https://farseer-go.gitee.io/#/web/webapi/container
func ToList(pageSize, pageIndex int, repository order.Repository) collections.PageList[order.DomainObject] {
	if pageIndex == 0 {
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	// 从仓储接口获取数据
	return repository.ToPageList(pageSize, pageIndex)
}

// Count 获取订单数量
// repository：订单仓储，webapi自动注入实例
// webapi注入请参考：https://farseer-go.gitee.io/#/web/webapi/container
func Count(repository order.Repository) int64 {
	return repository.Count()
}
