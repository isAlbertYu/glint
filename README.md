# glint
一款自行开发、小巧轻便的Go Web应用框架

相关特性：
version 1：
静态路由，一切由Engin处理：
```go
type HandlerFunc func(http.ResponseWriter, *http.Request)

//http服务端的处理引擎，内置一个路由表，根据路由表做出相应的处理动作
type Engine struct {
	router map[string]HandlerFunc
}
```