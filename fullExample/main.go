package main

import (
	"fullExample/application/products/productApp"
	"github.com/farseer-go/fs"
	"github.com/farseer-go/webapi"
)

func main() {
	fs.Initialize[StartupModule]("demo")

	webapi.Area("/api/1.0/", func() {
		webapi.RegisterGET("/product/info", productApp.ToEntity)
		webapi.RegisterGET("/product/info/{productId}", productApp.ToEntity)
		webapi.RegisterGET("/product/list-{pageSize}-{pageIndex}", productApp.ToList, "pageIndex", "pageSize", "repository")
	})
	webapi.UseApiResponse()
	webapi.Run()
}
