package order

import (
	"github.com/farseer-go/data"
)

// Repository 订单仓储
type Repository interface {
	// IRepository 通用的仓储接口
	data.IRepository[DomainObject]
}
