# glint
一款自行开发、小巧轻便的Go Web应用框架

## 相关特性：

#### version 1.0.2：
定义了路由类router，Engine中包含一个router，路由映射则移交至router类来管理
```go
type router struct {
	handlers map[string]HandlerFunc
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handlerFunc, ok := r.handlers[key]; ok {
		handlerFunc(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
```

Engine类实现ServeHTTP接口，将一次http交互的上下文交予router处理
```go
type Engine struct {
	router *router
}

//使得Engine实现Handler接口，可以注册到http服务端
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
```

#### version 1.0.1：
定义了context上下文内容，在context中包含了一次HTTP交互（请求与响应）的全部信息。
```go
type Context struct {
	Req        *http.Request       //自客户端的http请求
	Writer     http.ResponseWriter //服务端的响应输出流
	Path       string              //路由url
	Method     string              //请求方法
	StatusCode int                 //状态码
}
```

#### version 1.0.0：
静态路由，一切交由Engin处理：
```go
type HandlerFunc func(http.ResponseWriter, *http.Request)

//http服务端的处理引擎，内置一个路由表，根据路由表做出相应的处理动作
type Engine struct {
	router map[string]HandlerFunc
}
```