// 该文件由fsctl route命令自动生成，请不要手动修改此文件
package main

import (
	"github.com/farseer-go/webapi"
	"github.com/farseer-go/webapi/context"
	"shopping/application/orderApp"
	"shopping/application/procateApp"
	"shopping/application/productApp"
)

var route = []webapi.Route{
    {"GET", "/api/1.0/order/List", orderApp.List, "查询成功", []context.IFilter{Jwt{},Auth{}}, []string{"pageSize", "pageIndex", ""}},
    {"GET", "/api/1.0/order/Count", orderApp.Count, "查询成功", []context.IFilter{Jwt{},Auth{}}, []string{""}},
    {"GET", "/api/1.0/cate/List", procateApp.List, "查询成功", []context.IFilter{Jwt{},Auth{}}, []string{""}},
    {"GET", "/api/1.0/product/info", productApp.ToEntity, "查询成功", []context.IFilter{Jwt{},Auth{}}, []string{"productId", "", ""}},
    {"GET", "/api/1.0/product/List", productApp.List, "查询成功", []context.IFilter{Jwt{},Auth{}}, []string{"cateId", "pageSize", "pageIndex", "", ""}},
    {"POST", "/api/1.0/product/Buy", productApp.Buy, "下单成功", []context.IFilter{Jwt{},Auth{}}, []string{"productId", "", "", "default", "buyOrder"}},
}
