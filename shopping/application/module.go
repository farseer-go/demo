package application

import (
	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/tasks"
	"shopping/application/job"
	"shopping/domain"
	"time"
)

type Module struct {
}

func (module Module) DependsModule() []modules.FarseerModule {
	return []modules.FarseerModule{domain.Module{}}
}

func (module Module) PostInitialize() {
	// 模拟供应商补货，每分钟补一次，数量0-10
	tasks.Run("ProductReplenishment", 1*time.Minute, job.ProductReplenishmentJob, fs.Context)

	// 注册下单事件
	eventBus.RegisterEvent("buyOrder", job.CreateOrderEvent, job.NoticeWarehouseOrderEvent)
}
