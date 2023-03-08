package job

import (
	"fullExample/domain/order"
	"fullExample/domain/product"
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/snowflake"
	"time"
)

// CreateOrderEvent 订阅下单事件：创建订单
func CreateOrderEvent(message any, ea core.EventArgs) {
	productDO := message.(*product.DomainObject)
	orderRepository := container.Resolve[order.Repository]()
	orderRepository.Add(order.DomainObject{
		Id:             snowflake.GenerateId(),
		ProductId:      productDO.Id,
		ProductCaption: productDO.Caption,
		ProductImgSrc:  productDO.ImgSrc,
		ProductPrice:   productDO.Price,
		ProductCount:   1,
		CreateAt:       time.Now(),
		CreateId:       fs.AppId,
	})
}
