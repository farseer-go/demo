package main

import (
	"fullExample/application/products/productApp"
	"github.com/farseer-go/fs"
	"github.com/farseer-go/webapi"
)

func main() {
	fs.Initialize[StartupModule]("demo")

	webapi.Area("/api/1.0/", func() {
		// get http://localhost:8888/api/1.0/product/info?productId=1
		webapi.RegisterGET("/product/info", productApp.ToEntity)

		// get http://localhost:8888/api/1.0/product/list?pageIndex=1&pageSize=3
		webapi.RegisterGET("/product/list", productApp.ToList, "pageIndex", "pageSize", "repository")
	})
	webapi.UseApiResponse()
	webapi.Run()
}
