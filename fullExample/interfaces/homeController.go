package interfaces

import (
	"github.com/farseer-go/webapi/action"
	"github.com/farseer-go/webapi/controller"
)

// HomeController 这里演示MVC模式
type HomeController struct {
	controller.BaseController // 定义控制器基类
}

// Index 首页
func (receiver *HomeController) Index() action.IResult {
	// 视图路径：./views/home/index.html
	// 视图规则：./views/[route][.html]

	data := make(map[string]any)
	data["procate"] = 123
	
	// 使用视图的方式
	return action.ViewData(data)
}
