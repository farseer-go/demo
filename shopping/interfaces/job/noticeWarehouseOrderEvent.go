package job

import (
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/flog"
	"shopping/domain/product"
)

// NoticeWarehouseOrderEvent 订阅下单事件：通知仓库
func NoticeWarehouseOrderEvent(message any, ea core.EventArgs) {
	productDO := message.(*product.DomainObject)
	// 这里只是一个模拟，为了演示事件驱动的使用场景
	// 一个下单事件，可以订阅多个消费者
	flog.Infof("通知仓库打包商品：%d %s ￥%s元 1件", productDO.Id, productDO.Caption, flog.Red(productDO.Price))
}
