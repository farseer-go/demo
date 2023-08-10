package main

import (
	"github.com/farseer-go/fs"
	"github.com/farseer-go/webapi"
)

// @test
func main() {
	fs.Initialize[StartupModule]("demo")
	/*
		// 让所有api带前缀："/api/1.0/"
		webapi.Area("/api/1.0/", func() {
			// 商品分类列表
			// get http://localhost:8888/api/1.0/cate/list
			webapi.RegisterGET("/cate/list", procateApp.List)
			// 购买商品
			// get http://localhost:8888/api/1.0/product/buy
			webapi.RegisterPOST("/product/buy", productApp.Buy, "productId", "", "", "default", "buyOrder")
			// 商品信息
			// get http://localhost:8888/api/1.0/product/info?productId=1
			webapi.RegisterGET("/product/info", productApp.ToEntity)
			// 商品列表
			// get http://localhost:8888/api/1.0/product/list?pageIndex=1&pageSize=3&cateId=0
			webapi.RegisterGET("/product/list", productApp.List, "cateId", "pageSize", "pageIndex", "", "")
			// 订单列表
			// get http://localhost:8888/api/1.0/order/list?pageIndex=1&pageSize=3
			webapi.RegisterGET("/order/list", orderApp.List, "pageSize", "pageIndex", "")
			// 订单数量
			// get http://localhost:8888/api/1.0/order/count
			webapi.RegisterGET("/order/count", orderApp.Count)
		})
	*/
	webapi.RegisterRoutes(route...)

	// 让所有的返回值，包含在core.ApiResponse中
	webapi.UseApiResponse()
	// 使用静态文件 在根目录./wwwroot中的文件
	webapi.UseStaticFiles()
	// 运行web服务，端口配置在：farseer.yaml Webapi.Url 配置节点
	webapi.Run()
}
