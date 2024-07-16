package job

import (
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/tasks"
	"math/rand"
	"shopping/domain/product"
	"shopping/domain/stock"
	"time"
)

// ProductReplenishmentJob 模拟供应商补货
func ProductReplenishmentJob(context *tasks.TaskContext) {
	// 读取所有商品
	productRepository := container.Resolve[product.Repository]()
	lstProduct := productRepository.ToList()

	// 遍历商品，给每件商品进货
	stockRepository := container.Resolve[stock.Repository]()
	for i := 0; i < lstProduct.Count(); i++ {
		stockVal := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(5)
		productDO := lstProduct.Index(i)
		stockRepository.Set(productDO.Id, stockVal)

		flog.Infof("商品%d %s 补货%s件", productDO.Id, productDO.Caption, flog.Red(stockVal))
	}
}
