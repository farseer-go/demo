package interfaces

import (
	"shopping/application"
	"shopping/interfaces/job"
	"time"

	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/tasks"
	"github.com/farseer-go/webapi"
)

type Module struct {
}

func (module Module) DependsModule() []modules.FarseerModule {
	return []modules.FarseerModule{webapi.Module{}, application.Module{}}
}

func (module Module) PostInitialize() {
	// 模拟供应商补货，每分钟补一次，数量0-10
	tasks.RunNow("ProductReplenishment", 1*time.Minute, job.ProductReplenishmentJob, fs.Context)

	// 注册下单事件
	eventBus.RegisterEvent("buyOrder").
		RegisterSubscribe("创建订单", job.CreateOrderEvent).
		RegisterSubscribe("通知仓库", job.NoticeWarehouseOrderEvent)
}
