package infrastructure

import (
	"fmt"
	"github.com/farseer-go/fs/modules"
	"github.com/olivere/elastic/v7"
)

type Module struct {
}

func (module Module) DependsModule() []modules.FarseerModule {
	// 使用到了redis模块、data(orm)模块、eventBus（事件总线）模块、queue（本地队列）模块
	// 这些模块都是farseer-go内置的模块
	return []modules.FarseerModule{}
}

func (module Module) PostInitialize() {
	//lst := collections.NewList("")
	//fmt.Println(lst.Count())
	fmt.Println(elastic.Client{})
	//validator.New(validator.WithRequiredStructEnabled())
}

// fs 4.5
// collections 5.2
// cacheMemory 4.6
// elasticSearch 14.4
// etcd 19.2
// eventBus 4.6
// fSchedule 18
// linkTrace 16.3
// mapper 5.3
// queue 4.6
// rabbit 9.2
// redis 11.2
// tasks 4.5
// utils 4.5
// webapi 11.9  7.2
// data 31.9
