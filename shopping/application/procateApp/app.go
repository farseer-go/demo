package procateApp

import (
	"github.com/farseer-go/collections"
	"shopping/domain/procate"
)

// ToList 读取商品分类列表
// repository：商品分类仓储，webapi自动注入实例
// webapi注入请参考：https://farseer-go.gitee.io/#/web/webapi/container
func ToList(repository procate.Repository) collections.List[procate.DomainObject] {
	return repository.ToList().OrderBy(func(item procate.DomainObject) any {
		return item.Id
	}).ToList()
}
