package procateApp

import (
	"fullExample/domain/procate"
	"github.com/farseer-go/collections"
)

// ToList 读取商品分类列表
func ToList(repository procate.Repository) collections.List[procate.DomainObject] {
	return repository.ToList().OrderBy(func(item procate.DomainObject) any {
		return item.Id
	}).ToList()
}
