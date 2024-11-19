package main

import "github.com/farseer-go/webapi/context"

// IFilter 过滤器
type Auth struct {
}

func (receiver Auth) OnActionExecuting(httpContext *context.HttpContext) {

}

func (receiver Auth) OnActionExecuted(httpContext *context.HttpContext) {

}
