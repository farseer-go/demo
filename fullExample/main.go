package main

import (
	"fullExample/application/procateApp"
	"fullExample/application/productApp"
	"fullExample/interfaces"
	"github.com/farseer-go/fs"
	"github.com/farseer-go/webapi"
	"github.com/farseer-go/webapi/controller"
)

func main() {
	fs.Initialize[StartupModule]("demo")
	webapi.RegisterController(&interfaces.HomeController{
		BaseController: controller.BaseController{
			Action: map[string]controller.Action{
				"Index": {Method: "GET", Params: ""},
			},
		},
	})

	// 让所有api带前缀："/api/1.0/"
	webapi.Area("/api/1.0/", func() {
		// get http://localhost:8888/api/1.0/cate/list
		webapi.RegisterGET("/cate/list", procateApp.ToList)

		// get http://localhost:8888/api/1.0/product/info?productId=1
		webapi.RegisterGET("/product/info", productApp.ToEntity)

		// get http://localhost:8888/api/1.0/product/list?pageIndex=1&pageSize=3
		webapi.RegisterGET("/product/list", productApp.ToList, "cateId", "pageSize", "pageIndex", "repository")
	})

	// 让所有的返回值，包含在core.ApiResponse中
	webapi.UseApiResponse()
	// 使用静态文件 在根目录./wwwroot中的文件，直接以静态文件提供服务
	webapi.UseStaticFiles()
	// 运行web服务，端口配置在：farseer.yaml Webapi.Url 配置节点
	webapi.Run()
}
