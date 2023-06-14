package order

import (
	"github.com/farseer-go/data/repository"
)

// Repository 订单仓储
type Repository interface {
	// IRepository 通用的仓储接口
	repository.IRepository[DomainObject]
}
