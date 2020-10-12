package glint

import (
	"log"
	"net/http"
)

//定义了基于map实现的路由表，可以高效查询
//请求方法-匹配模式 : 处理函数
type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

//添加路由映射
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

//处理路由
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handlerFunc, ok := r.handlers[key]; ok {
		handlerFunc(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
