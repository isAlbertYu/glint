# glint
一款自行开发、小巧轻便的Go Web应用框架

## 相关特性：
#### version 1.0.1：
定义了context上下文内容，在context中包含了一次HTTP交互（请求与响应）的全部信息。
type Context struct {
	Req        *http.Request       //自客户端的http请求
	Writer     http.ResponseWriter //服务端的响应输出流
	Path       string              //路由url
	Method     string              //请求方法
	StatusCode int                 //状态码
}

#### version 1.0.0：
静态路由，一切交由Engin处理：
```go
type HandlerFunc func(http.ResponseWriter, *http.Request)

//http服务端的处理引擎，内置一个路由表，根据路由表做出相应的处理动作
type Engine struct {
	router map[string]HandlerFunc
}
```