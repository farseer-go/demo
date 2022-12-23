package main

import (
	"fullExample/application/products/productApp"
	"github.com/farseer-go/fs"
	"github.com/farseer-go/webapi"
)

func main() {
	fs.Initialize[StartupModule]("demo")

	webapi.Area("/api/1.0/", func() {
		webapi.RegisterPOST("/product/ToEntity", productApp.ToEntity)
	})
	webapi.UseApiResponse()
	webapi.Run()
}
