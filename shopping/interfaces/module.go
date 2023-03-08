package interfaces

import (
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/webapi"
	"shopping/application"
)

type Module struct {
}

func (module Module) DependsModule() []modules.FarseerModule {
	return []modules.FarseerModule{webapi.Module{}, application.Module{}}
}
