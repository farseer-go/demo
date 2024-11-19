package main

import "github.com/farseer-go/webapi/context"

// IFilter 过滤器
type Jwt struct {
}

func (receiver Jwt) OnActionExecuting(httpContext *context.HttpContext) {

}

func (receiver Jwt) OnActionExecuted(httpContext *context.HttpContext) {

}
