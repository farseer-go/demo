package job

import (
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/sonyflake"
	"shopping/domain/order"
	"shopping/domain/product"
	"time"
)

// CreateOrderEvent 订阅下单事件：创建订单
func CreateOrderEvent(message any, ea core.EventArgs) {
	productDO := message.(*product.DomainObject)
	orderRepository := container.Resolve[order.Repository]()
	orderRepository.Add(order.DomainObject{
		Id:             sonyflake.GenerateId(),
		ProductId:      productDO.Id,
		ProductCaption: productDO.Caption,
		ProductImgSrc:  productDO.ImgSrc,
		ProductPrice:   productDO.Price,
		ProductCount:   1,
		CreateAt:       time.Now(),
		CreateId:       core.AppId,
	})
}
