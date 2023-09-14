// 该文件由fsctl route命令自动生成，请不要手动修改此文件
package main

import (
	"github.com/farseer-go/webapi"
	"shopping/application/orderApp"
	"shopping/application/procateApp"
	"shopping/application/productApp"
)

var route = []webapi.Route{
    {"GET", "/area/order/list", orderApp.List, "查询成功", []string{"pageSize", "pageIndex", ""}},
    {"GET", "/area/order/count", orderApp.Count, "查询成功", []string{""}},
    {"GET", "/cate/list", procateApp.List, "查询成功", []string{""}},
    {"GET", "/product/info", productApp.ToEntity, "查询成功", []string{"productId", "", ""}},
    {"GET", "/product/list", productApp.List, "查询成功", []string{"cateId", "pageSize", "pageIndex", "", ""}},
    {"GET", "/product/buy", productApp.Buy, "下单成功", []string{"productId", "", "", "default", "buyOrder"}},
}
