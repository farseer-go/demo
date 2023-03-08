package main

import (
	"fullExample/application/orderApp"
	"fullExample/application/procateApp"
	"fullExample/application/productApp"
	"github.com/farseer-go/fs"
	"github.com/farseer-go/webapi"
)

func main() {
	fs.Initialize[StartupModule]("demo")

	// 让所有api带前缀："/api/1.0/"
	webapi.Area("/api/1.0/", func() {
		// 产品分类列表
		// get http://localhost:8888/api/1.0/cate/list
		webapi.RegisterGET("/cate/list", procateApp.ToList)

		// 购买产品
		// get http://localhost:8888/api/1.0/product/buy
		webapi.RegisterPOST("/product/buy", productApp.Buy)

		// 产品信息
		// get http://localhost:8888/api/1.0/product/info?productId=1
		webapi.RegisterGET("/product/info", productApp.ToEntity)

		// 产品列表
		// get http://localhost:8888/api/1.0/product/list?pageIndex=1&pageSize=3&cateId=0
		webapi.RegisterGET("/product/list", productApp.ToList, "cateId", "pageSize", "pageIndex", "", "")

		// 订单列表
		// get http://localhost:8888/api/1.0/order/list?pageIndex=1&pageSize=3
		webapi.RegisterGET("/order/list", orderApp.ToList, "pageSize", "pageIndex", "")

		// 订单数量
		// get http://localhost:8888/api/1.0/order/count
		webapi.RegisterGET("/order/count", orderApp.Count)
	})

	// 让所有的返回值，包含在core.ApiResponse中
	webapi.UseApiResponse()
	// 使用静态文件 在根目录./wwwroot中的文件
	webapi.UseStaticFiles()
	// 运行web服务，端口配置在：farseer.yaml Webapi.Url 配置节点
	webapi.Run()
}
