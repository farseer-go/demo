package main

import (
	"github.com/farseer-go/fs"
	"github.com/farseer-go/webapi"
)

func main() {
	fs.Initialize[StartupModule]("fss")

	//webapi.RegisterRoutes(routeMeta)
	//webapi.RegisterController(&interfaces.TaskController{})

	webapi.UseApiResponse()
	webapi.Run()
}
